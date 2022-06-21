
const hre = require("hardhat");

async function main() {

 
  const oracleDAOContractAddress = "0x6f83DA2C5f1C5AAC259aD8d817Bb92c2D863F74c";

  const addressList = [
    "0x800210CfB747992790245eA878D32F188d01a03A",
    "0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d",
    "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1"
  ];

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const result = await daoOracleInstance.getCarPaymentVotes('6976653', 'filecoin_mainnet', '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a');
  console.log(result);

  const result1 = await daoOracleInstance.getCidList('6976653', 'filecoin_mainnet');
  console.log(result1);

  const result2 = await daoOracleInstance.getSignatureList('6976653', 'filecoin_mainnet');
  console.log(result2);

  
  // const result2 = await daoOracleInstance.isCarPaymentAvailable('169017', 'filecoin_calibration', '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a');
  // console.log(result2);

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
