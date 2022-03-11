// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

// import "@openzeppelin/contracts/access/Ownable.sol";

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";

/**
 * Request testnet LINK and ETH here: https://faucets.chain.link/
 * Find information on LINK Token Contracts and get the latest ETH and LINK faucets here: https://docs.chain.link/docs/link-token-contracts/
 */

/**
 * THIS IS AN EXAMPLE CONTRACT WHICH USES HARDCODED VALUES FOR CLARITY.
 * PLEASE DO NOT USE THIS CODE IN PRODUCTION.
 */
contract FilinkConsumer is ChainlinkClient {
    using Chainlink for Chainlink.Request;
  
    mapping(bytes32 => string) private mapRequestDeal;
    mapping(string => uint256) private mapDealPrice;
    
    address private oracle;
    bytes32 private jobId;
    uint256 private fee;

    uint256 public price;

    constructor(address _chainlinkToken, address _oracle, uint256 _fee) {
        setChainlinkToken(_chainlinkToken);
        oracle = _oracle;
        jobId = "2bb15c3f9cfc4336b95012872ff05092";
        fee = _fee; // (Varies by network and job)
    }

    function concatenate(
        string memory s1,
        string memory s2
    ) private pure returns (string memory) {
        return string(abi.encodePacked(s1, s2));
    }
    
    /**
     * Create a Chainlink request to retrieve API response, find the target
     * data, then multiply by 1000000000000000000 (to remove decimal places from data).
     */

     // todo: use call back to pay
    function requestDealInfo(string calldata deal) public returns (bytes32 requestId) 
    {
        require(mapDealPrice[deal] == 0, "deal price is already on-chain, call getPrice(deal)");

        Chainlink.Request memory request = buildChainlinkRequest(jobId, address(this), this.fulfill.selector);

        string memory params = concatenate(deal, "?network=filecoin_mainnet");
        
        // Set the URL to perform the GET request on
        request.add("get", concatenate("http://35.168.51.2:7886/deal/", params));

        // request.add("deal", deal);

        request.add("path", "data.deal.storage_price");
      
        bytes32 id = sendChainlinkRequestTo(oracle, request, fee);
        mapRequestDeal[id] = deal;

        return id;
    }
    
    /**
     * Receive the response in the form of uint256
     */ 
    function fulfill( bytes32  _requestId, uint256 _price) public recordChainlinkFulfillment(_requestId)
    {
        price = _price;
        mapDealPrice[mapRequestDeal[_requestId]] = _price;
    }

    function getPrice(string calldata deal) public view returns (uint256)
    {
        return mapDealPrice[deal];
    }

    // function withdrawLINK(address _to, uint256 _amount) onlyowner public returns (bool)
    // {
    //     return transfer(address(this), _to, _amount);
    // }
}