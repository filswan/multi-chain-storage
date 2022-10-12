const { ethers } = require('hardhat')

const overrides = {
  gasLimit: 9999999,
}

async function main() {
  const deployer = await ethers.getSigner()
  console.log('deployer: ', deployer.address)

  const Router = await ethers.getContractFactory('UniswapV2Router01')
  const router = Router.attach('0xd5328354c83BA6bCe97CD24b444b070Dd00AB1cf')

  const USDC = await ethers.getContractFactory('USDCoin')
  const usdc = USDC.attach('0x1FdDc28136a57CF1713E5Fc416953687Fe2Ba339')

  const ERC20 = await ethers.getContractFactory('TestERC20')
  const renFil = ERC20.attach('0x2da6871c3Fd3598Bc2249901df01139bDfFe815e')

  const UNIX_NOW = Math.floor(Date.now() / 1000)
  const tx = await router.addLiquidity(
    usdc.address,
    renFil.address,
    '500000000',
    ethers.utils.parseEther('100'),
    '500000000',
    ethers.utils.parseEther('100'),
    deployer.address,
    UNIX_NOW + 60 * 60 * 24 * 3, // deadline = 3 days from now
  )
  await tx.wait()
  console.log(tx.hash)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
