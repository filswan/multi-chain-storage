//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "hardhat/console.sol";
import "../interfaces/IPaymentGateway.sol";
import "../interfaces/IPriceFeed.sol";

abstract contract SwanPaymentNoRefund is IPaymentMinimal, Initializable {
    address public constant NATIVE_TOKEN =
        0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    address private _ERC20_TOKEN;

    address private _owner;
    address private _priceFeed;

    uint256 private lockTime;
    mapping(string => TxInfo) private txMap;
    mapping(string => TxInfo) private txCarMap;

    function initialize(
        address owner,
        address ERC20_TOKEN,
        address priceFeed
    ) public initializer {
        _owner = owner;
        _ERC20_TOKEN = ERC20_TOKEN;
        _priceFeed = priceFeed;
        lockTime = 5 days;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(_owner == msg.sender, "Caller is not the owner");
        _;
    }

    function setPriceFeed(address priceFeed) public onlyOwner returns (bool) {
        _priceFeed = priceFeed;
        return true;
    }

    function getLockedPaymentInfo(string calldata cId)
        public
        view
        override
        returns (TxInfo memory tx)
    {
        // default value is 0
        return txCarMap[cId];
    }

    event LockPayment(
        string id,
        address token,
        uint256 lockedFee,
        uint256 minPayment,
        address recipient,
        uint256 deadline,
        uint256 size,
        uint8 copyLimit
    );

    /// @notice Deposits the amount of token for specific transaction
    /// @param param The transaction information for which to deposit balance
    /// @return Returns true for a successful deposit, false for an unsuccessful deposit
    function lockTokenPayment(lockPaymentParam calldata param)
        public
        override
        returns (bool)
    {
        require(
            !txMap[param.id]._isExisted && !txCarMap[param.id]._isExisted,
            "Payment of transaction is already locked"
        );
       
        require(
            param.minPayment > 0 && param.amount > param.minPayment,
            "payment should greater than min payment"
        );

        require(
            IERC20(_ERC20_TOKEN).allowance(msg.sender, address(this)) >=
                param.amount,
            "please approve spending token"
        );
        IERC20(_ERC20_TOKEN).transferFrom(
            msg.sender,
            address(this),
            param.amount
        );

        if (param.size > 0) {
            TxInfo storage t = txCarMap[param.id];
            t.owner = msg.sender;
            t.minPayment = param.minPayment;
            t.recipient = param.recipient;
            t.deadline = block.timestamp + param.lockTime;
            t.lockedFee = param.amount;
            t.token = _ERC20_TOKEN;
            t._isExisted = true;
            t.size = param.size;
            t.copyLimit = param.copyLimit;
            t.blockNumber = block.number;
            
            emit LockPayment(
                param.id,
                t.token,
                t.lockedFee,
                param.minPayment,
                param.recipient,
                t.deadline,
                t.size,
                t.copyLimit
            );
        } else {
            TxInfo storage t = txMap[param.id];
            t.owner = msg.sender;
            t.minPayment = param.minPayment;
            t.recipient = param.recipient;
            t.deadline = block.timestamp + param.lockTime;
            t.lockedFee = param.amount;
            t.token = _ERC20_TOKEN;
            t._isExisted = true;
            t.copyLimit = param.copyLimit;
            t.blockNumber = block.number;

            // emit LockPayment(
            //     param.id,
            //     t.token,
            //     t.lockedFee,
            //     param.minPayment,
            //     param.recipient,
            //     t.deadline
            // );
        }

        return true;
    }

    
}