// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

import "./Finalizable.sol";

contract TBSCertList is Finalizable {
  uint8[] public raw;
	int32 public verison;
	uint8[] public signature;
	uint8[] public issuer;
	uint256 public thisUpdate;
	uint256 public nextUpdate;
	uint8[] public extentions;
  address[] revokedCertList;

  function addRevokedCert(address ref) public onlyNotFinalized {   
    Finalizable cert = Finalizable(ref);
    assert(cert.isFinalized());
    revokedCertList.push(ref);
  }

  function setRaw(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      raw.push(data[index]); 
    }
  }

  function setSignature(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      signature.push(data[index]); 
    }
  }

  function setIssuer(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      issuer.push(data[index]); 
    }
  }

  function setExtentions(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      extentions.push(data[index]); 
    }
  }

  function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 version) public onlyNotFinalized {
    thisUpdate = thisUpd;
    nextUpdate = nextUpd;
    version = version;
  }
}
