
const hre = require("hardhat");

async function main() {

 
  const oracleDAOContractAddress = "0xe3262c0848b0cc5cd43df7139103f1fbf26558cc";

  const addressList = [
    "0x800210CfB747992790245eA878D32F188d01a03A",
    "0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d",
    "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1"
  ];

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const result = await daoOracleInstance.getCarPaymentVotes('87204','0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a');
  console.log(result);

  // const hasRole = await daoOracleInstance.connect(signer).hasRole(
  //   role,
  //   addressList[0]);

  // console.log(hasRole);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
