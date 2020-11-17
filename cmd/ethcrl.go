package main

import (
	"crypto/x509"
	"encoding/asn1"
	"ethcrl"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./064b632533662a24381872437a3bb7cbb2cafc73.crl")
	defer f.Close()

	_crl, err := ioutil.ReadAll(f)
	//_crl, err := crl.GetCRLBytes()
	if err != nil {
		log.Fatalf("error reading CRL: %s", err)
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
