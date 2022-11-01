const { expect } = require('chai')
const { upgrades } = require('hardhat')

describe('Collection Factory', function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  // async function deployContracts() {
  // Contracts are deployed using the first signer/account by default

  //  return { renFil, nft, owner, otherAccount }
  // }

  let Factory
  let factory
  let u1
  let u2
  let u3

  describe('Deployment', function () {
    it('Should set the right owner', async function () {
      ;[u1, u2, u3] = await ethers.getSigners()

      Factory = await ethers.getContractFactory('CollectionFactory')
      factory = await upgrades.deployProxy(Factory, [])
      await factory.deployed()
      //const { nft, owner } = await loadFixture(deployContracts)

      expect(await factory.owner()).to.equal(u1.address)
    })
  })

  describe('Creating Collections', function () {
    it('u1 should have 2', async () => {
      let tx = await factory.createCollection('')
      let tx2 = await factory.createCollection('')
      expect((await factory.getCollections(u1.address)).length).to.equal(2)

      let collections = await factory.getCollections(u1.address)
      let tx3 = await factory.mint(collections[0], u1.address, 1, '')

      let tx4 = await tx3.wait()

      console.log(tx4.events)
    })
  })
})
