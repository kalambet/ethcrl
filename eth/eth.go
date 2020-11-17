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
	Key string
)

func InitEthClient(addr string, key string) {
	var err error
	Cli, err = ethclient.Dial(addr)
	if err != nil {
		log.Fatalf("error initialising Eth client: %s", err)
	}
	Key = key
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
	privateKey, err := crypto.HexToECDSA(Key)
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
