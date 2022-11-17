const { expect } = require('chai')
const { upgrades, ethers } = require('hardhat')

describe('FilswanOracle', function () {
  let SwanPayment
  let payment
  let caller
  let id

  describe('Deployment', function () {
    it('Attach SwanPayment', async function () {
      caller = await ethers.getSigner()

      SwanPayment = await ethers.getContractFactory('SwanPayment')
      payment = SwanPayment.attach('0xA1f32c758c4324cC3070A3AA107C4dC7DdFe1a6f')

      //console.log(payment.functions)
    })
  })

  describe('Payment', () => {
    it('Should lock tokens', async () => {
      //   const USDC = (await ethers.getContractFactory('USDCoin')).attach(
      //     '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
      //   )
      //   const approve = await USDC.approve(payment.address, '5')
      //   console.log(approve.hash)

      //   id = `test-script-${Date.now()}`
      //   console.log('id:', id)

      //   const lockObj = {
      //     id: id,
      //     minPayment: '2',
      //     amount: '3',
      //     lockTime: 60 * 5,
      //     recipient: caller.address,
      //     size: 100,
      //     copyLimit: 5,
      //   }

      //   const tx = await payment.lockTokenPayment(lockObj, { gasLimit: 8000000 })
      //   console.log(tx.hash)

      const info = await payment.getLockedPaymentInfo(
        'test-script-1668704439192',
      )
      //console.log(info)

      expect(info._isExisted).to.equal(true)
    })
  })
})
