//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "hardhat/console.sol";

contract FilecoinOracle is Initializable {
    struct TxOracleStatus {
        uint256 actualPaid;
        uint256 timestamp;
    }

    address private _owner;
    address[] private _daoUsers;
    mapping(string => mapping(address => TxOracleStatus)) statusMap;

    uint8 private _threshold; // threhold to agree off-chain payment info

    function initialize(address owner, uint8 threshold) public initializer {
        _owner = owner;
        _threshold = threshold;
    }

    constructor(address owner) {
        _owner = owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(_owner == msg.sender, "Caller is not the owner");
        _;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyDAOUser() {
        bool found = false;
        for (uint8 i = 0; i < _daoUsers.length; i++) {
            if (_daoUsers[i] == msg.sender) {
                found = true;
                break;
            }
        }
        require(found, "Caller is not the DAO user");
        _;
    }

    function setDAOUsers(address[] calldata daoUsers)
        public
        onlyOwner
        returns (bool)
    {
        _daoUsers = daoUsers;
        return true;
    }

    // /**
    //  * @dev Throws if called by any account other than the owner.
    //  */
    // modifier onlyParticipant() {
    //     require(_owner == msg.sender, "Caller is not the owner");
    //     _;
    // }

    function getPaymentInfo(string calldata txId)
        public
        view
        returns (uint256 actualPaid)
    {
        // default value is 0
        // todo: every oracle should update the same payment info.
        mapping(address => TxOracleStatus) storage status = statusMap[txId];
        for (uint8 i = 0; i < _daoUsers.length; i++) {
            if (status[_daoUsers[i]].actualPaid == 0) {
                return 0;
            }
        }

        return status[_daoUsers[0]].actualPaid;
    }

    // todo: add only DAO group users
    function updatePaymentInfo(string calldata txId, uint256 actualPaid)
        public
        onlyDAOUser
        returns (bool)
    {
        // default value is 0
        if (statusMap[txId][msg.sender].timestamp == 0) {
            statusMap[txId][msg.sender].timestamp = block.timestamp;
            statusMap[txId][msg.sender].actualPaid = actualPaid;
        }
        return true;
    }
}
