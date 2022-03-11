
const hre = require("hardhat");

async function main() {

  const [signer, signer2, signer3] = await ethers.getSigners();

  const oracleDAOContractAddress = "0x00233B4d7A9d84b9c6440015A287DE2c5436F5D3";
  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);


  const cid = "bafykbzacebrb2yrwvrznyp7ca3xidfjyv4sz7yzqys2ernrzdvuzcmx5aaiz6";
  const orderId = "";
  const dealId = "4227297";


  const paid = "10000000000000000"; // paid filcoins
  const recipient = "0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a";
  //const terms = "2000000000000000";
  const status = true; // true for successful paid, false for failed paid

  const tx = await daoOracleInstance.connect(signer3).signTransaction(
    cid,
    dealId,
    recipient
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
