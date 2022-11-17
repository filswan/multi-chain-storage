const { expect } = require('chai')
const { upgrades } = require('hardhat')

describe('FilinkConsumer', function () {
  let FilinkConsumer
  let filink
  let caller

  describe('Deployment', function () {
    it('Attach FilinkConsumer', async function () {
      caller = await ethers.getSigner()

      FilinkConsumer = await ethers.getContractFactory('FilinkConsumer')
      filink = FilinkConsumer.attach(
        '0x2Bf5dBde4Fdd30de18b36405CF587044172ffD33',
      )

      //   console.log(await filink.functions)
    })
  })

  describe('Check Storage Price', function () {
    it('Should request deal info', async () => {
      //   try {
      //     let tx = await filink.requestDealInfo('123456', 'filecoin_mainnet', {
      //       gasLimit: 5000000,
      //     })
      //     await tx.wait()
      //   } catch (err) {}

      let price = await filink.getPrice('123456', 'filecoin_mainnet')
      console.log('storage price:', price.toString())
    })
  })
})
