package ethcrl

import (
	"context"
	"crypto/x509/pkix"
	"encoding/asn1"
	"ethcrl/sol/gocontracts/rc"
	"fmt"
	"log"
	"math/big"
	"time"

	"ethcrl/eth"
	"ethcrl/sol/gocontracts/crlv0"
	"ethcrl/sol/gocontracts/tbscl"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	quantum = 128
)

func MustPush(list *pkix.CertificateList) common.Address {
	addr, err := deployCRLv0(list)
	if err != nil {
		log.Fatalf("error deploying CRLv0 cntract: %s", err)
	}

	return *addr
}

func deployCRLv0(list *pkix.CertificateList) (*common.Address, error) {
	opts, err := eth.Translator()
	if err != nil {
		return nil, fmt.Errorf("error getting tx options: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	addr, tx, crlContract, err := crlv0.DeployCRLv0(opts, eth.Cli)
	if err != nil {
		return nil, fmt.Errorf("error deploying the CRLv0 contract: %s", err)
	}

	addr, err = bind.WaitDeployed(context.Background(), eth.Cli, tx)
	if err != nil {
		return nil, fmt.Errorf("error wating for a CRLv0 contract to deploy: %s", err)
	}
	log.Printf("CRLv0 deployed at %s", addr.String())

	data, err := asn1.Marshal(list.SignatureAlgorithm)
	if err != nil {
		return nil, fmt.Errorf("error encoding Signture Alorithm filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = crlContract.SetSigAlgo(opts, data)
	if err != nil {
		return nil, fmt.Errorf("error uploading Signture Alorithm data: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	data, err = asn1.Marshal(list.SignatureValue)
	if err != nil {
		return nil, fmt.Errorf("error encoding Signture Value filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = crlContract.SetSigValue(opts, data)
	if err != nil {
		return nil, fmt.Errorf("error uploading Signture Value data: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	tbsAddr, err := deployTBSCertList(list.TBSCertList, opts)
	if err != nil {
		return nil, fmt.Errorf("error deploying TBSCertList contract: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = crlContract.SetTBSCertList(opts, *tbsAddr)
	if err != nil {
		return nil, fmt.Errorf("error setting TBSCertList contract address: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	opts, err = eth.IterateOptions(opts)
	tx, err = crlContract.Finalize(opts, true)
	if err != nil {
		return nil, fmt.Errorf("error finalising CRLv0 contract: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func deployTBSCertList(list pkix.TBSCertificateList, opts *bind.TransactOpts) (*common.Address, error) {
	opts, err := eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	addr, tx, tbsContract, err := tbscl.DeployTBSCertList(opts, eth.Cli)
	if err != nil {
		return nil, fmt.Errorf("error deploying the TBSCertList contract: %s", err)
	}

	addr, err = bind.WaitDeployed(context.Background(), eth.Cli, tx)
	if err != nil {
		return nil, fmt.Errorf("error wating for a TBSCertList contract to deploy: %s", err)
	}
	log.Printf("TBSCertList deployed at %s", addr.String())

	data, err := asn1.Marshal(list.Raw)
	if err != nil {
		return nil, fmt.Errorf("error encoding Raw data filed: %s", err)
	}

	total := len(data)
	parts := total / quantum
	var cursor, _len int
	for i := 0; i <= parts; i++ {
		_len = quantum
		if i == parts {
			_len = total - parts*quantum
		}
		opts, err = eth.IterateOptions(opts)
		if err != nil {
			return nil, err
		}
		tx, err = tbsContract.SetRaw(opts, data[cursor:cursor+_len])
		if err != nil {
			return nil, fmt.Errorf("error setting Raw data: %s", err)
		}
		err = eth.WaitTx(tx)
		if err != nil {
			return nil, err
		}
		cursor += _len
	}

	data, err = asn1.Marshal(list.Signature)
	if err != nil {
		return nil, fmt.Errorf("error encoding Signture filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = tbsContract.SetSignature(opts, data)
	if err != nil {
		return nil, fmt.Errorf("error setting Signature: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	data, err = asn1.Marshal(list.Issuer)
	if err != nil {
		return nil, fmt.Errorf("error encoding Issuer filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = tbsContract.SetIssuer(opts, data)
	if err != nil {
		return nil, fmt.Errorf("error setting Issuer: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	data, err = asn1.Marshal(list.Extensions)
	if err != nil {
		return nil, fmt.Errorf("error encoding Extensions filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = tbsContract.SetExtensions(opts, data)
	if err != nil {
		return nil, fmt.Errorf("error setting Extensions: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = tbsContract.SetTimeAndVersion(opts, big.NewInt(list.ThisUpdate.Unix()), big.NewInt(list.NextUpdate.Unix()), int32(list.Version))
	if err != nil {
		return nil, fmt.Errorf("error setting Times and Version: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}

	addrs, err := deployRevokeCertList(list.RevokedCertificates, opts)
	if err != nil {
		return nil, fmt.Errorf("error deploying Revoked Cert List: %s", err)
	}

	for idx := range addrs {
		opts, err := eth.IterateOptions(opts)
		if err != nil {
			return nil, err
		}
		tx, err = tbsContract.AddRevokedCert(opts, addrs[idx])
		if err != nil {
			return nil, fmt.Errorf("error adding revoked cert: %s", err)
		}
		err = eth.WaitTx(tx)
		if err != nil {
			return nil, err
		}
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = tbsContract.Finalize(opts, true)
	if err != nil {
		return nil, fmt.Errorf("error finalising TBSCertList contract: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func deployRevokeCertList(certs []pkix.RevokedCertificate, opts *bind.TransactOpts) ([]common.Address, error) {
	res := make([]common.Address, 0, len(certs))
	for idx := range certs {
		opts, err := eth.IterateOptions(opts)
		if err != nil {
			return nil, err
		}
		addr, tx, rcContract, err := rc.DeployRevokedCert(opts, eth.Cli)
		if err != nil {
			return nil, fmt.Errorf("error deploying the RevokedCert contract: %s", err)
		}

		addr, err = bind.WaitDeployed(context.Background(), eth.Cli, tx)
		if err != nil {
			return nil, fmt.Errorf("error wating for a RevokedCert contract to deploy: %s", err)
		}
		log.Printf("RevokedCert[%d] deployed at %s", idx, addr.String())

		data, err := asn1.Marshal(certs[idx].Extensions)
		if err != nil {
			return nil, fmt.Errorf("error encoding Extensions filed: %s", err)
		}

		opts, err = eth.IterateOptions(opts)
		if err != nil {
			return nil, err
		}
		tx, err = rcContract.SetExtensions(opts, data)
		if err != nil {
			return nil, fmt.Errorf("error setting Extensions: %s", err)
		}
		err = eth.WaitTx(tx)
		if err != nil {
			return nil, err
		}

		opts, err = eth.IterateOptions(opts)
		if err != nil {
			return nil, err
		}
		tx, err = rcContract.SetCertData(opts, certs[idx].SerialNumber, big.NewInt(certs[idx].RevocationTime.Unix()))
		if err != nil {
			return nil, fmt.Errorf("error setting CertData: %s", err)
		}
		err = eth.WaitTx(tx)
		if err != nil {
			return nil, err
		}

		opts, err = eth.IterateOptions(opts)
		if err != nil {
			return nil, err
		}
		tx, err = rcContract.Finalize(opts, true)
		if err != nil {
			return nil, fmt.Errorf("error Finalizing RevokedCert conteract: %s", err)
		}
		err = eth.WaitTx(tx)
		if err != nil {
			return nil, err
		}
		res = append(res, addr)
	}
	return res, nil
}

func MustGet(addr common.Address) *pkix.CertificateList {
	list, err := buildCertificateList(addr)
	if err != nil {
		log.Fatalf("error building certificate back: %s", err)
	}
	return list
}

func buildCertificateList(addr common.Address) (*pkix.CertificateList, error) {
	log.Printf("restoring CertificateList from contract: %s", addr.String())
	trans, err := eth.Translator()
	if err != nil {
		return nil, err
	}
	opts := &bind.CallOpts{From: trans.From}

	crlContract, err := crlv0.NewCRLv0Caller(addr, eth.Cli)
	if err != nil {
		return nil, fmt.Errorf("error instantiating CRLv0 contract: %s", err)
	}
	res := new(pkix.CertificateList)

	var iter int64 = 0
	data := make([]byte, 0)
	for {
		piece, err := crlContract.SigAlgo(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	ai := new(pkix.AlgorithmIdentifier)
	_, err = asn1.Unmarshal(data, ai)
	if err != nil {
		return nil, err
	}
	res.SignatureAlgorithm = *ai

	iter = 0
	data = make([]byte, 0)
	for {
		piece, err := crlContract.SigValue(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	value := new(asn1.BitString)
	_, err = asn1.Unmarshal(data, value)
	if err != nil {
		return nil, err
	}
	res.SignatureValue = *value

	tbsAddr, err := crlContract.TbsCertList(opts)
	if err != nil {
		return nil, fmt.Errorf("error retrieving TBSCertList contract address: %s", err)
	}

	tbs, err := buildTBSCertList(tbsAddr, opts)
	if err != nil {
		return nil, fmt.Errorf("error building TBSCertList contract address: %s", err)
	}
	res.TBSCertList = *tbs
	return res, nil

}

func buildTBSCertList(addr common.Address, opts *bind.CallOpts) (*pkix.TBSCertificateList, error) {
	log.Printf("restoring TBSCertificateList from contract: %s", addr.String())
	tbsContract, err := tbscl.NewTBSCertListCaller(addr, eth.Cli)
	if err != nil {
		return nil, fmt.Errorf("error instantiating TBSCertList contract: %s", err)
	}
	res := new(pkix.TBSCertificateList)

	var iter int64 = 0
	data := make([]byte, 0)
	for {
		piece, err := tbsContract.Raw(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	rawValue := new(asn1.RawContent)
	_, err = asn1.Unmarshal(data, rawValue)
	if err != nil {
		return nil, err
	}
	res.Raw = *rawValue

	iter = 0
	data = make([]byte, 0)
	for {
		piece, err := tbsContract.Issuer(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	rdnsValue := new(pkix.RDNSequence)
	_, err = asn1.Unmarshal(data, rdnsValue)
	if err != nil {
		return nil, err
	}
	res.Issuer = *rdnsValue

	iter = 0
	data = make([]byte, 0)
	for {
		piece, err := tbsContract.Signature(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	algoIdentValue := new(pkix.AlgorithmIdentifier)
	_, err = asn1.Unmarshal(data, algoIdentValue)
	if err != nil {
		return nil, err
	}
	res.Signature = *algoIdentValue

	iter = 0
	data = make([]byte, 0)
	for {
		piece, err := tbsContract.Extensions(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	extValue := new([]pkix.Extension)
	_, err = asn1.Unmarshal(data, extValue)
	if err != nil {
		return nil, err
	}
	res.Extensions = *extValue

	nextUpdEpoch, err := tbsContract.NextUpdate(opts)
	if err != nil {
		return nil, err
	}
	res.NextUpdate = time.Unix(nextUpdEpoch.Int64(), 0)

	thisUpdEpoch, err := tbsContract.ThisUpdate(opts)
	if err != nil {
		return nil, err
	}
	res.ThisUpdate = time.Unix(thisUpdEpoch.Int64(), 0)

	version, err := tbsContract.Version(opts)
	if err != nil {
		return nil, err
	}
	res.Version = int(version)

	res.RevokedCertificates = make([]pkix.RevokedCertificate, 0)
	iter = 0
	for {
		addr, err := tbsContract.RevokedCertList(opts, big.NewInt(iter))
		if err != nil {
			break
		}

		rcert, err := buildRevokedCert(addr, opts)
		if err != nil {
			return nil, fmt.Errorf("error building Revoked Certificate: %s", err)
		}
		res.RevokedCertificates = append(res.RevokedCertificates, *rcert)
		iter++
	}

	return res, nil
}

func buildRevokedCert(addr common.Address, opts *bind.CallOpts) (*pkix.RevokedCertificate, error) {
	log.Printf("restoring RevokedCertificate from contract: %s", addr.String())
	rcContract, err := rc.NewRevokedCertCaller(addr, eth.Cli)
	if err != nil {
		return nil, fmt.Errorf("error instantiating RevokedCert contract: %s", err)
	}
	res := new(pkix.RevokedCertificate)

	var iter int64 = 0
	data := make([]byte, 0)
	for {
		piece, err := rcContract.Extensions(opts, big.NewInt(iter))
		if err != nil {
			break
		}
		data = append(data, piece)
		iter++
	}

	extValue := new([]pkix.Extension)
	_, err = asn1.Unmarshal(data, extValue)
	if err != nil {
		return nil, err
	}
	res.Extensions = *extValue

	serial, err := rcContract.SerialNumber(opts)
	if err != nil {
		return nil, err
	}
	res.SerialNumber = serial

	timeEpoch, err := rcContract.RevocationTime(opts)
	if err != nil {
		return nil, err
	}
	res.RevocationTime = time.Unix(timeEpoch.Int64(), 0)
	return res, nil
}
