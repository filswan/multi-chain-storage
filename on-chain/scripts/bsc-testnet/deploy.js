async function main() {
  const FILSWAN_ORACLE_THRESHOLD = 3
  const USDC_ADDRESS = '0x28fC65CF1F2bDe09ab2876fddaA7788340bAf1D7'
  const wFIL_ADDRESS = '0x67F4f8d7229e3296936AB63E6d0F3E8aB44B1118'
  const USDC_wFIL_PAIR = '0x30cc00a22c782e686c4646f0dd5d78220d53c2f3'

  const ONE = '1000000000000000000'

  const CHAINLINK_ADDRESS = '0x84b9b910527ad5c03a9ca831909e21e236ea7b06'
  const ORACLE_ADDRESS = '0x4246103C6fF2e3eC839751156b518d066aab5e5A'
  const FEE = '50000000000000000'

  const deployer = await ethers.getSigner()

  const FilinkConsumer = await ethers.getContractFactory('FilinkConsumer')
  const FilswanOracle = await ethers.getContractFactory('FilswanOracle')
  const MCSNFT = await ethers.getContractFactory('MCSNFT')
  const PriceFeed = await ethers.getContractFactory('PriceFeed')
  const SwanFaucet = await ethers.getContractFactory('SwanFaucet')
  const SwanPayment = await ethers.getContractFactory('SwanPayment')
  const SwanVRFConsumer = await ethers.getContractFactory('SwanVRFConsumer')

  console.log('deployer address:', deployer)

  const filinkConsumer = await FilinkConsumer.deploy(
    CHAINLINK_ADDRESS,
    ORACLE_ADDRESS,
    FEE,
  )

  console.log('filink address:', filink.address)

  const filswanOracle = await FilswanOracle.deploy(
    deployer,
    FILSWAN_ORACLE_THRESHOLD,
  )
  await filswanOracle.deployed()
  console.log('filswanOracle address:', filswanOracle.address)

  const nft = await upgrades.deployProxy(MCSNFT, [], overrides)
  await nft.deployed()
  console.log('MCSNFT address:', nft.address)

  const priceFeed = await PriceFeed.deploy(USDC_wFIL_PAIR, 1, overrides)
  await priceFeed.deployed()
  console.log('priceFeedInstance address:', priceFeed.address)
  const price = await priceFeed.consult(USDC_ADDRESS, ONE)
  console.log(`1 USDC to wFIL price: ${price}`)

  const swanFaucet = await upgrades.deployProxy(
    SwanFaucet,
    [deployer.address],
    overrides,
  )
  await swanFaucet.deployed()
  console.log(`swanFaucet address: ${swanFaucet.address}`)

  const swanPayment = await upgrades.deployProxy(
    SwanPayment,
    [
      deployer.address,
      USDC_ADDRESS,
      filswanOracle.address,
      priceFeed.address,
      filinkConsumer.address,
    ],
    overrides,
  )
  await swanPayment.deployed()
  console.log('swanPayment address:', swanPayment.address)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})
