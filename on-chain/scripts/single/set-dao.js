
const hre = require("hardhat");

async function main() {

 // todo: change address
  const oracleDAOContractAddress = "0x00233B4d7A9d84b9c6440015A287DE2c5436F5D3";

  // todo: update addresses
  const addressList = [
    "0x800210CfB747992790245eA878D32F188d01a03A",
    "0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d",
    "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1",
    "0x591f62C3FDC087dADC8A02dF76fD0a2Bd2168CDF",
    "0xE41c53Eb9fce0AC9D204d4F361e28a8f28559D54"
  ];

  

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const tx = await daoOracleInstance.connect(signer).setDAOUsers(addressList);
  await tx.wait();

  const role = await daoOracleInstance.DAO_ROLE.call();
  console.log(role);

  const hasRole = await daoOracleInstance.connect(signer).hasRole(
    role,
    addressList[0]);

  console.log(hasRole);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
