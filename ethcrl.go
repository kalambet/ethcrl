package ethcrl

import (
	"context"
	"crypto/x509/pkix"
	"encoding/asn1"
	"ethcrl/sol/gocontracts/rc"
	"fmt"
	"log"
	"math/big"

	"ethcrl/eth"
	"ethcrl/sol/gocontracts/crlv0"
	"ethcrl/sol/gocontracts/tbscl"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func MustPush(list *pkix.CertificateList) {
	err := deployCRLv0(list)
	if err != nil {
		log.Fatalf("error deploying CRLv0 cntract: %s", err)
	}
}

func deployCRLv0(list *pkix.CertificateList) error {
	opts, err := eth.Translator()
	if err != nil {
		return fmt.Errorf("error getting tx options: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return err
	}
	addr, tx, crlContract, err := crlv0.DeployCRLv0(opts, eth.Cli)
	if err != nil {
		return fmt.Errorf("error deploying the CRLv0 contract: %s", err)
	}

	addr, err = bind.WaitDeployed(context.Background(), eth.Cli, tx)
	if err != nil {
		return fmt.Errorf("error wating for a CRLv0 contract to deploy: %s", err)
	}
	log.Printf("CRLv0 deployed at %s", addr.String())

	data, err := asn1.Marshal(list.SignatureAlgorithm)
	if err != nil {
		return fmt.Errorf("error encoding Signture Alorithm filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return err
	}
	tx, err = crlContract.SetSigAlgo(opts, data)
	if err != nil {
		return fmt.Errorf("error uploading Signture Alorithm data: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return err
	}

	data, err = asn1.Marshal(list.SignatureValue)
	if err != nil {
		return fmt.Errorf("error encoding Signture Value filed: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return err
	}
	tx, err = crlContract.SetSigValue(opts, data)
	if err != nil {
		return fmt.Errorf("error uploading Signture Value data: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return err
	}

	tbsAddr, err := deployTBSCertList(list.TBSCertList, opts)
	if err != nil {
		return fmt.Errorf("error deploying TBSCertList contract: %s", err)
	}

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return err
	}
	tx, err = crlContract.SetTBSCertList(opts, *tbsAddr)
	if err != nil {
		return fmt.Errorf("error setting TBSCertList contract address: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return err
	}

	opts, err = eth.IterateOptions(opts)
	tx, err = crlContract.Finalize(opts, true)
	if err != nil {
		return fmt.Errorf("error finalising CRLv0 contract: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return err
	}
	return nil
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

	opts, err = eth.IterateOptions(opts)
	if err != nil {
		return nil, err
	}
	tx, err = tbsContract.SetRaw(opts, data)
	if err != nil {
		return nil, fmt.Errorf("error setting Raw data: %s", err)
	}
	err = eth.WaitTx(tx)
	if err != nil {
		return nil, err
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
	tx, err = tbsContract.SetExtentions(opts, data)
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
		tx, err = rcContract.SetExtentions(opts, data)
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

func MustGet() []byte {
	return nil
}
