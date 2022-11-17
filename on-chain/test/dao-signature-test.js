const { expect } = require('chai')

// Important, this test case only works on local test
// please comments FilswanOracle.sol filink call before run this test case
describe('DAO Signatures', function () {
  const overrides = {
    gasLimit: 9999999,
  }

  let accounts
  let oracleAccounts
  let oracleInstance
  let filink

  const fee = {
    // To convert Ether to Wei:
    value: ethers.utils.parseEther('0.72'), // ether in this case MUST be a string

    // Or you can use Wei directly if you have that:
    // value: someBigNumber
    // value: 1234   // Note that using JavaScript numbers requires they are less than Number.MAX_SAFE_INTEGER
    // value: "1234567890"
    // value: "0x1234"

    // Or, promises are also supported:
    // value: provider.getBalance(addr)
  }

  const threshold = 2

  before('Deploy DAO contract', async () => {
    accounts = await ethers.getSigners()
    oracleAccounts = accounts.slice(0, -1)

    const daoAddressList = oracleAccounts.map((a) => a.address)

    // console.log('deploy payment instance')
    // const contract = await ethers.getContractFactory("SwanPayment");
    // paymentInstance = await contract.deploy(accounts[0].address);
    // await paymentInstance.deployed(); // have to wait block to include this transaction

    // console.log("payment instance deployed at:", paymentInstance.address);
    const filinkContract = await ethers.getContractFactory('MockFilink')
    filink = await filinkContract.deploy()

    const oracleContract = await ethers.getContractFactory('FilswanOracleV2')
    oracleInstance = await upgrades.deployProxy(
      oracleContract,
      [accounts[0].address, threshold],
      overrides,
    )
    await oracleInstance.deployed()
    await oracleInstance.setDAOUsers(daoAddressList)
    await oracleInstance.setFilinkOracle(filink.address)

    // console.log('set oracle')
    // await paymentInstance.setOracle(oracleInstance.address);

    // const tx = {
    //   to: wethAddress,
    //   value: ethers.utils.parseEther("0.5")
    // };

    // await accounts[0].sendTransaction(tx);
  })

  describe('Oracle Signature', async () => {
    // const accounts = await ethers.getSigners();

    // const zero = ethers.BigNumber.from("0");
    // const actualPay20Native = ethers.BigNumber.from("200000000000000000");

    const cidList1 = [
      '0aac3d53-d293-4f90-9fb3-1f057df261a5QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C9z',
      '519fac5e-069c-492f-bb7b-9e22ce0b0cd3QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C9z',
      '10b77004-e283-4779-ab0f-b866ed246822QmRXm5G4s4gFJMhXq2Wzu4N1QVLX5ygUbR48s25ffz5oVd',
      '10b77004-e283-4779-ab0f-b866ed246822QmRXm5G4s4gFJMhXq2Wzu4N1QVLX5ygUbR48s25ffz5oVd',
    ]

    const cidList2 = [
      '0aac3d53-d293-4f90-9fb3-1f057df261a5QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C90',
      '519fac5e-069c-492f-bb7b-9e22ce0b0cd3QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C91',
      '10b77004-e283-4779-ab0f-b866ed246822QmRXm5G4s4gFJMhXq2Wzu4N1QVLX5ygUbR48s25ffz5oV2',
    ]

    const user2CidList1 = [
      '0aac3d53-d293-4f90-9fb3-1f057df261a5QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C9z',
      '519fac5e-069c-492f-bb7b-9e22ce0b0cd3QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C9z',
      '10b77004-e283-4779-ab0f-b866ed246822QmRXm5G4s4gFJMhXq2Wzu4N1QVLX5ygUbR48s25ffz5oVd',
      '0aac3d53-d293-4f90-9fb3-1f057df261a5QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C90',
    ]

    const user2CidList2 = [
      '519fac5e-069c-492f-bb7b-9e22ce0b0cd3QmUKn5wLZWPzbsy4ihnoDMaTSSNhrZY7c34kSWeqLv9C91',
      '10b77004-e283-4779-ab0f-b866ed246822QmRXm5G4s4gFJMhXq2Wzu4N1QVLX5ygUbR48s25ffz5oV2',
    ]

    const batch = 2

    const network = 'filecoin_mainnet'

    const recipient = '0xc4fcaAdCb0b00a9501e56215c37B10fAF9e79c0a'

    async function doPreSign(deal) {
      for (let i = 0; i < oracleAccounts.length - 1; i++) {
        const tx = await oracleInstance
          .connect(oracleAccounts[i])
          .preSign(deal, network, recipient, batch)
        await tx.wait()

        const result = await oracleInstance.getOracleInfo(
          deal,
          network,
          oracleAccounts[i].address,
        )
        expect(true).to.equal(result.flag)
        expect(3).to.equal(result.signStatus)
      }
    }

    it('Test correct signature', async function () {
      const deal = '6976653'

      await doPreSign(deal)

      let tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, cidList1, 0)
      await tx.wait()

      tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, cidList2, 1)
      await tx.wait()

      let cidList = await oracleInstance.getCidList(deal, network)
      expect(cidList.length).to.equal(0)

      tx = await oracleInstance
        .connect(oracleAccounts[1])
        .sign(deal, network, user2CidList1, 0)
      await tx.wait()

      // sign cid list 1 one time
      tx = await oracleInstance
        .connect(oracleAccounts[1])
        .sign(deal, network, user2CidList2, 1)
      await tx.wait()

      cidList = await oracleInstance.getCidList(deal, network)

      expect(cidList.length).to.equal(6)
    })

    it('Test different cid list signature', async function () {
      const deal = '6976655'
      await doPreSign(deal)

      let tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, cidList1, 0)
      await tx.wait()

      tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, cidList1, 1)
      await tx.wait()

      let cidList = await oracleInstance.getCidList(deal, network)
      expect(cidList.length).to.equal(0)

      tx = await oracleInstance
        .connect(oracleAccounts[1])
        .sign(deal, network, user2CidList1, 0)
      await tx.wait()

      tx = await oracleInstance
        .connect(oracleAccounts[1])
        .sign(deal, network, user2CidList2, 1)
      await tx.wait()

      cidList = await oracleInstance.getCidList(deal, network)

      expect(cidList.length).to.equal(0)
    })

    it('Test wrong batch no', async function () {
      const deal = '6976656'
      await doPreSign(deal)

      let tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, cidList1, 0)
      await tx.wait()

      // todo: err is not correct
      await expect(
        oracleInstance
          .connect(oracleAccounts[0])
          .sign(deal, network, cidList2, 2),
      ).to.be.revertedWith('wrong batch No')
    })

    it('Test same batch signature', async function () {
      const deal = '6976657'
      await doPreSign(deal)

      let tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, cidList1, 0)
      await tx.wait()

      await expect(
        oracleInstance
          .connect(oracleAccounts[0])
          .sign(deal, network, cidList1, 0),
      ).to.be.revertedWith('already signed the batch')

      let cidList = await oracleInstance.getCidList(deal, network)
      expect(cidList.length).to.equal(0)
    })

    it('Should not be able to sign twice with the same address', async function () {
      const deal = '6976658'
      await doPreSign(deal)

      let hashKey = await oracleInstance.getHashKey(deal, network, recipient, [
        ...user2CidList1,
        ...user2CidList2,
      ])
      //console.log('key:', hashKey)

      let tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, user2CidList1, 0)
      tx = await tx.wait()

      tx = await oracleInstance
        .connect(oracleAccounts[0])
        .sign(deal, network, user2CidList2, 1)
      tx = await tx.wait()

      let votes = await oracleInstance.getVotes(hashKey)
      //console.log('votes:', votes)

      expect(votes).to.equal(1)

      await expect(
        oracleInstance
          .connect(oracleAccounts[0])
          .signHash(deal, network, recipient, hashKey),
      ).to.be.revertedWith('you already signed this hash')

      await oracleInstance
        .connect(oracleAccounts[1])
        .signHash(deal, network, recipient, hashKey)

      votes = await oracleInstance.getVotes(hashKey)
      //console.log('votes:', await oracleInstance.getVotes(hashKey))

      expect(votes).to.equal(2)
    })
  })
})
