// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.8.0;

contract Finalizable {
  bool public finalized;
  
  modifier onlyNotFinalized() {
    require(finalized == false, "Finalizable: contract is already finalized");
    _;
  }
  
  function isFinalized() external view returns (bool) {return finalized;}
  function finalize(bool fin) external onlyNotFinalized {finalized = fin;}
}
