const { ethers } = require('hardhat')

const ONE_ETH = ethers.utils.parseEther('1')
const USDC_WFIL_PAIR = '0x30cc00a22c782e686c4646f0dd5d78220d53c2f3'

async function main() {
  const [deployer] = await ethers.getSigners()
  console.log('deployer: ', deployer.address)

  const PriceFeed = await ethers.getContractFactory('PriceFeed')
  //const feed = await PriceFeed.deploy(USDC_WFIL_PAIR, 1, 6)
  const feed = PriceFeed.attach('0x58d089CF567900cB63F82947Da4495A6c509Aa1b')
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
