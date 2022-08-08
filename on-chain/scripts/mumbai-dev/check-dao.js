
const hre = require("hardhat");

const overrides = {
  gasLimit: 9999999
}

async function main() {

 
  const oracleDAOContractAddress = "0x6f83DA2C5f1C5AAC259aD8d817Bb92c2D863F74c";

  const addressList = [
    "0x800210CfB747992790245eA878D32F188d01a03A",
    "0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d",
    "0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1"
  ];

  const recipient = "0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a";

  const [signer, signer2, signer3] = await ethers.getSigners();

  const contract = await hre.ethers.getContractFactory("FilswanOracle");
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress);

  // const result = await daoOracleInstance.getCarPaymentVotes('6976653', 'filecoin_mainnet', '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a');
  // console.log(result);

  // const tx1 = await daoOracleInstance.connect(signer3).preSign('20015', 'filecoin_calibration', recipient, 1);
  // await tx1.wait();

  // const cidList = ['e38bb9da-67d2-4374-ae46-9d67ce93b39eQmf4N6uo9NqYPgRTvCnkhMZh1wjGAxdFtgh74RGFyh1mW1',
  // 'a6aa114b-df42-442c-92dd-537db6905115Qmf4N6uo9NqYPgRTvCnkhMZh1wjGAxdFtgh74RGFyh1mW1'];

  // const tx2 = await daoOracleInstance.connect(signer3).sign('20015', 'filecoin_calibration', cidList, 0);
  // await tx2.wait();

  const result1 = await daoOracleInstance.getCidList('8460147', 'filecoin_mainnet');
  console.log(result1);

  

  // const result2 = await daoOracleInstance.getSignatureList('6976653', 'filecoin_mainnet');
  // console.log(result2);

  // const result3 = await daoOracleInstance.getHashKey('20011', 'filecoin_calibration', "0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a", ["d337b095-8337-41a3-b66f-5555e2f48b0dQmdytmR4wULMd3SLo6ePF4s3WcRHWcpnJZ7bHhoj3QB13v","03e4f239-3f4f-4190-866d-674c40640afcQmarHqYRBJA2CAZD78vPcMLWyRrLLVJH32ewzqtF11VTVT","edcdb983-db0d-4b79-be58-d5a4d2819170QmQi8DVQjGbqbEYsbHaGkJNLFwfHkWFYsJNDjCnjsaJzmd","3694d473-49a2-46e1-8b7f-7b978c8a3e93QmcrqzqovbqgrhoPZPes6p4FejSRmJDdNfQ1Zic7eZaErp","92e03f36-83a9-4e51-a005-1f9ba92a0e3fQmf4N6uo9NqYPgRTvCnkhMZh1wjGAxdFtgh74RGFyh1mW1"]);
  // console.log(result3);
  
  
  const result2 = await daoOracleInstance.isCarPaymentAvailable('8460147', 'filecoin_mainnet', '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a');
  console.log(result2);

  // const hasRole = await daoOracleInstance.connect(signer).hasRole(
  //   role,
  //   addressList[0]);

  // console.log(hasRole);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
