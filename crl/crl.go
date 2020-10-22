package crl

import (
	"io/ioutil"

	"github.com/cloudflare/cfssl/crl"
)

const (
	tryTwoCert = "testdata/caTwo.pem"
	tryTwoKey  = "testdata/ca-keyTwo.pem"
	serialList = "testdata/serialList"
)

func GetCRLBytes() ([]byte, error) {
	tryTwoKeyBytes, err := ioutil.ReadFile(tryTwoKey)
	if err != nil {
		return nil, err
	}

	tryTwoCertBytes, err := ioutil.ReadFile(tryTwoCert)
	if err != nil {
		return nil, err
	}

	serialListBytes, err := ioutil.ReadFile(serialList)
	if err != nil {
		return nil, err
	}

	return crl.NewCRLFromFile(serialListBytes, tryTwoCertBytes, tryTwoKeyBytes, "0")
}
