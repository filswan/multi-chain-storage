// scripts/upgrade-box.js
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const swanPaymentContract = await ethers.getContractFactory("SwanPayment");
  // old address 0xABeAAb124e6b52afFF504DB71bbF08D0A768D053
  const swanPaymentAddress  = "0x7ab09f9Ab4D39cfBE0551dfb6AdAc63C89bB955b";
  
  await upgrades.upgradeProxy(swanPaymentAddress, swanPaymentContract);

  console.log("swanPayment upgraded");

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });