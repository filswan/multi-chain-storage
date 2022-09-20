const { ethers, upgrades } = require('hardhat')

async function main() {
  const FILSWAN_ORACLE_THRESHOLD = 2
  const USDC_ADDRESS = '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174'
  const wFIL_ADDRESS = '0xEde1B77C0Ccc45BFa949636757cd2cA7eF30137F'
  const USDC_wFIL_PAIR = '0x230e57B69E3d45557e20a6238462564EBf4Fe2a9'
  const SUSHISWAP_FACTORY_ADDRESS = '0xc35DADB65012eC5796536bD9864eD8773aBc74C4'

  const ONE = '1000000000000000000'

  const CHAINLINK_ADDRESS = '0xb0897686c545045afc77cf20ec7a532e3120e0f1'
  const ORACLE_ADDRESS = '0x3d5552a177Fe380CDDe28a034EA93C2d30b80b2D'
  const FEE = '100000000000000000'
  const FEEx20 = '2000000000000000000'

  const DAO_USERS = [
    '0xd44CDe0f3BeEF47Af166fC763b52977A43bf8Fe7',
    '0xa1c24757Da070E62b12bf2c762213c35FA30Fae5',
    '0x74D5Ef1C805EA2D42ee3163fA312c9e156635C68',
  ]

  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms))
  }

  const deployer = await ethers.getSigner()

  const overrides = {
    gasLimit: 9999999,
  }

  const FilinkConsumer = await ethers.getContractFactory('PolygonFilink')
  const FilswanOracle = await ethers.getContractFactory('FilswanOracle')
  const MCSNFT = await ethers.getContractFactory('MCSNFT')
  // const PriceFeed = await ethers.getContractFactory('PriceFeed')
  // const SwanFaucet = await ethers.getContractFactory('SwanFaucet')
  // const SwanPayment = await ethers.getContractFactory('SwanPayment')
  // const PriceOracleFeed = await ethers.getContractFactory('PriceOracleFeed')

  console.log('deployer address:', deployer.address)
  console.log('USDC address:', USDC_ADDRESS)
  console.log('wFIL address:', wFIL_ADDRESS)
  console.log('dex pair address:', USDC_wFIL_PAIR)
  console.log('sushiswap factory address:', SUSHISWAP_FACTORY_ADDRESS)

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

  //   const priceFeed = await PriceFeed.deploy(USDC_wFIL_PAIR, 1)
  //   await priceFeed.deployed()
  //   console.log('priceFeed address:', priceFeed.address)

  //   const priceOracleFeed = await PriceOracleFeed.deploy(
  //     FACTORY,
  //     USDC_ADDRESS,
  //     wFIL_ADDRESS,
  //   )
  //   await priceOracleFeed.deployed()
  //   console.log('priceOracleFeed address:', priceOracleFeed.address)

  /** TEST PRICE FEED 
  const price = await priceFeed.consult(USDC_ADDRESS, ONE)
  console.log(`1 USDC to wFIL price: ${price}`)
  */

  //   const swanFaucet = await upgrades.deployProxy(
  //     SwanFaucet,
  //     [deployer.address],
  //     overrides,
  //   )
  //   await swanFaucet.deployed()
  //   console.log(`swanFaucet address: ${swanFaucet.address}`)

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
