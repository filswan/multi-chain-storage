// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");

const erc20ABI = require('../../artifacts/contracts/test/ERC20.sol/TestERC20.json').abi;

const one = "10000000000000";
const ten = "10000000000000000000";
const oneThousand = "1000000000000000000000";

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";

  const recipientAddress = "0xE53AEd6DEA9e44116D4551a93eEeE28bC8684916";

  const gatewayContractAddress = "0x80a186DCD922175019913b274568ab172F6E20b1";

  const cid = "eb1bf504-456e-4188-a5f7-932a9c36bef4QmcDsFG1ENSECBgPTE29kKFD9E1YVgtBrrybt1cWzY41qC";

  const [payer] = await ethers.getSigners();


  const USDCInstance = new ethers.Contract(usdcAddress, erc20ABI);

  await USDCInstance.connect(payer).approve(gatewayContractAddress, oneThousand);


  const contract = await hre.ethers.getContractFactory("SwanPayment");
  const paymentInstance = await contract.attach(gatewayContractAddress);

  // const tx = await paymentInstance.connect(payer).lockTokenPayment({
  //   id: cid,
  //   minPayment: one,
  //   amount: ten,
  //   lockTime: 60, // 6 days
  //   recipient: recipientAddress,
  //   size:100,
  //   copyLimit: 2,
  // }, overrides);

  // await tx.wait();

  const info = await paymentInstance.getLockedPaymentInfo(cid);

  console.log(info);

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
