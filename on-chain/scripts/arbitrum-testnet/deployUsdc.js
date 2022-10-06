const { ethers } = require('hardhat')

const overrides = {
  gasLimit: 9999999,
}

async function main() {
  const deployer = await ethers.getSigner()

  console.log('deployer: ', deployer.address)

  const USDC = await ethers.getContractFactory('USDCoin')
  const usdc = await USDC.deploy()
  await usdc.deployed()
  console.log(`USDCInstance address: ${usdc.address}`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
