// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "../interfaces/IPriceFeed.sol";

interface IERC20 {
    function decimals() external view returns (uint8);
}

contract ChainlinkPriceFeed is IPriceFeed {

    AggregatorV3Interface internal priceFeed;

    /**
     * Network: Polygon
     * Aggregator: FIL/USD
     * Address: 0xa07703E5C2eD1516107c7c72A494493Dcb99C676
     */
    constructor() {
        priceFeed = AggregatorV3Interface(0xa07703E5C2eD1516107c7c72A494493Dcb99C676);
    }

    /**
     * Returns the latest price
     */
    function getLatestPrice() public view returns (int) {
        (
            /*uint80 roundID*/,
            int price,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = priceFeed.latestRoundData();
        return price;
    }

    function consult(
        address token,
        uint256 amount
    ) external view override returns (uint) {
        uint price = uint(getLatestPrice());
        uint amountInWei = price * amount / 10 ** decimals();
        uint amountInTokenDecimals = amountInWei / 10 ** (18 - IERC20(token).decimals());

        return amountInTokenDecimals;
    }

    function decimals() public view returns (uint8) {
        return priceFeed.decimals();
    }

    function description() public view returns (string memory) {
        return priceFeed.description();
    }

    function version() public view returns (uint) {
        return priceFeed.version();
    }
}