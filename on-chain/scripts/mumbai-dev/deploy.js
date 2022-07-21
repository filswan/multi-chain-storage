
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  // const swanOracleContract = await ethers.getContractFactory("FilswanOracle");

  // // threshold 2, meaning 2 votes are required to pass
  // const swanOracleInstance = await upgrades.deployProxy(swanOracleContract, [deployer.address, 2], overrides);
 
  // await swanOracleInstance.deployed();
  // console.log(`swanOracleInstance address: ${swanOracleInstance.address}`);

  // const swanOracleAddress = swanOracleInstance.address;

  // const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";
  // const priceFeedOracleAddress = "0xe8a67994c114e0c17E1c135d0CB599a2394f1505";

  // // todo: ?? different contract for each flink service?
  // const chainlinkContract = await ethers.getContractFactory("FilinkConsumer");

  // const chainlinkAddress = "0x326C977E6efc84E512bB9C30f76E30c160eD06FB";
  // const httpOracleAddress = "0x0bDDCD124709aCBf9BB3F824EbC61C87019888bb";
  // const fee = "10000000000000000"

  // const chainlinkInstance = await chainlinkContract.deploy(chainlinkAddress, httpOracleAddress, fee);

  // console.log(`chainlinkInstance address: ${chainlinkInstance.address}`)

  // const flinkOracleAddress = chainlinkInstance.address;

  const swanOracleAddress="0xA12EB17A664E206f363bB240e01dbAa746d2f804";
  const flinkOracleAddress = "0xef4828525f78991a2b7b1f108751948F16f25a3F";
  const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";
  const priceFeedOracleAddress = "0xe8a67994c114e0c17E1c135d0CB599a2394f1505";

  const swanPaymentContract = await ethers.getContractFactory("SwanPayment");

  const swanPaymentInstance = await upgrades.deployProxy(swanPaymentContract, [deployer.address, usdcAddress, swanOracleAddress, priceFeedOracleAddress, flinkOracleAddress], overrides);
  await swanPaymentInstance.deployed();
  console.log(`swanPaymentInstance address: ${swanPaymentInstance.address}`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });