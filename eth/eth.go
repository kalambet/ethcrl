package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"ethcrl/sol"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	Cli   bind.ContractBackend
	CRLv0 *sol.Sol
)

const (
	addr = "0x05FD872F322ef6e3D70A561Ad2234365fC692eAb"
	key  = "7b90ee8f413863ae7d1dea2886f3bec462f56ebb22a9bf379a9a76f5de7ba335"
)

func init() {
	var err error
	Cli, err = ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatalf("error initialising Eth client: %s", err)
	}

	address := common.HexToAddress(addr)
	CRLv0, err = sol.NewSol(address, Cli)
	if err != nil {
		log.Fatalf("error binding to contract: %s", err)
	}
}

func Options(opts *bind.TransactOpts, from common.Address) (*bind.TransactOpts, error) {
	nonce, err := Cli.PendingNonceAt(context.Background(), from)
	if err != nil {
		return nil, err
	}
	log.Printf("nonce: %d\n", nonce)

	gasPrice, err := Cli.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	log.Printf("gas price: %d\n", gasPrice)

	opts.From = from
	opts.Value = big.NewInt(0)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = gasPrice
	opts.GasLimit = uint64(6700000)

	return opts, nil
}

func Translator() (*bind.TransactOpts, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, [20]byte{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, [20]byte{}, errors.New("error casting public key to ECDSA")
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("from: %s\n", from)

	return bind.NewKeyedTransactor(privateKey), from, nil
}
