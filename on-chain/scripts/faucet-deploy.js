
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  const faucetContract = await ethers.getContractFactory("SwanFaucet");

  // const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";

  // const faucetInstance = await upgrades.deployProxy(faucetContract, [deployer.address], overrides);
  // await faucetInstance.deployed();
  // console.log(`faucetInstance address: ${faucetInstance.address}`)

  // upgrade
  const faucetAddress  = "0xd5b31FB565d608692d6422beB31Bf0875dad4fC3";
  await upgrades.upgradeProxy(faucetAddress, faucetContract);

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });