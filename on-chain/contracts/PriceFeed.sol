//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.4;

// import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "hardhat/console.sol";
import '@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol';
import "./interfaces/IPriceFeed.sol";

contract PriceFeed is IPriceFeed {

    using SafeMath for uint112;
    using SafeMath for uint256;

    address private _dexPair;

    uint8 private _tokenIndex; // wfil token address index in the pair

    constructor(address dexPair, uint8 tokenIndex) public {

        _dexPair = dexPair;
        _tokenIndex = tokenIndex;
    }

    // how many tokens of tokenIput are equal amount of wfil
    function consult(
        address tokenInput,
        uint256 amount
    ) external 
    override
    view returns (uint256){
        require(amount>0, "amount must greater than 0");
        (uint112 _reserve0, uint112 _reserve1, ) = IUniswapV2Pair(_dexPair).getReserves();

        uint256 cp = uint(_reserve0).mul(_reserve1);
        uint256 retAmount = 0;

        if(_tokenIndex == 0){
            uint256 tAmt = _reserve0.sub(amount);
            require(tAmt>0, "not enough token to return");
            uint256 amt = cp.div(tAmt);
            retAmount = amt.sub(_reserve1);
        }else if(_tokenIndex == 1){
            uint256 tAmt = _reserve1.sub(amount);
            require(tAmt>0, "not enough token to return");
            uint256 amt = cp.div(tAmt);
            retAmount = amt.sub(_reserve0);
        }
        return retAmount;
    }
}