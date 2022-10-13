const { ethers } = require('hardhat')

const ONE_ETH = ethers.utils.parseEther('1')
const USDC_WFIL_PAIR = '0xFF2278d5C5C5761d951Fd87bA077947Aa31d737f'

async function main() {
  const [deployer] = await ethers.getSigners()
  console.log('deployer: ', deployer.address)

  const PriceFeed = await ethers.getContractFactory('PriceFeed')
  //const feed = await PriceFeed.deploy(USDC_WFIL_PAIR, 1, 6)
  const feed = PriceFeed.attach('0x150FAc29ABc9E790e06C726920aad3cFE2862275')
  console.log('price feed address:', feed.address)

  const price = await feed.consult(USDC_WFIL_PAIR, ONE_ETH)

  console.log('1 wFIL to USDC:', ethers.utils.formatEther(price))
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
