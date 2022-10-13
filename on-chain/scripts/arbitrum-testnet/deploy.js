async function main() {
  const USDC_ADDRESS = '0x1FdDc28136a57CF1713E5Fc416953687Fe2Ba339'
  const RENFIL_ADDRESS = '0x2da6871c3Fd3598Bc2249901df01139bDfFe815e'
  const PAIR_ADDRESS = '0xFF2278d5C5C5761d951Fd87bA077947Aa31d737f'

  const ONE = ethers.utils.parseEther('1')

  const deployer = await ethers.getSigner()

  const MCSNFT = await ethers.getContractFactory('MCSNFT')
  const PriceFeed = await ethers.getContractFactory('PriceFeed')
  const SwanPayment = await ethers.getContractFactory('SwanPaymentNoUnlock')
  const SwanFaucet = await ethers.getContractFactory('SwanFaucet')

  console.log('deployer address:', deployer.address)
  console.log('usdc address:', USDC_ADDRESS)
  console.log('renFil address:', RENFIL_ADDRESS)
  console.log('usdc/renFil pair address:', PAIR_ADDRESS)

  const nft = await upgrades.deployProxy(MCSNFT, [])
  await nft.deployed()
  console.log('MCSNFT address:', nft.address)

  const priceFeed = await PriceFeed.deploy(PAIR_ADDRESS, 1, 6)
  await priceFeed.deployed()
  console.log('priceFeedInstance address:', priceFeed.address)
  //   const price = await priceFeed.consult(USDC_ADDRESS, ONE)
  //   console.log(`1 wFIL to USDC price: ${price}`)

  const swanFaucet = await upgrades.deployProxy(SwanFaucet, [deployer.address])
  await swanFaucet.deployed()
  console.log(`swanFaucet address: ${swanFaucet.address}`)

  const swanPayment = await upgrades.deployProxy(SwanPayment, [
    deployer.address,
    USDC_ADDRESS,
    priceFeed.address,
  ])
  await swanPayment.deployed()
  console.log('swanPayment address:', swanPayment.address)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})
