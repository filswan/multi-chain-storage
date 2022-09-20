const { ethers, upgrades } = require('hardhat')

const uniswapv2FactoryABI = require('@uniswap/v2-core/build/IUniswapV2Factory.json')
  .abi
const sushiswapFactoryAddress = '0xc35DADB65012eC5796536bD9864eD8773aBc74C4'

const overrides = {
  gasLimit: 9999999,
}

const ONE = '1000000000000000000'

async function main() {
  const [deployer] = await ethers.getSigners()
  console.log('Deploying contracts with the account:', deployer.address)
  console.log('Account balance:', (await deployer.getBalance()).toString())

  const swanOracleContract = await ethers.getContractFactory('FilswanOracle')
  // threshold 2, meaning 2 votes are required to pass
  const swanOracleInstance = await upgrades.deployProxy(
    swanOracleContract,
    [deployer.address, 2],
    overrides,
  )
  await swanOracleInstance.deployed()
  console.log(`swanOracleInstance address: ${swanOracleInstance.address}`)

  /* MUMBAI ADDRESSES */
  // const swanOracleAddress = '0xe3262c0848b0cc5cd43df7139103f1fbf26558cc'
  // const usdcAddress = '0xe11A86849d99F524cAC3E7A0Ec1241828e332C62'
  // const wFilAddress = '0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055'
  // const pairAddress = '0x74038ed7D891A043d4aF41FeA242ED01914c2636'

  /* BSC TESTNET ADDRESSES */
  // const swanOracleAddress = ''
  // const usdcAddress = '0x28fC65CF1F2bDe09ab2876fddaA7788340bAf1D7'
  // const wFilAddress = '0x67F4f8d7229e3296936AB63E6d0F3E8aB44B1118'
  // const pairAddress = '0x30cc00a22c782e686c4646f0dd5d78220d53c2f3' // USDC/wFIL

  const priceFeedContract = await ethers.getContractFactory('PriceFeed')
  const priceFeedInstance = await priceFeedContract.deploy(
    pairAddress,
    1,
    overrides,
  )
  await priceFeedInstance.deployed()
  console.log(`priceFeedInstance address: ${priceFeedInstance.address}`)
  const price = await priceFeedInstance.consult(usdcAddress, ONE)
  console.log(`1 USDC to wFIL price: ${price}`)

  // const priceFeedOracleAddress = '0xe8a67994c114e0c17E1c135d0CB599a2394f1505'
  // const flinkOracleAddress = '0xcE9A9e594db39dCD449E392d68F60959533c0D75'

  const swanPaymentContract = await ethers.getContractFactory('SwanPayment')
  const swanPaymentInstance = await upgrades.deployProxy(
    swanPaymentContract,
    [
      deployer.address,
      usdcAddress,
      swanOracleAddress,
      priceFeedOracleAddress,
      flinkOracleAddress,
    ],
    overrides,
  )
  await swanPaymentInstance.deployed()
  console.log(`swanPaymentInstance address: ${swanPaymentInstance.address}`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
