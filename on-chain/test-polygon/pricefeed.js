const { expect } = require('chai')
const { upgrades } = require('hardhat')

describe('PriceFeed', function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  // async function deployContracts() {
  // Contracts are deployed using the first signer/account by default

  //  return { renFil, nft, owner, otherAccount }
  // }

  let PriceFeed
  let feed
  let caller

  describe('Deployment', function () {
    it('Attach PriceFeed', async function () {
      caller = await ethers.getSigner()

      PriceFeed = await ethers.getContractFactory('ChainlinkPriceFeed')
      feed = PriceFeed.attach('0xFC8B846fEd57579F91973F0561a08a235A39a8dA')

      expect(feed.address).to.equal(
        '0xFC8B846fEd57579F91973F0561a08a235A39a8dA',
      )

      console.log('version:', (await feed.version()).toString())
    })
  })

  describe('Check Feed Info', function () {
    it('Should be FIL/USD', async () => {
      let description = await feed.description()
      expect(description).to.equal('FIL / USD')
    })
    it('Should be 8 decimal', async () => {
      expect(await feed.decimals()).to.equal(8)
    })
  })

  describe('Check Price', function () {
    let latestPrice
    let consultPrice
    it('Should get latest price', async () => {
      latestPrice = await feed.getLatestPrice()
      console.log('latest price:', latestPrice.toString())
    })
    it('Should consult price', async () => {
      consultPrice = await feed.consult(
        '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
        '1000000000000000000',
      )
      console.log('latest price:', consultPrice.toString())
    })
  })
})
