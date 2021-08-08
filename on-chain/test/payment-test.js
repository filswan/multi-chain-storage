const { expect } = require("chai");

describe("Payment gateway", function () {
  const chainlinkAddress = "0x514910771af9ca656af840dff83e8264ecf986ca";
  const wethAddress = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2";

  let paymentInstance;
  let accounts;
  let oracleInstance;
  before('deploy payment gateway contract', async()=>{
    accounts = await ethers.getSigners();

    console.log('deploy payment instance')
    const contract = await ethers.getContractFactory("SwanPayment");
    paymentInstance = await contract.deploy(accounts[0].address);
    await paymentInstance.deployed(); // have to wait block to include this transaction

    console.log('deploy oracle instance')
    const oracleContract = await ethers.getContractFactory("FilecoinOracle");
    oracleInstance = await oracleContract.deploy(accounts[0].address);
    await oracleInstance.deployed();

    console.log('set oracle')
    await paymentInstance.setOracle(oracleInstance.address);

    const tx = {
      to: wethAddress,
      value: ethers.utils.parseEther("0.5")
    };
  
    await accounts[0].sendTransaction(tx);
  
  });

  describe('payment workflow', async()=>{

    const cid = "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72";
    const recipient = "0xdD2FD4581271e230360230F9337D5c0430Bf44C0";

    let overrides = {
      // To convert Ether to Wei:
      value: ethers.utils.parseEther("10.0")     // ether in this case MUST be a string
  
      // Or you can use Wei directly if you have that:
      // value: someBigNumber
      // value: 1234   // Note that using JavaScript numbers requires they are less than Number.MAX_SAFE_INTEGER
      // value: "1234567890"
      // value: "0x1234"
  
      // Or, promises are also supported:
      // value: provider.getBalance(addr)
  };

    it("Test Lock payment", async function () {
      const tx = await paymentInstance.connect(accounts[0]).lockPayment({
        id: cid,
        minPayment: 10,
        lockedFee: 0,
        owner: accounts[0].address,
        recipient: recipient, //todo:
        deadline:0,
      _isExisted: true,
      }, overrides);

      await tx.wait();

      const result = await paymentInstance.getLockedPaymentInfo(cid);
      expect(accounts[0].address).to.equal(result.owner);
      expect(recipient).to.equal(result.recipient);
      expect(overrides.value).to.equal(result.lockedFee);
    });
  
    // it("Test Get lock payment info", async function () {
    //   const result = await paymentInstance.getLockedPaymentInfo(cid);
    //   expect(accounts[0].address).to.equal(result.owner);
    //   expect(recipient).to.equal(result.owner);
    // });

  });

  
});
