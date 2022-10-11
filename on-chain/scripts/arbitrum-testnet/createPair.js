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

  const ERC20 = await ethers.getContractFactory('TestERC20')
  const renFil = await ERC20.deploy('RenFil', 'renFIL')
  await renFil.deployed()
  console.log(`RenFil address: ${renFil.address}`)

  const UniswapFactory = await ethers.getContractFactory('UniswapV2Library')
  const uniswap = await UniswapFactory.deploy(deployer.address)
  await uniswap.deployed()
  console.log(`UniswapFactory address: ${uniswap.address}`)

  const tx = await uniswap.createPair(usdc.address, renFil.address)
  console.log(tx.hash)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
