// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity >=0.8.4;

/// @title Minimal ERC20 interface for Uniswap
/// @notice Contains a subset of the full ERC20 interface that is used in Uniswap V3
interface IPaymentMinimal {

    struct lockPaymentParam{
        string id; // in filecoin, is CID
        uint256 minPayment;
        uint256 lockTime;
        address recipient;
    }
    struct TxInfo {
        string id; // in filecoin, is CID
        uint256 minPayment;
        uint256 lockedFee;
        address owner;
        address recipient;
        uint256 deadline;
        bool _isExisted;
    }

    /// @notice Returns the locked balance of caller in specific transaction
    /// @param txId The transaction for which to look up the balance it has
    /// @return tx The transaction object
    function getLockedPaymentInfo(string calldata txId) external view returns (TxInfo memory tx);

    /// @notice Deposits the amount of token for specific transaction
    /// @param txInfo The transaction information for which to deposit balance
    /// @return Returns true for a successful deposit, false for an unsuccessful deposit
    function lockPayment(lockPaymentParam calldata txInfo) external payable returns (bool);

    /// @notice Returns the current allowance given to a spender by an owner
    /// @param txId transaction id
    /// @return Returns true for a successful payment, false for an unsuccessful payment
    function unlockPayment(string calldata txId) external returns (bool);

    //function updateTxStatus((bytes calldata txId, uint256 cost, uint8 status) external returns (bool);
}