package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Cli *ethclient.Client
)

const (
	key = "7b90ee8f413863ae7d1dea2886f3bec462f56ebb22a9bf379a9a76f5de7ba335"
)

func init() {
	var err error
	Cli, err = ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatalf("error initialising Eth client: %s", err)
	}
}

func IterateOptions(opts *bind.TransactOpts) (*bind.TransactOpts, error) {
	if opts == nil {
		return nil, errors.New("initial options needs to be provided")
	}
	nonce, err := Cli.PendingNonceAt(context.Background(), opts.From)
	if err != nil {
		return nil, err
	}
	log.Printf("nonce: %d\n", nonce)

	gasPrice, err := Cli.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	log.Printf("gas price: %d\n", gasPrice)

	opts.Value = big.NewInt(0)
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = gasPrice
	opts.GasLimit = uint64(6700000)

	return opts, nil
}

func Translator() (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("account: %s\n", from.String())

	return bind.NewKeyedTransactor(privateKey), nil
}

func WaitTx(tx *types.Transaction) error {
	r, err := bind.WaitMined(context.Background(), Cli, tx)
	if err != nil {
		return err
	}
	if r.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("trnsaction failes %#v", r.Logs)
	}
	return nil
}
