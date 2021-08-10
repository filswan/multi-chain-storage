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

  const recipientAddress = "0xEC9Ee466a7372A4D49f64AAF4d348dBE9b5DE02A";
  const gatewayContractAddress = "0xad8cE271beE7b917F2a1870C8b64EDfF6aAF3342";

  const cid = "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72";

  const minPay10Native = ethers.utils.parseEther("0.011");

  const [payer] = await ethers.getSigners();

  const fee = {
    // To convert Ether to Wei:
    value: ethers.utils.parseEther("0.021")     // ether in this case MUST be a string
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
