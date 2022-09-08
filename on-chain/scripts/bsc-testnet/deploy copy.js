const { ethers, upgrades } = require('hardhat')

async function main() {
  const FILSWAN_ORACLE_THRESHOLD = 2
  const USDC_ADDRESS = '0x28fC65CF1F2bDe09ab2876fddaA7788340bAf1D7'
  const wFIL_ADDRESS = '0x67F4f8d7229e3296936AB63E6d0F3E8aB44B1118'
  const USDC_wFIL_PAIR = '0x30cc00a22c782e686c4646f0dd5d78220d53c2f3'

  const ONE = '1000000000000000000'

  const CHAINLINK_ADDRESS = '0x84b9b910527ad5c03a9ca831909e21e236ea7b06'
  const ORACLE_ADDRESS = '0x4246103C6fF2e3eC839751156b518d066aab5e5A'
  const FEE = '50000000000000000'
  const FEEx20 = '1000000000000000000'

  const DAO_USERS = [
    '0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5',
    '0x21aE11DF412002378b73A28EF137FBfC59332BA4',
    '0x71632B0e6b5347BAc09E85a40B329397af473933',
  ]

  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms))
  }

  const deployer = await ethers.getSigner()

  const overrides = {
    gasLimit: 9999999,
  }

  const FilinkConsumer = await ethers.getContractFactory('FilinkConsumer')
  const FilswanOracle = await ethers.getContractFactory('FilswanOracle')
  const MCSNFT = await ethers.getContractFactory('MCSNFT')
  const PriceFeed = await ethers.getContractFactory('PriceFeed')
  const SwanFaucet = await ethers.getContractFactory('SwanFaucet')
  const SwanPayment = await ethers.getContractFactory('SwanPayment')

  console.log('deployer address:', deployer.address)
  console.log('USDC address:', USDC_ADDRESS)
  console.log('wFIL address:', wFIL_ADDRESS)
  console.log('dex pair address:', USDC_wFIL_PAIR)

  const filinkConsumer = await FilinkConsumer.deploy(
    CHAINLINK_ADDRESS,
    ORACLE_ADDRESS,
    FEE,
  )
  console.log('filinkConsumer address:', filinkConsumer.address)

  /** TEST ORACLE
  // transfer LINK
  const ERC20 = await ethers.getContractFactory('TestERC20')
  const linkToken = await ERC20.attach(CHAINLINK_ADDRESS)
  let transfer = await linkToken.transfer(filinkConsumer.address, FEEx20)

  let tx = await filinkConsumer.requestDealInfo('8588990', 'filecoin_mainnet')
  await sleep(15000)
  const storagePrice = await filinkConsumer.getPrice(
    '8588990',
    'filecoin_mainnet',
  )
  console.log('storagePrice: ', storagePrice)
  */

  const filswanOracle = await upgrades.deployProxy(
    FilswanOracle,
    [deployer.address, FILSWAN_ORACLE_THRESHOLD],
    overrides,
  )
  await filswanOracle.deployed()
  console.log('filswanOracle address:', filswanOracle.address)

  const nft = await upgrades.deployProxy(MCSNFT, [], overrides)
  await nft.deployed()
  console.log('MCSNFT address:', nft.address)

  const priceFeed = await PriceFeed.deploy(USDC_wFIL_PAIR, 1)
  await priceFeed.deployed()
  console.log('priceFeed address:', priceFeed.address)

  /** TEST PRICE FEED 
  const price = await priceFeed.consult(USDC_ADDRESS, ONE)
  console.log(`1 USDC to wFIL price: ${price}`)
  */

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

  console.log('setting DAO_USERS...')
  let tx = await filswanOracle.setDAOUsers(DAO_USERS)
  console.log(tx.hash)
  console.log('setting filinkOracle...')
  tx = await filswanOracle.setFilinkOracle(filinkConsumer.address)
  console.log(tx.hash)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})
