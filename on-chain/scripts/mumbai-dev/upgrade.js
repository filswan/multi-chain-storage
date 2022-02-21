// scripts/upgrade-box.js
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const swanPaymentContract = await ethers.getContractFactory("SwanPayment");
  // old address 0xABeAAb124e6b52afFF504DB71bbF08D0A768D053
  const swanPaymentAddress  = "0x24B9c56BB6419f4c5AE6a63Fd64dE0dCFA1841F1";
  
  await upgrades.upgradeProxy(swanPaymentAddress, swanPaymentContract);

  console.log("swanPayment upgraded");

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });