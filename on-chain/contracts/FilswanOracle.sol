//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "hardhat/console.sol";

contract FilswanOracle is OwnableUpgradeable, AccessControlUpgradeable {

    bytes32 public constant DAO_ROLE = keccak256("DAO_ROLE");

    // todo: add get threshold function
    uint8 private _threshold;

    mapping(string => mapping(address => TxOracleInfo)) txInfoMap;
    mapping(bytes32 => uint8) txVoteMap;
    struct TxOracleInfo {
        uint256 paid;
        uint256 terms;
        address recipient;
        bool status;
        bool flag; // check existence of signature
    }

    event SignTransaction(
        string cid,
        string orderId,
        string dealId,
        address recipient,
        uint256 paid,
        uint256 terms,
        bool status
    );

    function initialize(address admin, uint8 threshold) public initializer {
        __Ownable_init();
        __AccessControl_init();
        _setupRole(DEFAULT_ADMIN_ROLE, admin);
        _threshold = threshold;
    }

    function updateThreshold(uint8 threshold)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (bool)
    {
        _threshold = threshold;
        return true;
    }

    function updateThreshold(uint8 threshold)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (bool)
    {
        _threshold = threshold;
        return true;
    }

    function setDAOUsers(address[] calldata daoUsers)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (bool)
    {
        for (uint8 i = 0; i < daoUsers.length; i++) {
            grantRole(DAO_ROLE, daoUsers[i]);
        }
        return true;
    }

     function concatenate(
        string memory s1,
        string memory s2,
        string memory s3
    ) private pure returns (string memory) {
        return string(abi.encodePacked(s1, s2, s3));
    }

    function signTransaction(
        string memory cid,
        string memory orderId,
        string memory dealId,
        uint256 paid,
        address recipient,
        uint256 terms, // todo: term is for updated 
        bool status
    ) public onlyRole(DAO_ROLE) {
        string memory key = concatenate(cid, orderId, dealId);

        // todo: improvement, DAO only can sign once for each transaction 
        require(
            txInfoMap[key][msg.sender].flag == false,
            "You already sign this transaction"
        );


        txInfoMap[key][msg.sender].recipient = recipient;
        txInfoMap[key][msg.sender].paid = paid;
        txInfoMap[key][msg.sender].terms = terms;
        txInfoMap[key][msg.sender].status = status;
        txInfoMap[key][msg.sender].flag = true;

        bytes32 voteKey = keccak256(
            abi.encodePacked(cid, orderId, dealId, paid, recipient, status)
        );

        txVoteMap[voteKey] = txVoteMap[voteKey] + 1;

        emit SignTransaction(
            cid,
            orderId,
            dealId,
            recipient,
            paid,
            terms,
            status // bool
        );
    }

    function isPaymentAvailable(
        string memory cid,
        string memory orderId,
        string memory dealId,
        uint256 paid,
        address recipient,
        bool status
    ) public view returns (bool) {
        bytes32 voteKey = keccak256(
            abi.encodePacked(cid, orderId, dealId, paid, recipient, status)
        );
        return txVoteMap[voteKey] >= _threshold;
    }

}
