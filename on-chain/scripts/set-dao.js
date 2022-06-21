
const hre = require("hardhat");

async function main() {

 
  const oracleDAOContractAddress = "0xe3262c0848b0cc5cd43df7139103f1fbf26558cc";

  const addressList = [
    "0x800210CfB747992790245eA878D32F188d01a03A",
    "0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d",
    "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1",
    "0x71632B0e6b5347BAc09E85a40B329397af473933",
    "0x21aE11DF412002378b73A28EF137FBfC59332BA4",
    "0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5",
  ];

  const [signer] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);


  const tx = await daoOracleInstance.connect(signer).setDAOUsers(addressList);
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
