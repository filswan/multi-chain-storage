const { expect } = require("chai");

describe("Payment gateway", function () {
  // const chainlinkAddress = "0x514910771af9ca656af840dff83e8264ecf986ca";
  const wethAddress = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2";

  let accounts;

  let paymentInstance;
  let oracleInstance;

  before('Deploy payment gateway contract', async () => {
    accounts = await ethers.getSigners();

    const daoAddressList = accounts.map(a => a.address);

    console.log('deploy payment instance')
    const contract = await ethers.getContractFactory("SwanPayment");
    paymentInstance = await contract.deploy(accounts[0].address);
    await paymentInstance.deployed(); // have to wait block to include this transaction

    console.log('deploy oracle instance')
    const oracleContract = await ethers.getContractFactory("FilecoinOracle");
    oracleInstance = await oracleContract.deploy(accounts[0].address);
    await oracleInstance.deployed();

    await oracleInstance.setDAOUsers(daoAddressList);

    console.log('set oracle')
    await paymentInstance.setOracle(oracleInstance.address);

    const tx = {
      to: wethAddress,
      value: ethers.utils.parseEther("0.5")
    };

    await accounts[0].sendTransaction(tx);

  });

  describe('Payment oracle workflow', async () => {
    // const accounts = await ethers.getSigners();


    const zero = ethers.BigNumber.from("0");
    const actualPay20Native = ethers.BigNumber.from("20000000000000000000");

    it("Test only some dao users update payment info", async function () {
      const cid_zero = "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw00";
      for (let i = 0; i < accounts.length - 1; i++) {
        const tx = await oracleInstance.connect(accounts[i]).updatePaymentInfo(cid_zero, actualPay20Native);
        await tx.wait();
      }

      const result = await oracleInstance.getPaymentInfo(cid_zero);
      console.log("current payment is ", result.toString());
      expect(zero).to.equal(result);
    });

    it("Test all dao users update payment info", async function () {
      const cid_20 = "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw20";
      for (let i = 0; i < accounts.length; i++) {
        const tx = await oracleInstance.connect(accounts[i]).updatePaymentInfo(cid_20, actualPay20Native);
        await tx.wait();
      }

      const result = await oracleInstance.getPaymentInfo(cid_20);
      console.log("current payment is ", result.toString());
      expect(actualPay20Native).to.equal(result);
    });


    it("Test Lock payment", async function () {
      const cid = "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72";
      const minPay10Native = ethers.BigNumber.from("10000000000000000000");

      const payer = accounts[3];
      const fileswanRecipient = accounts[4];

      const fee = {
        // To convert Ether to Wei:
        value: ethers.utils.parseEther("89.0")     // ether in this case MUST be a string

        // Or you can use Wei directly if you have that:
        // value: someBigNumber
        // value: 1234   // Note that using JavaScript numbers requires they are less than Number.MAX_SAFE_INTEGER
        // value: "1234567890"
        // value: "0x1234"

        // Or, promises are also supported:
        // value: provider.getBalance(addr)
      };

      const tx = await paymentInstance.connect(payer).lockPayment({
        id: cid,
        minPayment: minPay10Native,
        lockTime: 86400000, // one day
        recipient: fileswanRecipient.address, //todo:
      }, fee);

      await tx.wait();

      const result = await paymentInstance.getLockedPaymentInfo(cid);
      expect(payer.address).to.equal(result.owner);
      expect(fileswanRecipient.address).to.equal(result.recipient);
      expect(fee.value).to.equal(result.lockedFee);
    });

    // it("Test Unlock payment before oracle updates info", async function () {
    //   const tx = await paymentInstance.connect(payer).unlockPayment(cid);

    //   await tx.wait();

    //   const result = await paymentInstance.getLockedPaymentInfo(cid);
    //   expect(payer.address).to.equal(result.owner);
    //   expect(fileswanRecipient.address).to.equal(result.recipient);
    //   expect(fee.value).to.equal(result.lockedFee);
    // });

    // it("Test Unlock payment after oracle updated info", async function () {
    //   const tx = await paymentInstance.connect(payer).unlockPayment(cid);

    //   await tx.wait();

    //   const result = await paymentInstance.getLockedPaymentInfo(cid);
    //   expect(payer.address).to.equal(result.owner);
    //   expect(fileswanRecipient.address).to.equal(result.recipient);
    //   expect(fee.value).to.equal(result.lockedFee);
    // });

  });




});
