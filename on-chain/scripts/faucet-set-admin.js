
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  const faucetContract = await ethers.getContractFactory("SwanFaucet");

  const faucetAddress = "0xd5b31FB565d608692d6422beB31Bf0875dad4fC3"

  const faucetInstance = await faucetContract.attach(faucetAddress);
  // const owner = await await faucetInstance.connect(deployer).owner();
  // console.log(owner);

  // const tx1 = await faucetInstance.connect(deployer).setAdmin("0xA878795d2C93985444f1e2A077FA324d59C759b0", overrides);
  // await tx1.wait();

  const tx2 = await faucetInstance.connect(deployer).setAdmin("0x8f9B5Fc53DC3035122D7e6f08289f4A927Bd064e", overrides);
  // await tx2.wait();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });