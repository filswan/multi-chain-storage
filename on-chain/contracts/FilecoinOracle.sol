//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "hardhat/console.sol";

contract FilecoinOracle {

    struct TxOracleStatus {
        uint256 actualPaid;
        uint256 timestamp;
    }


    address private _owner;
    address[] private  _daoUsers;
    mapping(string => mapping(address => TxOracleStatus)) statusMap;

    constructor() {
        console.log("Deploying Filecoin Oracle contract");
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
         for(uint8 i = 0; i < daoUsers.length; i++){
            if(daoUsers[i] == msg.sender){
                found = true;
                break;
            }
        }
        require(found, "Caller is not the DAO user");
        _;
    }

    function setDAOUsers(address[] daoUsers) public
        onlyOwner
        returns (bool){
            _daoUsers =  daoUsers;
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
        override
        returns (uint256 actualPaid)
    {
        // default value is 0
        mapping(address => TxOracleStatus) memory status = txMap[txId];
        for(uint8 i = 0; i < daoUsers.length; i++){
            if(status[daoUsers[i]].actualPaid == 0){
                return 0;
            }
        }
        return status[daoUsers[0]].actualPaid;
    }

    // todo: add only DAO group users
    function updatePaymentInfo(string calldata txId, uint256 actualPaid)
        public
        onlyDAO
        returns (bool)
    {
        // default value is 0
        if(txMap[txId][msg.sender].timestamp == 0)
        {
            txMap[txId][msg.sender].timestamp = block.timestamp;
            txMap[txId][msg.sender].actualPaid = actualPaid;
        }
    }

    
}
