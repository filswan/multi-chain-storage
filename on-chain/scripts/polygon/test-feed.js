const { ethers } = require('hardhat')

const ONE_ETH = ethers.utils.parseEther('1')
const USDC_WFIL_PAIR = '0x230e57B69E3d45557e20a6238462564EBf4Fe2a9'

async function main() {
  const [deployer] = await ethers.getSigners()
  console.log('deployer: ', deployer.address)

  const PriceFeed = await ethers.getContractFactory('PriceFeed')
  //const feed = await PriceFeed.deploy(USDC_WFIL_PAIR, 1, 6)
  const feed = PriceFeed.attach('0xa976fb4e934f07aa3b57e167879772dee2ca7a00')
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
