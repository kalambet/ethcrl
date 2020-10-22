package main

import (
	"crypto/x509"
	"ethcrl"
	"ethcrl/crl"
	"log"
)

func main() {
	_crl, err := crl.GetCRLBytes()
	if err != nil {
		log.Fatalf("error reading CRL: %s", err)
	}

	list, err := x509.ParseDERCRL(_crl)
	log.Printf("crl: %#v", list)

	ethcrl.MustPush(_crl)
	_crl = ethcrl.MustGet()

	list, err = x509.ParseDERCRL(_crl)
	if err != nil {
		log.Fatalf("error reading returned CRL: %s", err)
	}
	log.Printf("crl: %#v", list)
}
