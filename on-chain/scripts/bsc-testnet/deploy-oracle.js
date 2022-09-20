const { ethers, upgrades } = require('hardhat')

const oneThousand = '1000000000000000000000'
const oneHundred = '1000000000000000000000'
const tenThousand = '1000000000000000000000'
const oneMillion = '1000000000000000000000000'

const deadline = '1000000000000000000'

const overrides = {
  gasLimit: 9999999,
}

const CHAINLINK_ADDRESS = '0x84b9b910527ad5c03a9ca831909e21e236ea7b06'
const ORACLE_ADDRESS = '0x4246103C6fF2e3eC839751156b518d066aab5e5A'
const JOB_ID = '7599d3c8f31e4ce78ad2b790cbcfc673'
const FEE = '50000000000000000'
const FEEx2 = '100000000000000000'

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

async function main() {
  const [deployer] = await ethers.getSigners()

  console.log('deployer: ', deployer.address)

  const FilinkConsumer = await ethers.getContractFactory('FilinkConsumer')

  const filink = await FilinkConsumer.deploy(
    CHAINLINK_ADDRESS,
    ORACLE_ADDRESS,
    FEE,
  )

  console.log('filink address:', filink.address)

  // transfer LINK
  const ERC20 = await ethers.getContractFactory('TestERC20')
  const linkToken = await ERC20.attach(CHAINLINK_ADDRESS)
  let transfer = await linkToken.transfer(filink.address, FEEx2)
  console.log('transfer hash:', transfer.hash)

  let tx = await filink.requestDealInfo('8588990', 'filecoin_mainnet')
  await tx.wait()
  await sleep(10000)
  console.log(await filink.price())
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
