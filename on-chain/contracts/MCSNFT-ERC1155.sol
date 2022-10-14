// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/ERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract MCSNFT is Initializable, ERC1155Upgradeable, OwnableUpgradeable {
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() initializer public {
        __ERC1155_init("");
        __Ownable_init();
    }

    function setURI(string memory uri) public onlyOwner {
        _setURI(uri);
    }

    // does not use data parameter
    function mintData(address account, uint256 id, uint256 amount)
        public
    {
        _mint(account, id, amount, "");
    }

    function mint(address account, uint256 id, uint256 amount, bytes memory data)
        public
        onlyOwner
    {
        _mint(account, id, amount, data);
    }

    function mintBatch(address to, uint256[] memory ids, uint256[] memory amounts, bytes memory data)
        public
        onlyOwner
    {
        _mintBatch(to, ids, amounts, data);
    }

    function contractURI() external pure returns (string memory) {
        return "https://calibration-ipfs.filswan.com/ipfs/QmXpf69sK8KJ6RurtbdHeS7dsSKzQiCoXB2XybgBoZuMRH";
    }
}