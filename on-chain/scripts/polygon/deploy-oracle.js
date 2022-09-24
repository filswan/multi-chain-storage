const { ethers, upgrades } = require('hardhat')

const oneThousand = '1000000000000000000000'
const oneHundred = '1000000000000000000000'
const tenThousand = '1000000000000000000000'
const oneMillion = '1000000000000000000000000'

const deadline = '1000000000000000000'

const overrides = {
  gasLimit: 9999999,
}

const CHAINLINK_ADDRESS = '0xb0897686c545045afc77cf20ec7a532e3120e0f1'
const ORACLE_ADDRESS = '0x3d5552a177Fe380CDDe28a034EA93C2d30b80b2D'
const FEE = '100000000000000000'
const FEEx2 = '200000000000000000'

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

async function main() {
  const [deployer] = await ethers.getSigners()

  console.log('deployer: ', deployer.address)

  const FilinkConsumer = await ethers.getContractFactory('PolygonFilink')

  const filink = await FilinkConsumer.deploy(
    CHAINLINK_ADDRESS,
    ORACLE_ADDRESS,
    FEE,
  )

  console.log('filink address:', filink.address)

  // transfer LINK
  const ERC20 = await ethers.getContractFactory('TestERC20')
  const linkToken = ERC20.attach(CHAINLINK_ADDRESS)
  let transfer = await linkToken.transfer(filink.address, FEE)
  console.log('transfer hash:', transfer.hash)

  let tx = await filink.requestDealInfo('8588990', 'filecoin_mainnet')
  await sleep(20000)
  console.log(await filink.price())
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
