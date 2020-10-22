package ethcrl

import (
	"ethcrl/eth"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const quantum = 128

func MustPush(_crl []byte) {
	opts, from, err := eth.Translator()
	if err != nil {
		log.Fatalf("error getting tx options: %s", err)
	}

	total := len(_crl)
	parts := total / quantum
	var cursor, _len int
	for i := 0; i <= parts; i++ {
		_len = quantum
		if i == parts {
			_len = total - parts*quantum
		}
		log.Printf("slice[%d:%d]\n", cursor, cursor+_len)
		opts, err = eth.Options(opts, from)
		if err != nil {
			log.Fatalf("error updating tx options: %s", err)
		}
		tx, err := eth.CRLv0.Append(opts, _crl[cursor:cursor+_len], i == parts)
		if err != nil {
			log.Fatalf("error making transaction to contract: %s", err)
		}
		log.Printf("tx hash: %s\n", tx.Hash())
		cursor += _len
		log.Printf("cursor: %d\n", cursor)
	}
}

func MustGet() []byte {
	_, from, err := eth.Translator()
	if err != nil {
		log.Fatalf("error getting tx options: %s", err)
	}

	opts := new(bind.CallOpts)
	opts.From = from

	_len, err := eth.CRLv0.Length(opts)
	if err != nil {
		log.Fatalf("error reading data length: %s", err)
	}
	log.Printf("stored data length: %d", _len)

	res := make([]byte, 0, _len.Int64())
	var index int64
	for index = 0; index < _len.Int64(); index++ {
		data, err := eth.CRLv0.Data(opts, big.NewInt(index))
		if err != nil {
			log.Fatalf("error reading data from contract: %s", err)
		}
		res = append(res, data)
	}
	return res
}
