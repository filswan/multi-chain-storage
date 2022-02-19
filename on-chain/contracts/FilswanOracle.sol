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
        string[] cidList;
    }

    event SignTransaction(string cid, string dealId, address recipient);

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

    function setFilinkOracle(address filinkAddress)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (bool)
    {
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

    function concatenate(string memory s1, string memory s2)
        private
        pure
        returns (string memory)
    {
        return string(abi.encodePacked(s1, s2));
    }

    function signCarTransaction(
        string[] memory cidList,
        string memory dealId,
        string memory network,
        address recipient
    ) public onlyRole(DAO_ROLE) {
        string memory key = concatenate(dealId, network);

        require(
            txInfoMap[key][msg.sender].flag == false,
            "You already sign this transaction"
        );

        txInfoMap[key][msg.sender].recipient = recipient;
        txInfoMap[key][msg.sender].flag = true;
        txInfoMap[key][msg.sender].cidList = cidList;

        bytes32 voteKey = keccak256(
            abi.encodeWithSignature(
                "f(string, string,address,string[])",
                dealId,
                network,
                recipient,
                cidList
            )
        );

        txVoteMap[voteKey] = txVoteMap[voteKey] + 1;

        // todo: check cidList each time.
        if (txVoteMap[voteKey] == _threshold && _filinkAddress != address(0)) {
            cidListMap[key] = cidList;
            FilinkConsumer(_filinkAddress).requestDealInfo(dealId, network);
        }
    }

    function isCarPaymentAvailable(string memory dealId, string memory network, address recipient)
        public
        view
        returns (bool)
    {
        string[] memory cidList = cidListMap[dealId];
        bytes32 voteKey = keccak256(
            abi.encodeWithSignature(
                "f(string,string,address,string[])",
                dealId,
                network,
                recipient,
                cidList
            )
        );
        return txVoteMap[voteKey] >= _threshold;
    }

    function getCarPaymentVotes(string memory dealId, string memory network, address recipient)
        public
        view
        returns (uint8)
    {
        string[] memory cidList = cidListMap[dealId];
        bytes32 voteKey = keccak256(
            abi.encodeWithSignature(
                "f(string,string,address,string[])",
                dealId,
                network,
                recipient,
                cidList
            )
        );
        return txVoteMap[voteKey];
    }

    function getThreshold() public view returns (uint8) {
        return _threshold;
    }

    function getCidList(string memory dealId, string memory network)
        public
        view
        returns (string[] memory)
    {
        string memory key = concatenate(dealId, network);
        return cidListMap[key];
    }

    // function getSignatureList(
    //     string memory dealId,
    //     string memory network
    // ) public view returns (TxOracleInfo[] memory) {
    //     string memory key = concatenate(dealId, network);
    //     uint256 cnt = getRoleMemberCount(DAO_ROLE);
    //     TxOracleInfo[] memory result = new TxOracleInfo[](cnt);
    //     for (uint256 i = 0; i < cnt; i++) {
    //         address member = getRoleMember(DAO_ROLE, i);
    //         result[i] = txInfoMap[key][member];
    //     }
    //     return result;
    // }
}
