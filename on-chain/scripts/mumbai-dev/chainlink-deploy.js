
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  const chainlinkContract = await ethers.getContractFactory("FilinkConsumer");
  // const chainlinkInstance = await chainlinkContract.deploy();
  // await chainlinkInstance.deployed();

  const chainlinkAddress = "0x326C977E6efc84E512bB9C30f76E30c160eD06FB";
  const httpOracleAddress = "0x0bDDCD124709aCBf9BB3F824EbC61C87019888bb";
  const jobId = ethers.utils.hexZeroPad("0xc6a006e4f4844754a6524445acde84a0",32);
  const fee = "10000000000000000"

  const chainlinkInstance = await chainlinkContract.deploy(chainlinkAddress, httpOracleAddress, fee);

  console.log(`chainlinkInstance address: ${chainlinkInstance.address}`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });