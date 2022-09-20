const { ethers } = require('hardhat')

const ONE_ETH = ethers.utils.parseEther('1')
const USDC_WFIL_PAIR = '0x230e57B69E3d45557e20a6238462564EBf4Fe2a9'

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

async function main() {
  const [deployer] = await ethers.getSigners()
  console.log('deployer: ', deployer.address)

  const FilinkConsumer = await ethers.getContractFactory('PolygonFilink')

  const filink = FilinkConsumer.attach(
    '0x2Bf5dBde4Fdd30de18b36405CF587044172ffD33',
  )

  const id = await filink.requestDealInfo('8588990', 'filecoin_mainnet')
  await id.wait()

  await sleep(10000)
  const price = await filink.getPrice('8588990', 'filecoin_mainnet')
  console.log(price)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
