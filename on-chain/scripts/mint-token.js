
const { ethers, upgrades } = require("hardhat");

const { MaxUint256 } = require("ethers");

const erc20ABI = require('../artifacts/contracts/test/ERC20.sol/TestERC20.json').abi;

const uniswapv2FactoryABI = require('@uniswap/v2-core/build/IUniswapV2Factory.json').abi;

const uniswapv2RouterABI = require('@uniswap/v2-periphery/build/IUniswapV2Router02.json').abi;


const uniswapv2LibABI = require('@uniswap/v2-periphery/build/UniswapV2Library.json').abi;

const sushiswapFactoryAddress = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4";
const sushiswapRouterAddresses = "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506";

const oneThousand = "1000000000000000000000";
const oneHundred = "1000000000000000000000";
const tenThousand = "10000000000000000000000";
const oneMillion = "1000000000000000000000000";
const fiveMillion = "5000000000000000000000000";

const deadline = "1000000000000000000";

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const [k1, deployer] = await ethers.getSigners();

  console.log("deployer: ", deployer.address);

  const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";
  const wFilAddress = "0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055";

  const addressList = [
    // "0x800210CfB747992790245eA878D32F188d01a03A",
    // "0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d",
    // "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1",
    "0xE41c53Eb9fce0AC9D204d4F361e28a8f28559D54"
  ];

  // const USDCInstance = new ethers.Contract(usdcAddress, erc20ABI);
  const wFilInstance = new ethers.Contract(wFilAddress, erc20ABI);

  // addressList.forEach(async (address) => {
  //   console.log("address: ", address);
  //   await USDCInstance.connect(deployer).mint(address, tenThousand);
  //   await wFilInstance.connect(deployer).mint(address, oneHundred);
  // });

  // await USDCInstance.connect(deployer).mint(addressList[0], fiveMillion);
  const tx = await wFilInstance.connect(deployer).mint(addressList[0], oneMillion);
  await tx.wait();

  // await USDCInstance.connect(deployer).mint(addressList[1], tenThousand);
  // await wFilInstance.connect(deployer).mint(addressList[1], oneHundred);
  
  // await USDCInstance.connect(deployer).mint(addressList[2], tenThousand);
  // await wFilInstance.connect(deployer).mint(addressList[2], oneHundred);

  console.log("complete mint.")
}


main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
  