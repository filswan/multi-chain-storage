// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

const erc20ABI = require('../../artifacts/contracts/test/ERC20.sol/TestERC20.json').abi;

const one = "10000000000000";
const ten = "10000000000000000000";
const oneThousand = "1000000000000000000000";

const overrides = {
  gasLimit: 9999999
}

async function main() {


  const contract = await hre.ethers.getContractFactory("PriceFeed");
  const priceInstance = await contract.attach("0xe8a67994c114e0c17E1c135d0CB599a2394f1505");

  const info = await priceInstance.consult("0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055", one);

  console.log(info);

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
