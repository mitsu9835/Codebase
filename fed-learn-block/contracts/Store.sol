// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Store {
    // ─── EXISTING STORAGE ───────────────────────────────────────────
    bytes storedParams;
    bytes stored_g;
    bytes storedPK_C;
    bytes storedPK_F;
    string storedDId_M1;
    string stored_HashofF;
    bytes stored_HashofC;
    string storedDId_M2;
    string stored_HashofRC;
    string accessP;

    // ─── EXISTING METHODS ───────────────────────────────────────────
    function setParams(bytes memory input1, bytes memory input2) public {
        storedParams = input1;
        stored_g     = input2;
    }
    function getParams() public view returns (bytes memory, bytes memory) {
        return (storedParams, stored_g);
    }

    function setPKC(bytes memory input) public {
        storedPK_C = input;
    }
    function getPKC() public view returns (bytes memory) {
        return storedPK_C;
    }

    function setPKF(bytes memory input) public {
        storedPK_F = input;
    }
    function getPKF() public view returns (bytes memory) {
        return storedPK_F;
    }

    function setMeta1(
        string memory input1,
        string memory input2,
        bytes memory input3
    ) public {
        storedDId_M1   = input1;
        stored_HashofF = input2;
        stored_HashofC = input3;
    }
    function compareMeta1(string memory req) public view returns (bool) {
        return
            keccak256(abi.encodePacked(stored_HashofF)) ==
            keccak256(abi.encodePacked(req));
    }

    function setMeta2(string memory input1, string memory input2) public {
        storedDId_M2    = input1;
        stored_HashofRC = input2;
    }
    function compareMeta2(string memory req) public view returns (bool) {
        return
            keccak256(abi.encodePacked(stored_HashofRC)) ==
            keccak256(abi.encodePacked(req));
    }

    function setAccessP(string memory input1) public {
        accessP = input1;
    }
    function compareAccessPolicy(string memory req) public view returns (bool) {
        return
            keccak256(abi.encodePacked(accessP)) ==
            keccak256(abi.encodePacked(req));
    }

    // ─── NEW FUNCTION: GENERATE PSEUDO-NORMAL NOISE ──────────────────

    /// @notice Return a pseudo-random “noise” value ηᵢⱼ ≈ Normal(0,σ²)
    /// @dev Solidity can’t sample true normals; approximate via CLT over 4 uniforms.
    function generate_noise() public view returns (int256) {
        // seed from block data + sender
        uint256 seed = uint256(
            keccak256(
                abi.encodePacked(
                    block.timestamp,
                    block.prevrandao,        // replaced block.difficulty
                    msg.sender,
                    blockhash(block.number - 1)
                )
            )
        );

        // sum 4 uniform draws in [0,1e6)
        int256 sum = 0;
        for (uint256 i = 0; i < 4; i++) {
            seed = uint256(keccak256(abi.encodePacked(seed, i)));
            sum += int256(seed % 1_000_000);
        }

        // center at zero -> range ≈ [-2e6, +2e6]
        return sum - 2_000_000;
    }
}
