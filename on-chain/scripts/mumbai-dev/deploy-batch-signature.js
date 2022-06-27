
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const addressList = [
    "0x6d2e5279b106843f6E924194401B50e6e27FE12a",
    "0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5",
    "0xeA2bf08288bbfB0d3DBf534f35af32bF2c6E5e45",
    "0x800210cfb747992790245ea878d32f188d01a03a",
    "0x21aE11DF412002378b73A28EF137FBfC59332BA4",
    "0x71632B0e6b5347BAc09E85a40B329397af473933",
    "0x591f62C3FDC087dADC8A02dF76fD0a2Bd2168CDF"
  ];

  const [deployer, k1, k2, k3] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  const swanOracleContract = await ethers.getContractFactory("FilswanOracle");

  // threshold 2, meaning 2 votes are required to pass
  // const swanOracleInstance = await upgrades.deployProxy(swanOracleContract, [deployer.address, 2], overrides);
 
  // await swanOracleInstance.deployed();
  // console.log(`swanOracleInstance address: ${swanOracleInstance.address}`);

  const swanOracleInstance = await swanOracleContract.attach("0xA12EB17A664E206f363bB240e01dbAa746d2f804");

  // console.log("Setting filink address");
  // const filinkAddress = "0xef4828525f78991a2b7b1f108751948F16f25a3F";
  // const tx = await swanOracleInstance.connect(deployer).setFilinkOracle(filinkAddress);
  // await tx.wait();

  // console.log("Setting dao address");

  // const tx1 = await swanOracleInstance.connect(deployer).setDAOUsers(addressList);
  // await tx1.wait();

  console.log("do presign..");
  const txPresign = await swanOracleInstance.connect(k3).preSign("10001", "filecoin_calibration", "0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a", 1);
  await txPresign.wait();

  // console.log("Setting complete");

  // run this one later.....
  // const gatewayContractAddress = "0x80a186DCD922175019913b274568ab172F6E20b1";

  // const contract = await hre.ethers.getContractFactory("SwanPayment");
  // const paymentInstance = await contract.attach(gatewayContractAddress);

  // await paymentInstance.setOracle(swanOracleInstance.address);

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });