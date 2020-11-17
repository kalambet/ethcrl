# Pack CRL to Ethereum

This repo provides example implementation of an approach to place CRL (as Certificate Revocation List) to Ethereum as a set of smart contracts.

## Prerequisites

To run this code you'll need:
1. [Go](https://golang.org/doc/install)
2. [Truffle Suite](https://www.trufflesuite.com/docs/truffle/getting-started/installation)
3. [abigen](https://github.com/ethereum/go-ethereum)

## Steps
* Clone and build repo
```bash
$: git clone github.com/kalambet/ethcrl
$: cd ethcrl && go mod download
```
* Obtain supported solidity compiler
```bash
$: cd sol/src && truffle obtain && cd -
```
* Generate binding contracts with `abigen`
```bash
$: cd sol && ./gen.sh && cd -
```
* On a _separate console_ run Truffle Develop from `<repo_path>/sol/src` folder and leave it as is
```bash
$: cd sol/src && truffle develop
``` 
* Copy the first private key from the provided output
* Build an example code
```bash
$: go build cmd/ethcrl.go
```
* Run binary with previously copied private key
```bash
$: ./ethcrl -key <PRIVATE_KEY>
```

## Alternative CRL data
It is possible to test arbitrary CRL certificates, for that `--der` key can be provided.
For example, it is possible to download external CRL file, ex:
```bash
wget "http://testca2012.cryptopro.ru/cdp/064b632533662a24381872437a3bb7cbb2cafc73.crl"
``` 

And provide it as an input to `ethcrl`:
```bash
./bashcrl -key <PRIVATE_KEY> -der "./064b632533662a24381872437a3bb7cbb2cafc73.crl"
``` 

## Usage
```
Usage of ./ethcrl:
  -backend string
    	address to eth node listening on port (default "http://localhost:9545")
  -der string
    	path to DER CRL file
  -key string
    	private key data (default "7b90ee8f413863ae7d1dea2886f3bec462f56ebb22a9bf379a9a76f5de7ba335")
```
