// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    uint256 private storedData;
    event DataStored(uint256 data);

    // write a new value
    function set(uint256 x) public {
        storedData = x;
        emit DataStored(x);
    }

    // read the stored value
    function get() public view returns (uint256) {
        return storedData;
    }
}
