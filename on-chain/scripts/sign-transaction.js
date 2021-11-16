
const hre = require("hardhat");

async function main() {

  const [signer] = await ethers.getSigners();

  const oracleDAOContractAddress = "0xe3262c0848b0cc5cd43df7139103f1fbf26558cc";
  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  const cid = "bafykbzacea4cz7kz77wx6zcqajd3fykb3zwmewiggy4x33wacuwads5favmf2";
  const orderId = "";
  const dealId = "'bafyreieqkjsrmwkqybqhxwkqzxsixphu6z4wmvrfj6ham2zf6doehabt2y'";

  const paid = "10000000000000000"; // paid filcoins
  const recipient = "0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a";
  //const terms = "2000000000000000";
  const status = true; // true for successful paid, false for failed paid

  const tx = await daoOracleInstance.connect(signer).signTransaction(
    cid,
    orderId,
    dealId,
    paid,
    recipient,
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
