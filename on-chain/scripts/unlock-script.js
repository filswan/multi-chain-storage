// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";

  const recipientAddress = "0x591f62C3FDC087dADC8A02dF76fD0a2Bd2168CDF";

  const gatewayContractAddress = "0x12EDC75CE16d778Dc450960d5f1a744477ee49a0";

  const cid = "abcd2bzacedh6keeksywaoa3wjryqzihqixyfekqgfljfosrcoyaj2";

  const [payer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("SwanPayment");
  const paymentInstance = await contract.attach(gatewayContractAddress);

  const tx = await paymentInstance.connect(payer).unlockTokenPayment({
    id: cid,
    orderId: "",
    dealId: "4109",
    amount: "0",
    recipient: recipientAddress, //todo:
  }, overrides);

  await tx.wait();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
