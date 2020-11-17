// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

import "./Finalizable.sol";

contract CRLv0 is Finalizable {
  bool public finalised;

  address public tbsCertList;
  uint8[] public sigAlgo;
  uint8[] public sigValue;

  function setSigAlgo(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      sigAlgo.push(data[index]); 
    }
  }

  function setSigValue(uint8[] memory data) public onlyNotFinalized {
    uint inLength = data.length;

    for (uint index = 0; index < inLength; index++) {
      sigValue.push(data[index]); 
    }
  }

  function setTBSCertList(address ref) public onlyNotFinalized {
    assert(tbsCertList == address(0));
    tbsCertList = ref;
  }
}
