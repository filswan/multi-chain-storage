// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  const recipientAddress = "0xE53AEd6DEA9e44116D4551a93eEeE28bC8684916";
  const gatewayContractAddress = "0x5210ED929B5BEdBFFBA2F6b9A0b1B608eEAb3aa0";

  const cid = "bafk2bzaceb7cp727fxdrzudlgvsoivdwscydp35eb6wc3bzuflggfdhfa4rfe";

  const minPay10Native = ethers.utils.parseEther("0.03");

  const [payer] = await ethers.getSigners();

  const fee = {
    // To convert Ether to Wei:
    value: ethers.utils.parseEther("0.05")     // ether in this case MUST be a string
  };

  const contract = await hre.ethers.getContractFactory("SwanPayment");
  const paymentInstance = await contract.attach(gatewayContractAddress);

  const tx = await paymentInstance.connect(payer).lockPayment({
    id: cid,
    minPayment: minPay10Native,
    lockTime: 86400, // one day
    recipient: recipientAddress, //todo:
  }, fee);

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
