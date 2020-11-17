// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

import "./Finalizable.sol";

contract RevokedCert is Finalizable {
  uint256 serialNumber;
  uint256 revocationTime;
  uint8[] extentions;

  function setExtentions(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      extentions.push(data[index]); 
    }
  }

  function setCertData(uint256 serial, uint256 time) public onlyNotFinalized {
    serialNumber = serial;
    revocationTime = time;
  }
}
