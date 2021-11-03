
const hre = require("hardhat");

async function main() {

  const [signer] = await ethers.getSigners();
 
  const oracleDAOContractAddress = "0xe3262c0848b0cc5cd43df7139103f1fbf26558cc";
  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const cid = "bafk2bzacedh6keeksywaoa3wjryqzihqixyfekqgfljfosrcoyaj2kozl767s";
  const orderId = "";
  const dealId = "0x0000000000000000000000000000000000000000000000000000000000000002";

  const paid = "10000000000000000"; // paid filcoins
  const recipient = "0x591f62C3FDC087dADC8A02dF76fD0a2Bd2168CDF";
  const terms = "2000000000000000"; 
  const status = true; // true for successful paid, false for failed paid

  const tx = await daoOracleInstance.connect(signer).signTransaction(
    cid,
    orderId,
    dealId,
    paid,
    recipient,
    terms,
    status
  );
  await tx.wait();

  console.log("Sign transaction completed.");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
