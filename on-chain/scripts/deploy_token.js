
const { ethers, upgrades } = require("hardhat");

const {MaxUint256} = require("ethers");

// const {
//   abi: SWAP_ROUTER_ABI,
//   bytecode: SWAP_ROUTER_BYTECODE,
//  } = require('@uniswap/v3-periphery/artifacts/contracts/SwapRouter.sol/SwapRouter.json');

const erc20ABI = require('@openzeppelin/contracts/build/contracts/ERC20.json').abi;

const uniswapv2FactoryABI = require('@uniswap/v2-core/build/IUniswapV2Factory.json').abi;

const uniswapv2RouterABI = require('@uniswap/v2-periphery/build/IUniswapV2Router02.json').abi;


const uniswapv2LibABI = require('@uniswap/v2-periphery/build/UniswapV2Library.json').abi;

const sushiswapFactoryAddress = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4";
const sushiswapRouterAddresses = "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506";

const oneThousand = "1000000000000000000000";
const oneHundred = "1000000000000000000000";
const tenThousand = "1000000000000000000000";
const oneMillion = "1000000000000000000000000";

const deadline = "1000000000000000000";

const overrides = {
  gasLimit: 9999999
}

async function main() {

    const [deployer] = await ethers.getSigners();  

    console.log("deployer: ", deployer.address);

    // const testTokenContract = await ethers.getContractFactory("TestERC20");

    // const USDCInstance = await testTokenContract.deploy( 
    //   "USDC",
    //   "USDC"
    //   );
    // await USDCInstance.deployed();

    // console.log(`USDCInstance address: ${USDCInstance.address}`);

    // await USDCInstance.mint(deployer.address, oneMillion);

    // const wFilInstance = await testTokenContract.deploy( 
    //   "wFil",
    //   "wFil"
    //   );
    // await wFilInstance.deployed();

    // console.log(`wFilInstance address: ${wFilInstance.address}`);
    // await wFilInstance.mint(deployer.address, oneThousand);

    const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";
    const wFilAddress = "0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055";


    // const usdcContract = new ethers.Contract(usdcAddress, erc20ABI);
    // await usdcContract.connect(deployer).approve(sushiswapRouterAddresses, oneMillion);

    // const wFilContract = new ethers.Contract(wFilAddress, erc20ABI);
    // await wFilContract.connect(deployer).approve(sushiswapRouterAddresses, oneMillion);


    // const sushiRouterContract = new ethers.Contract(sushiswapRouterAddresses, uniswapv2RouterABI);
    // await sushiRouterContract.connect(deployer).addLiquidity(usdcAddress, wFilAddress, oneMillion, oneThousand, tenThousand, oneHundred, deployer.address, deadline, overrides);

    const sushiFactoryContract = new ethers.Contract(sushiswapFactoryAddress, uniswapv2FactoryABI);
    const pairAddress = await sushiFactoryContract.connect(deployer).getPair(usdcAddress, wFilAddress);

    if(pairAddress === "0x0000000000000000000000000000000000000000") {
      console.log("Pair not found");
    }else {
      console.log(`Pair address: ${pairAddress}`);
    }


  }

  
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });