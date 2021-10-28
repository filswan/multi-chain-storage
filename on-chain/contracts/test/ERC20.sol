//SPDX-License-Identifier: Unlicense
pragma solidity 0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TestERC20 is ERC20 {
    constructor (string memory _name, string memory _symbol) ERC20(_name, _symbol) {}

    function mint(address user, uint256 amount) external returns (bool) {
        _mint(user, amount);
        return true;
    }
}