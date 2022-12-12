// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract BucketPayment is OwnableUpgradeable {
    uint bucketFee;
    address paymentToken;
    mapping(bytes8 => address) private bucketToUsers;
    mapping(address => bytes8[]) private userToBuckets;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(uint initialBucketFee, address acceptedToken) public initializer {
        bucketFee = initialBucketFee;
        paymentToken = acceptedToken;
        __Ownable_init();
    }

    function payForBucket(bytes8 bucketId) public {
        require(bucketToUsers[bucketId] == address(0), 'Bucket: already owned');
        require(IERC20(paymentToken).allowance(msg.sender, address(this)) >= bucketFee, 'ERC20: allowance is less than bucket fee');
    
        IERC20(paymentToken).transferFrom(msg.sender, address(this), bucketFee);
    }

    function withdrawTokens(address tokenAddress) public onlyOwner {
        uint256 balance = IERC20(tokenAddress).balanceOf(address(this));
        IERC20(tokenAddress).transfer(msg.sender, balance);
    }
}