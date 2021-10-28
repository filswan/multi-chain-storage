//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "hardhat/console.sol";
import "./interfaces/IPaymentGateway.sol";
import "./FilswanOracle.sol";
import "./interfaces/IPriceFeed.sol";

contract SwanPayment is IPaymentMinimal, Initializable {
    address public constant NATIVE_TOKEN =
        0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    address private _ERC20_TOKEN;

    address private _owner;
    address private _oracle;
    address private _priceFeed;

    uint256 private lockTime;
    mapping(string => TxInfo) private txMap;

    function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed) public initializer {
        _owner = owner;
        _ERC20_TOKEN = ERC20_TOKEN;
        _oracle = oracle;
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

    function setOracle(address oracle) public onlyOwner returns (bool) {
        _oracle = oracle;
        return true;
    }

    function setPriceFeed(address priceFeed) public onlyOwner returns (bool) {
        _priceFeed = priceFeed;
        return true;
    }

    function getLockedPaymentInfo(string calldata txId)
        public
        view
        override
        returns (TxInfo memory tx)
    {
        // default value is 0
        return txMap[txId];
    }

    event LockPayment(
        string id,
        address token,
        uint256 lockedFee,
        uint256 minPayment,
        address recipient,
        uint256 deadline
    );

    /// @notice Deposits the amount of token for specific transaction
    /// @param param The transaction information for which to deposit balance
    /// @return Returns true for a successful deposit, false for an unsuccessful deposit
    function lockPayment(lockPaymentParam calldata param)
        public
        payable
        override
        returns (bool)
    {
        // require(
        //     !txMap[param.id]._isExisted,
        //     "Payment of transaction is already locked"
        // );
        // require(
        //     param.minPayment > 0 && msg.value > param.minPayment,
        //     "payment should greater than min payment"
        // );
        // TxInfo storage t = txMap[param.id];
        // t.owner = msg.sender;
        // t.minPayment = param.minPayment;
        // t.recipient = param.recipient;
        // t.deadline = block.timestamp + param.lockTime;
        // t.lockedFee = msg.value;
        // t.token = NATIVE_TOKEN;
        // t._isExisted = true;

        // emit LockPayment(
        //     param.id,
        //     NATIVE_TOKEN,
        //     t.lockedFee,
        //     param.minPayment,
        //     param.recipient,
        //     t.deadline
        // );
        return true;
    }

    

    /// @notice Returns the current allowance given to a spender by an owner
    /// @param txId transaction id
    /// @return Returns true for a successful payment, false for an unsuccessful payment
    function unlockPayment(string calldata txId)
        public
        override
        returns (bool)
    {
        return true;
        // TxInfo storage t = txMap[txId];
        // require(t._isExisted, "Transaction does not exist");
        // require(
        //     t.owner == msg.sender || t.recipient == msg.sender,
        //     "Invalid caller"
        // );
        // // if passed deadline, return payback to user
        // if (block.timestamp > t.deadline) {
        //     require(
        //         t.owner == msg.sender,
        //         "Tx passed deadline, only owner can get locked tokens"
        //     );
        //     t._isExisted = false;

        //     payable(address(t.owner)).transfer(t.lockedFee);
        // } else {
        //     uint256 actualFee = FilswanOracle(_oracle).getPaymentInfo(txId);
        //     require(actualFee > 0, "Transaction is incompleted");

        //     // todo: add convert rate function to get latest price

        //     if (actualFee < t.minPayment) {
        //         actualFee = t.minPayment;
        //     }
        //     t._isExisted = false;

        //     console.log("actualFee is %s", actualFee);
        //     console.log("locked fee is %s", t.lockedFee);

        //     if (t.lockedFee > actualFee) {
        //         payable(address(t.owner)).transfer(t.lockedFee - actualFee);
        //     } else {
        //         actualFee = t.lockedFee;
        //     }
        //     payable(address(t.recipient)).transfer(actualFee);
        // }
        // t.minPayment = 0;
        // t.lockedFee = 0;
       
        // return true;
    }

    /// @notice Deposits the amount of token for specific transaction
    /// @param param The transaction information for which to deposit balance
    /// @return Returns true for a successful deposit, false for an unsuccessful deposit
    function lockTokenPayment(lockPaymentParam calldata param)
        public
        override
        returns (bool)
    {
        require(
            !txMap[param.id]._isExisted,
            "Payment of transaction is already locked"
        );
        require(
            param.minPayment > 0 && param.amount > param.minPayment,
            "payment should greater than min payment"
        );

        require(IERC20(_ERC20_TOKEN).allowance(msg.sender, address(this)) >= param.amount, "please approve spending token");
        IERC20(_ERC20_TOKEN).transferFrom(msg.sender, address(this), param.amount);

        TxInfo storage t = txMap[param.id];
        t.owner = msg.sender;
        t.minPayment = param.minPayment;
        t.recipient = param.recipient;
        t.deadline = block.timestamp + param.lockTime;
        t.lockedFee = param.amount;
        t.token = _ERC20_TOKEN;
        t._isExisted = true;

        emit LockPayment(
            param.id,
            t.token,
            t.lockedFee,
            param.minPayment,
            param.recipient,
            t.deadline
        );
        return true;
    }


    /// @notice real fee is greater than tx.fee, take tx.fee
    /// real fee is less than tx.minPayment, take minPayment, return tx.fee - minPayment to tx.owner
    /// otherwise, take real fee, return tx.fee - real fee to tx.owner
    /// @param param data
    /// @return Returns true for a successful payment, false for an unsuccessful payment
    function unlockTokenPayment(
        unlockPaymentParam calldata param
    ) public override returns (bool){

        TxInfo storage t = txMap[param.id];
        require(t._isExisted, "Transaction does not exist");

        // if passed deadline, pay back to user
        if (block.timestamp > t.deadline) {
            require(
                t.owner == msg.sender,
                "Tx passed deadline, only owner can get locked tokens"
            );
            t._isExisted = false;
            t.minPayment = 0;
            t.lockedFee = 0;
            //payable(address(t.owner)).transfer(t.lockedFee);
            IERC20(_ERC20_TOKEN).transfer(t.owner, t.lockedFee);
        } else {
            require(param.amount > 0, "Transaction is incompleted");

            require(
                FilswanOracle(_oracle).isPaymentAvailable(
                    param.id,
                    param.orderId,
                    param.dealId,
                    param.amount,
                    param.recipient,
                    true
                ),
                "illegal unlock action"
            );
            // get spend token amount
            uint256 tokenAmount = IPriceFeed(_priceFeed).consult(_ERC20_TOKEN, param.amount);

            if (tokenAmount < t.minPayment) {
                tokenAmount = t.minPayment;
            }
            t._isExisted = false;

            console.log("actual token is %s", tokenAmount);

            if (t.lockedFee > tokenAmount) {
                uint256 tmp = t.lockedFee;
                t.lockedFee = 0; // prevent re-entrying
                IERC20(_ERC20_TOKEN).transfer(t.owner, tmp - tokenAmount);
                
            } else {
                tokenAmount = t.lockedFee;
                t.lockedFee = 0; // prevent re-entrying
            }
            IERC20(_ERC20_TOKEN).transfer(param.recipient, tokenAmount);
        }

        return true;
    }
}
