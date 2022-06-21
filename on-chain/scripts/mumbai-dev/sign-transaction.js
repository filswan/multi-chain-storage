
const hre = require("hardhat");

async function main() {

    // todo: change to singers
  const [signer, signer2, signer3] = await ethers.getSigners();

  const oracleDAOContractAddress = "0x6f83DA2C5f1C5AAC259aD8d817Bb92c2D863F74c";
  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);


  const cidList = ["0aac3d53-d293-4f90-9fb3-1f057df261a5QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C9z","519fac5e-069c-492f-bb7b-9e22ce0b0cd3QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C9z",
  "10b77004-e283-4779-ab0f-b866ed246822QmRXm5G4s4gFJMhXq2Wzu4N1QVLX5ygUbR48s25ffz5oVd"];
  const network = "filecoin_calibration";
  // const network = "filecoin_mainnet";

  const dealId = "172330";


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
