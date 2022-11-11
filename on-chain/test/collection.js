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

  let factory
  let defaultCollection
  let Collection
  let accounts

  describe('Deployment', function () {
    it('Should set the right owner', async function () {
      accounts = await ethers.getSigners()

      Collection = await ethers.getContractFactory('MCSCollection')
      defaultCollection = await Collection.deploy('')

      const Factory = await ethers.getContractFactory('CollectionFactory')
      factory = await upgrades.deployProxy(Factory, [defaultCollection.address])
      await factory.deployed()
      //const { nft, owner } = await loadFixture(deployContracts)

      expect(await factory.owner()).to.equal(accounts[0].address)
    })
  })

  describe('Creating Collections', function () {
    it('user 1 should have 2', async () => {
      let tx = await factory.createCollection('')
      let tx2 = await factory.createCollection('')
      expect(
        (await factory.getCollections(accounts[0].address)).length,
      ).to.equal(2)
    })
  })

  describe('Minting', function () {
    it('should mint to existing collection', async () => {
      let collections = await factory.getCollections(accounts[0].address)
      let tx = await factory.mint(collections[0], accounts[0].address, 1, '')

      //   let tx4 = await tx3.wait()
      //   console.log(tx4.events)
      let count = await Collection.attach(collections[0]).idCount()
      expect(count).to.equal(1)
    })

    it('should mint semi-fungible token', async () => {
      let collections = await factory.getCollections(accounts[0].address)
      let tx = await factory.mint(collections[1], accounts[0].address, 5, '')
      let supply = await Collection.attach(collections[1]).totalSupply(1)

      expect(supply).to.equal(5)
    })

    it('should mint to new collection', async () => {
      await factory.mintToNewCollection('', accounts[0].address, 1, '')

      let collections = await factory.getCollections(accounts[0].address)
      expect(collections.length).to.equal(3)

      let count = await Collection.attach(collections[2]).idCount()
      expect(count).to.equal(1)
    })

    it("should not be able to mint to another user's collection", async () => {
      let collections = await factory.getCollections(accounts[0].address)
      await expect(
        factory
          .connect(accounts[1])
          .mint(collections[0], accounts[0].address, 1, ''),
      ).to.be.revertedWith('caller is not admin')

      let count = await Collection.attach(collections[0]).idCount()
      expect(count).to.equal(1)
    })
  })
})
