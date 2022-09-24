const hre = require('hardhat')

const overrides = {
  gasLimit: 9999999,
}

async function main() {
  const oracleDAOContractAddress = '0xA3015fE8483e6675310D09CB8F314059D004Ce65'

  const addressList = [
    '0x800210CfB747992790245eA878D32F188d01a03A',
    '0x05856015d07F3E24936B7D20cB3CcfCa3D34B41d',
    '0x6f2B76024196e82D81c8bC5eDe7cff0B0276c9C1',
  ]

  const recipient = '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a'

  const [signer3] = await ethers.getSigners()

  const contract = await hre.ethers.getContractFactory('FilswanOracle')
  const daoOracleInstance = await contract.attach(oracleDAOContractAddress)

  // const result = await daoOracleInstance.getCarPaymentVotes('6976653', 'filecoin_mainnet', '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a');
  // console.log(result);

  // const tx1 = await daoOracleInstance.connect(signer3).preSign('20015', 'filecoin_calibration', recipient, 1);
  // await tx1.wait();

  //   const cidList = [
  //     'e38bb9da-67d2-4374-ae46-9d67ce93b39eQmf4N6uo9NqYPgRTvCnkhMZh1wjGAxdFtgh74RGFyh1mW1',
  //     'a6aa114b-df42-442c-92dd-537db6905115Qmf4N6uo9NqYPgRTvCnkhMZh1wjGAxdFtgh74RGFyh1mW1',
  //   ]

  //   const tx2 = await daoOracleInstance
  //     .connect(signer3)
  //     .sign('20015', 'filecoin_calibration', cidList, 0)
  //   await tx2.wait()

  const DEAL_ID = '206678'
  const NETWORK = 'filecoin_calibration'

  const result1 = await daoOracleInstance.getCidList(
    '206678',
    'filecoin_calibration',
  )
  console.log(result1)

  // const result2 = await daoOracleInstance.getSignatureList('6976653', 'filecoin_mainnet');
  // console.log(result2);

  const result3 = await daoOracleInstance.getHashKey(
    '206678',
    'filecoin_calibration',
    '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a',
    [
      '3f11483f-a393-4c10-a94d-bc5de04301d9QmXqqnbRtiYyErKiY6Y3REZ7kGppNFjyYTQAEBGoDo9ZVb',
    ],
  )
  console.log(result3)

  const result2 = await daoOracleInstance.isCarPaymentAvailable(
    '206678',
    'filecoin_calibration',
    '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a',
  )
  console.log(result2)

  // const hasRole = await daoOracleInstance.connect(signer).hasRole(
  //   role,
  //   addressList[0]);

  // console.log(hasRole);

  const result4 = await daoOracleInstance.getCarPaymentVotes(
    '206678',
    'filecoin_calibration',
    '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a',
  )
  console.log(result4)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
