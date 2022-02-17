
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const filinkAddress = "";

  const [deployer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const tx = await daoOracleInstance.connect(signer).setFilinkOracle(filinkAddress);
  await tx.wait();


  const paymentContract = await hre.ethers.getContractFactory("SwanPayment");
  const paymentInstance = await paymentContract.attach(oracleDAOContractAddress);

  tx = await paymentInstance.connect(signer).setChainlinkOracle(filinkAddress);
  await tx.wait();

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });