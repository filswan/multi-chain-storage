//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "./FilinkConsumer.sol";

contract FilswanOracle is OwnableUpgradeable, AccessControlUpgradeable {

    bytes32 public constant DAO_ROLE = keccak256("DAO_ROLE");

    uint8 private _threshold;

    mapping(string => mapping(address => TxOracleInfo)) txInfoMap;
    mapping(bytes32 => uint8) txVoteMap;

    address private _filinkAddress;
    mapping(string => string[]) cidListMap;

    struct TxOracleInfo {
        uint256 paid;
        uint256 terms;
        address recipient;
        bool status;
        bool flag; // check existence of signature
    }

    event SignTransaction(
        string[] cidList,
        string dealId,
        address recipient
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

    function setFilinkOracle(address filinkAddress)public
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (bool){
       _filinkAddress = filinkAddress;
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
        string memory s2
    ) private pure returns (string memory) {
        return string(abi.encodePacked(s1, s2));
    }

    function signTransaction(
        string[] memory cidList,
        string memory dealId,
        address recipient
    ) public onlyRole(DAO_ROLE) {
        string memory key = dealId;

        require(
            txInfoMap[key][msg.sender].flag == false,
            "You already sign this transaction"
        );

        txInfoMap[key][msg.sender].recipient = recipient;
        txInfoMap[key][msg.sender].flag = true;

        bytes32 voteKey = keccak256(
            abi.encodeWithSignature("f(string,address,string[])",dealId, recipient, cidList)
        );

        txVoteMap[voteKey] = txVoteMap[voteKey] + 1;

        if(txVoteMap[voteKey] == _threshold && _filinkAddress != address(0)){
            cidListMap[key] = cidList;
            FilinkConsumer(_filinkAddress).requestDealInfo(dealId);
        }

        emit SignTransaction(
            cidList,
            dealId,
            recipient
        );
    }

    function isPaymentAvailable(
        string memory dealId,
        address recipient
    ) public view returns (bool) {
        bytes32 voteKey = keccak256(
            abi.encodePacked(dealId, recipient)
        );
        return txVoteMap[voteKey] >= _threshold;
    }


    function getThreshold() public view returns (uint8) {
        return _threshold;
    }

    function getCidList(string memory dealId) public view returns (string[] memory) {
         return cidListMap[dealId];
    }
}
