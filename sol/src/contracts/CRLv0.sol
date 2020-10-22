// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

contract CRLv0 {
  string public version;
  bool public finalised;
  uint8[] public data;

  constructor() public {}

  function append(uint8[] memory part, bool isFinal) public {
    assert(finalised == false);
    uint inLength = part.length;

    for (uint index = 0; index < inLength; index++) {
      data.push(part[index]); 
    }
    finalised = isFinal;
  }

  function length() public view returns (uint256) {return data.length;}
}
