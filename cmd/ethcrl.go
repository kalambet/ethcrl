package main

import (
	"crypto/x509"
	"encoding/asn1"
	"ethcrl"
	"ethcrl/crl"
	"ethcrl/eth"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	derFile = flag.String("der", "", "path to DER CRL file")
	pKey    = flag.String("key", "7b90ee8f413863ae7d1dea2886f3bec462f56ebb22a9bf379a9a76f5de7ba335", "private key data")
	backend = flag.String("backend", "http://localhost:9545", "address to eth node listening on port")
)

func main() {
	flag.Parse()
	flag.Usage()

	eth.InitEthClient(*backend, *pKey)

	_crl := make([]byte, 0)
	if len(*derFile) != 0 {
		f, err := os.Open(*derFile)
		defer f.Close()
		_crl, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("error reading CRL DER file: %s", err)
		}
		log.Printf("Using data from provided DER file %s", *derFile)
	} else {
		var err error
		_crl, err = crl.GetCRLBytes()
		if err != nil {
			log.Fatalf("error reading CRL DER data from embeded test data: %s", err)
		}
		log.Println("Using embeded example CRL")
	}

	list, err := x509.ParseDERCRL(_crl)
	if err != nil {
		log.Fatalf("error parsing CRL: %s", err)
	}

	list = ethcrl.MustGet(ethcrl.MustPush(list))

	_crlBack, err := asn1.Marshal(*list)
	if err != nil {
		log.Fatalf("error marshaling CRL: %s", err)
	}

	if len(_crl) != len(_crlBack) {
		log.Fatalf(
			"lenth of sored and retrieved certificate differs: stored = %d, retrieved = %d", len(_crl), len(_crlBack))
	}

	for idx := 0; idx < len(_crl); idx++ {
		if _crl[idx] != _crlBack[idx] {
			log.Fatalf(
				"byte %d in stored and retrieved certs differs: stored = %b, retrieved = %b", idx, _crl[idx], _crlBack[idx])
		}
	}
}
