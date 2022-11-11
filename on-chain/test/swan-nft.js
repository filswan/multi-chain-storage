const { expect } = require('chai')
const { upgrades } = require('hardhat')

describe('SwanNFT', function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  // async function deployContracts() {
  // Contracts are deployed using the first signer/account by default

  //  return { renFil, nft, owner, otherAccount }
  // }

  let SwanNFT
  let nft
  let owner

  describe('Deployment', function () {
    it('Should set the right owner', async function () {
      owner = await ethers.getSigner()

      SwanNFT = await ethers.getContractFactory('SwanNFT')
      nft = await upgrades.deployProxy(SwanNFT, [])
      await nft.deployed()
      //const { nft, owner } = await loadFixture(deployContracts)

      expect(await nft.owner()).to.equal(owner.address)
    })
  })

  describe('Minting NFT (with URI)', function () {
    it('Should have supply of 1', async () => {
      const tx = await nft.mintUnique(owner.address, 'test')
      expect(await nft.totalSupply(1)).to.equal(1)
    })
    it('Should be unique', async () => {
      expect(await nft.isUnique(1)).to.equal(true)
    })
    it('Should have URI set', async () => {
      const uri = await nft.uri(1)
      expect(uri).to.equal('test')
    })
    it('Should not be able to mint more', async () => {
      await expect(nft.mintMore(owner.address, 1, 2, '')).to.be.reverted
    })
    it('Should not be able to change URI', async () => {
      await expect(nft.setURI(1, 'test2')).to.be.reverted
    })
  })

  describe('Minting NFT (without URI)', function () {
    it('Should have supply of 1', async () => {
      const tx = await nft.mintUnique(owner.address, '')
      expect(await nft.totalSupply(2)).to.equal(1)
    })
    it('Should be unique', async () => {
      expect(await nft.isUnique(2)).to.equal(true)
    })
    it('Should not have URI set', async () => {
      const uri = await nft.uri(2)
      expect(uri).to.equal('')
    })
    it('Should not be able to mint more', async () => {
      await expect(nft.mintMore(owner.address, 1, 2, '')).to.be.reverted
    })
    it('Should be able to change URI', async () => {
      await nft.setURI(2, 'test2')
      expect(await nft.uri(2)).to.equal('test2')
    })
  })

  describe('Minting SFT (without URI)', function () {
    it('Should have supply of 100', async () => {
      const tx = await nft.mint(owner.address, 100, 0)
      expect(await nft.totalSupply(3)).to.equal(100)
    })
    it('Should not be unique', async () => {
      expect(await nft.isUnique(3)).to.equal(false)
    })
    it('Should not have URI set', async () => {
      const uri = await nft.uri(3)
      expect(uri).to.equal('')
    })
    it('Should be able to mint more', async () => {
      await nft.mintMore(owner.address, 3, 100, 0)
      expect(await nft.totalSupply(3)).to.equal(200)
    })
    it('Should be able to change URI', async () => {
      await nft.setURI(3, 'test3')
      expect(await nft.uri(3)).to.equal('test3')
    })
  })
})
