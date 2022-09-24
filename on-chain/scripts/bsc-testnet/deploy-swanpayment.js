const { ethers, upgrades } = require('hardhat')

const overrides = {
  gasLimit: 9999999,
}

const USDC_ADDRESS = '0x28fC65CF1F2bDe09ab2876fddaA7788340bAf1D7'
const FILESWAN_ORACLE_ADDRESS = '0xA3015fE8483e6675310D09CB8F314059D004Ce65'
const PRICE_FEED_ADDRESS = '0x58d089CF567900cB63F82947Da4495A6c509Aa1b'
const FILINK_ADDRESS = '0x04a1b73FB8dD289E807FD0c584a060017ee0B26F'

async function main() {
  const [deployer] = await ethers.getSigners()

  console.log('deployer: ', deployer.address)

  const SwanPayment = await ethers.getContractFactory('SwanPayment')

  const swanPayment = await upgrades.deployProxy(
    SwanPayment,
    [
      deployer.address,
      USDC_ADDRESS,
      FILESWAN_ORACLE_ADDRESS,
      PRICE_FEED_ADDRESS,
      FILINK_ADDRESS,
    ],
    overrides,
  )
  await swanPayment.deployed()
  console.log('swanPayment address:', swanPayment.address)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
