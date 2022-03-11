
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const filinkAddress = "0x917347c86B6bd1536d125d1b2FE8687F9fa3091F";

  const oracleDAOContractAddress = "0x00233B4d7A9d84b9c6440015A287DE2c5436F5D3";

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const tx = await daoOracleInstance.connect(signer).setFilinkOracle(filinkAddress);
  await tx.wait();

  const swanPaymentAddress  = "0x7ab09f9Ab4D39cfBE0551dfb6AdAc63C89bB955b";
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