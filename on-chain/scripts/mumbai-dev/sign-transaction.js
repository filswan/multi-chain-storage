
const hre = require("hardhat");

async function main() {

    // todo: change to singers
  const [signer, signer2, signer3] = await ethers.getSigners();

  const oracleDAOContractAddress = "0x6f83DA2C5f1C5AAC259aD8d817Bb92c2D863F74c";
  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);


  const cidList = ["baga6ea4seaqmw53htgtnx53c45sg33btibgoz7oseyqcm7kdc2tiosbdnq72sjy","123a6ea4seaqmw53htgtnx53c45sg33btibgoz7oseyqcm7kdc2tiosbdnq72abc"];
  const network = "filecoin_calibration";
  const dealId = "2763";


  const paid = "10000000000000000"; // paid filcoins
  const recipient = "0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a";
  //const terms = "2000000000000000";
  const status = true; // true for successful paid, false for failed paid

  const tx = await daoOracleInstance.connect(signer3).signCarTransaction(
    cidList,
    dealId,
    network,
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
