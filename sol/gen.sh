#!/bin/zsh

set -xe

abigen --pkg=rc --sol=src/contracts/RevokedCert.sol --out=gocontracts/rc/contract.go
abigen --pkg=tbscl --sol=src/contracts/TBSCertList.sol --out=gocontracts/tbscl/contract.go
abigen --pkg=crlv0 --sol=src/contracts/CRLv0.sol --out=gocontracts/crlv0/contract.go
