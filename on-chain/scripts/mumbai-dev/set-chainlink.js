
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {
  const oracleDAOContractAddress = "0x9208C2B417Ec2699454843A06A5E49fA6dd88422";

  const filinkAddress = "0xB6312D719F6B496647703c81F6965EF38bF58B8D";

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const tx = await daoOracleInstance.connect(signer).setFilinkOracle(filinkAddress);
  await tx.wait();

  const swanPaymentAddress  = "0x24B9c56BB6419f4c5AE6a63Fd64dE0dCFA1841F1";
  const paymentContract = await hre.ethers.getContractFactory("SwanPayment");
  const paymentInstance = await paymentContract.attach(swanPaymentAddress);

  const tx2 = await paymentInstance.connect(signer).setChainlinkOracle(filinkAddress);
  await tx2.wait();

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });