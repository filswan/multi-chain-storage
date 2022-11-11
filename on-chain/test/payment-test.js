const { expect } = require('chai')

describe('Swan Payment', function () {
  let accounts
  let paymentInstance
  let token

  before('Deploy SwanPayment', async () => {
    accounts = await ethers.getSigners()

    const TokenContract = await ethers.getContractFactory('USDCoin')
    token = await TokenContract.deploy()

    const contract = await ethers.getContractFactory('SwanPayment')
    paymentInstance = await upgrades.deployProxy(contract, [
      accounts[0].address,
      token.address,
      token.address,
      token.address,
      token.address,
    ])
    await paymentInstance.deployed() // have to wait block to include this transaction

    // console.log('deploy oracle instance')
    // const oracleContract = await ethers.getContractFactory('FilecoinOracle')
    // oracleInstance = await oracleContract.deploy(accounts[0].address)
    // await oracleInstance.deployed()

    // console.log('payment instance deployed at:', oracleInstance.address)

    // await oracleInstance.setDAOUsers(daoAddressList)

    // console.log('set oracle')
    // await paymentInstance.setOracle(oracleInstance.address)

    // const tx = {
    //   to: wethAddress,
    //   value: ethers.utils.parseEther('0.5'),
    // }

    // await accounts[0].sendTransaction(tx)

    await token.mint(accounts[0].address, '100000000') // 100 USDC
  })

  describe('Payment oracle workflow', async () => {
    // const accounts = await ethers.getSigners();
    // it('Test only some dao users update payment info', async function () {
    //   const cid_zero =
    //     'bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw00'
    //   for (let i = 0; i < oracleAccounts.length - 1; i++) {
    //     const tx = await oracleInstance
    //       .connect(oracleAccounts[i])
    //       .updatePaymentInfo(cid_zero, actualPay20Native)
    //     await tx.wait()
    //   }

    //   const result = await oracleInstance.getPaymentInfo(cid_zero)
    //   console.log('current payment is ', result.toString())
    //   expect(zero).to.equal(result)
    // })

    // it('Test all dao users update payment info', async function () {
    //   const cid_20 =
    //     'bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw20'
    //   for (let i = 0; i < oracleAccounts.length; i++) {
    //     const tx = await oracleInstance
    //       .connect(oracleAccounts[i])
    //       .updatePaymentInfo(cid_20, actualPay20Native)
    //     await tx.wait()
    //   }

    //   const result = await oracleInstance.getPaymentInfo(cid_20)
    //   console.log('current payment is ', result.toString())
    //   expect(actualPay20Native).to.equal(result)
    // })

    it('Test Lock payment', async function () {
      const cid =
        'bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72'

      const payer = accounts[3]
      const fileswanRecipient = accounts[5]
      const minAmount = '1000000'
      const amount = '1500000'
      const lockTime = 60 * 60 * 24 * 6

      await token.mint(payer.address, '100000000')
      await token.connect(payer).approve(paymentInstance.address, amount)

      const params = {
        id: cid,
        minPayment: minAmount,
        amount: amount,
        lockTime: 86400 * lockTime,
        recipient: fileswanRecipient.address,
        size: '20000',
        copyLimit: 5,
      }

      const tx = await paymentInstance.connect(payer).lockTokenPayment(params)
      await tx.wait()

      const result = await paymentInstance.getLockedPaymentInfo(cid)
      //console.log(result)
      expect(payer.address).to.equal(result.owner)
      expect(fileswanRecipient.address).to.equal(result.recipient)
      expect(amount).to.equal(result.lockedFee)
    })

    //   it('Test Unlock payment great than minPayment after oracle updated info', async function () {
    //     const cid =
    //       'bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72'

    //     const payer = accounts[3]
    //     const fileswanRecipient = accounts[5]

    //     const beforePaymentBalance = await fileswanRecipient.getBalance()
    //     expect(beforePaymentBalance).to.equal(0)

    //     for (let i = 0; i < oracleAccounts.length; i++) {
    //       const tx = await oracleInstance
    //         .connect(oracleAccounts[i])
    //         .updatePaymentInfo(cid, actualPay20Native)
    //       await tx.wait()
    //     }

    //     // payer and recipient can unlock the payment
    //     const tx = await paymentInstance.connect(payer).unlockPayment(cid)

    //     await tx.wait()

    //     const result = await paymentInstance.getLockedPaymentInfo(cid)
    //     expect(result._isExisted).to.equal(false)
    //     expect(result.lockedFee).to.equal(0)

    //     const afterPaymentBalance = await fileswanRecipient.getBalance()
    //     expect(afterPaymentBalance).to.equal(actualPay20Native)
    //   })
  })
})
