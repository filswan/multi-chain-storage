// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;


import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./MCSCollection.sol";

contract CollectionFactory is Initializable, OwnableUpgradeable {
    mapping(address => address[]) userToCollections;

    event CreateCollection(address collectionOwner, address collectionAddress);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
    }

    function createCollection(string memory contractURI) public returns (address) {
        // user will be owner AND admin, factory will be admin
        MCSCollection newCollection = new MCSCollection(contractURI); // this contract will be owner and admin
        newCollection.setAdmin(msg.sender, true); // set the user as another admin
        newCollection.transferOwnership(msg.sender); // then set user as admin

        userToCollections[msg.sender].push(address(newCollection));

        emit CreateCollection(msg.sender, address(newCollection));

        return address(newCollection);
    }

    //TODO: mint functions
    function mint(address collection, address recipient, uint amount, string memory uri) public {
        require(MCSCollection(collection).isAdmin(msg.sender), 'caller is not admin');

        MCSCollection(collection).mint(recipient, amount, uri, "");
    }

    function mintToNewCollection(string memory contractURI, address recipient, uint amount, string memory uri) public {
        address collection = createCollection(contractURI);
        mint(collection, recipient, amount, uri);
    }

    // getter functions
    function getCollections(address user) public view returns (address[] memory) {
        return userToCollections[user];
    }
}