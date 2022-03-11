
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  const vrfContract = await ethers.getContractFactory("SwanVRFConsumer");
  // const chainlinkInstance = await chainlinkContract.deploy();
  // await chainlinkInstance.deployed();

  const chainlinkAddress = "0x326C977E6efc84E512bB9C30f76E30c160eD06FB";
  const coordinator = "0x8C7382F9D8f56b33781fE506E897a4F1e2d17255";
  const keyHash = ethers.utils.hexZeroPad("0x6e75b569a01ef56d18cab6a8e71e6600d6ce853834d4a5748b720d06f878b3a4",32);
  const fee = "100000000000000" // 0.0001 in mumbai network

  const vrfInstance = await vrfContract.deploy(coordinator, chainlinkAddress, keyHash, fee);

  console.log(`vrfInstance address: ${vrfInstance.address}`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });