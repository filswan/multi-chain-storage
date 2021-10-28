
const hre = require("hardhat");

async function main() {

  const recipientAddress = "0xE53AEd6DEA9e44116D4551a93eEeE28bC8684916";
  const oracleContractAddress = "0xf3963AF369974F557e987A1B5CD327413B2608BE";

  const cid = "bafk2bzaceajwszlar4obpnmifoydefxlhfd7npbnmq3onkfzkincyy4fdj5xk";

  const realFee = ethers.utils.parseEther("0.03");

  const [daoUser] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilecoinOracle");
  const oracleInstance = await contract.attach(oracleContractAddress);


  const tx = await oracleInstance.connect(daoUser).updatePaymentInfo(cid, realFee);
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
