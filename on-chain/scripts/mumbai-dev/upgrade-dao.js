// scripts/upgrade-box.js
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const [signer] = await ethers.getSigners();

  const contract = await ethers.getContractFactory("FilswanOracle");
  const oracleDAOContractAddress  = "0x9208C2B417Ec2699454843A06A5E49fA6dd88422";

  
  await upgrades.upgradeProxy(oracleDAOContractAddress, contract);

  console.log("FilswanOracle upgraded");

  // const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  // const tx = await daoOracleInstance.connect(signer).setFilinkOracle(
  //   "0xcE9A9e594db39dCD449E392d68F60959533c0D75"
  // );
  // await tx.wait();

  // console.log("Set Filink Oracle completed.");

  // const tx = await daoOracleInstance.connect(signer).updateThreshold(2);
  // await tx.wait();

  // console.log("update threshold completed.");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });