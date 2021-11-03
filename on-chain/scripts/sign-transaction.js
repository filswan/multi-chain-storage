
const hre = require("hardhat");

async function main() {

  const [signer] = await ethers.getSigners();

  const oracleDAOContractAddress = "0xe3262c0848b0cc5cd43df7139103f1fbf26558cc";
  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const cid = "bafykbzacea7nzf7cr4olwxmjjvzszaho2mv34wwbs46yfr7fyegwmxee742b4";
  const orderId = "";
  const dealId = "bafyreifgrtipfm5u6asz7wuf6gx4747ukcx4nuahbzs3npg3y7hqux5sby";

  const paid = "10000000000000000"; // paid filcoins
  const recipient = "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1";
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
