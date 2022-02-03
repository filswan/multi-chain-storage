
const { ethers, upgrades } = require("hardhat");

const sleep = require('sleep-promise');

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Executing contracts with the account:", deployer.address);

  const vrfContract = await ethers.getContractFactory("SwanVRFConsumer");

  const address = "0x8A61dEff846240ce57a36b0D83391260cD794db1"

  const vrfInstance = await vrfContract.attach(address);
  const requestId = await vrfInstance.connect(deployer).callStatic.getRandomNumber();
  const tx = await vrfInstance.connect(deployer).getRandomNumber();
  await tx.wait();
  
  console.log(`random number requestId is: ${requestId}`);

  console.log("Waiting for generating random number (30-120 seconds)......");

  let number = 0;
  let count = 0;
  while(number == 0 && count++ < 12)
  {
    process.stdout.write(`waiting seconds: ${count}0...\r`)
    await sleep(10000); 
    number = await vrfInstance.requestIdToRandomNumber(requestId);
  }
  console.log(`random number is:${number}`)
  
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });