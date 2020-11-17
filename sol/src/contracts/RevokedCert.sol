// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

import "./Finalizable.sol";

contract RevokedCert is Finalizable {
  uint256 public serialNumber;
  uint256 public revocationTime;
  uint8[] public extensions;

  function setExtensions(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      extensions.push(data[index]);
    }
  }

  function setCertData(uint256 serial, uint256 time) public onlyNotFinalized {
    serialNumber = serial;
    revocationTime = time;
  }
}
