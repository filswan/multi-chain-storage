const { ethers, upgrades } = require('hardhat')

const CONSULT_AMOUNT = '1'

async function main() {
  const PriceFeed = await ethers.getContractFactory('ChainlinkPriceFeed')
  const priceFeed = PriceFeed.attach(
    '0xFC8B846fEd57579F91973F0561a08a235A39a8dA',
  )

  console.log('priceFeed address:', priceFeed.address)

  let price = await priceFeed.getLatestPrice()
  let decimals = await priceFeed.decimals()
  let name = await priceFeed.description()

  console.log('feed:', name)
  console.log('price:', price.toString())
  console.log('decimals:', decimals)

  console.log('\nconsulting...')
  const USDC_ADDRESS = '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174'
  const consultPrice = await priceFeed.consult(
    USDC_ADDRESS,
    ethers.utils.parseEther(CONSULT_AMOUNT),
  )
  console.log(`${CONSULT_AMOUNT} FIL = ${consultPrice.toString()} (6 decimal)`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
