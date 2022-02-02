
const { ethers, upgrades } = require("hardhat");

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Executing contracts with the account:", deployer.address);

  const faucetContract = await ethers.getContractFactory("SwanFaucet");

  const faucetAddress = "0xd5b31FB565d608692d6422beB31Bf0875dad4fC3"
  const receiver = "0x591f62C3FDC087dADC8A02dF76fD0a2Bd2168CDF";

  // execute faucet
  const tokens = ["0xe11A86849d99F524cAC3E7A0Ec1241828e332C62", "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"];
  const amounts = ["100000000000000000000", "100000000000000000"];
  const faucetInstance = await faucetContract.attach(faucetAddress);
  const tx = await faucetInstance.connect(deployer).sendMultiTokens(tokens, amounts, receiver, overrides);
  await tx.wait();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });