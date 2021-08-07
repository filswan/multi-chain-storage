//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "hardhat/console.sol";
import "./interfaces/IPaymentGateway.sol";

contract SwanPayment is IPaymentMinimal {
    uint256 lockTime = 1 days;
    mapping(string => TxInfo) txMap;

    constructor() {
        console.log("Deploying swan payment contract");
    }

    // /**
    //  * @dev Throws if called by any account other than the owner.
    //  */
    // modifier onlyParticipant() {
    //     require(_owner == msg.sender, "Caller is not the owner");
    //     _;
    // }

    function getLockedPaymentInfo(string calldata txId)
        public
        view
        override
        returns (TxInfo memory tx)
    {
        // default value is 0
        return txMap[txId];
    }

    /// @notice Deposits the amount of token for specific transaction
    /// @param param The transaction information for which to deposit balance
    /// @return Returns true for a successful deposit, false for an unsuccessful deposit
    function lockPayment(TxInfo calldata param)
        public
        payable
        override
        returns (bool)
    {
        require(param.minPayment > 0 && msg.value > param.minPayment);
        TxInfo storage t = txMap[param.id];
        t.owner = msg.sender;
        t.minPayment = param.minPayment;
        t.recipient = param.recipient;
        t.deadline = block.timestamp + lockTime;
        t.lockedFee = msg.value;
        t._isExisted = true;
    }

    /// @notice Returns the current allowance given to a spender by an owner
    /// @param txId transaction id
    /// @return Returns true for a successful payment, false for an unsuccessful payment
    function unlockPayment(string calldata txId, uint256 actualFee)
        public
        override
        returns (bool)
    {
        TxInfo storage t = txMap[txId];
        require(t._isExisted, "Transaction does not exist");
        require(
            t.owner == msg.sender || t.recipient == msg.sender,
            "Transaction has no connection wtih caller"
        );
        // if passed deadline, return payback to user
        if (block.timestamp > t.deadline) {
            payable(address(t.owner)).transfer(t.lockedFee);
        } else {
            if (actualFee < t.minPayment) {
                actualFee = t.minPayment;
            }
            payable(address(t.recipient)).transfer(actualFee);
            if (t.lockedFee > actualFee) {
                payable(address(t.owner)).transfer(t.lockedFee - actualFee);
            }
        }
        // todo: get status from oralce/other contract, status include status, real fee
        // check status, if not complete, return

        txMap[txId]._isExisted= false;
        return true;
        // real fee is greater than tx.fee, take tx.fee
        // real fee is less than tx.minPayment, take minPayment, return tx.fee - minPayment to tx.owner
        // otherwise, take real fee, return tx.fee - real fee to tx.owner
    }

    // function greet() public view returns (string memory) {
    //   return greeting;
    // }

    // function setGreeting(string memory _greeting) public {
    //   console.log("Changing greeting from '%s' to '%s'", greeting, _greeting);
    //   greeting = _greeting;
    // }
}
