
const hre = require("hardhat");

async function main() {

 // todo: change address
  const oracleDAOContractAddress = "0x6f83DA2C5f1C5AAC259aD8d817Bb92c2D863F74c";

  // todo: update addresses
  const addressList = [
    "0x6d2e5279b106843f6E924194401B50e6e27FE12a",
    "0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5",
    "0xeA2bf08288bbfB0d3DBf534f35af32bF2c6E5e45",
    "0x800210cfb747992790245ea878d32f188d01a03a",
    "0x21aE11DF412002378b73A28EF137FBfC59332BA4",
    "0x71632B0e6b5347BAc09E85a40B329397af473933"
  ];

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  let tx = await daoOracleInstance.connect(signer).setFilinkOracle(filinkAddress);
  await tx.wait();

  tx = await daoOracleInstance.connect(signer).setDAOUsers(addressList);
  await tx.wait();

  const role = await daoOracleInstance.DAO_ROLE.call();
  console.log(role);

  const hasRole = await daoOracleInstance.connect(signer).hasRole(
    role,
    addressList[0]);

  console.log(hasRole);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
