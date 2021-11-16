// scripts/upgrade-box.js
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const contract = await ethers.getContractFactory("FilswanOracle");
  const oracleDAOContractAddress  = "0xE3262C0848B0cc5cD43df7139103f1Fbf26558cc";

  await upgrades.upgradeProxy(oracleDAOContractAddress, contract);

  console.log("FilswanOracle upgraded");

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });