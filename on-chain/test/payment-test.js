const { expect } = require("chai");

describe("Payment gateway", function () {

  let paymentInstance;
  let accounts;
  before('deploy payment gateway contract', async()=>{
    accounts = await ethers.getSigners();

    const contract = await ethers.getContractFactory("SwanPayment");
    paymentInstance = await contract.deploy();
    // await sp.deployed();
  });

  describe('payment workflow', async()=>{

    const cid = "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72";
    const recipient = "0xdD2FD4581271e230360230F9337D5c0430Bf44C0"
    it("Test Lock payment", async function () {
      const result = await paymentInstance.lockPayment({
        id: cid,
        minPayment: 100,
        recipient: recipient, //todo:
      });
      expect(result).to.equal(true);
    });
  
    it("Test Get lock payment info", async function () {
      const result = await paymentInstance.getLockedPaymentInfo(cid);
      expect(result.owner).to.equal(accounts[0].address);
      expect(result.owner).to.equal(recipient);
    });

  });

  
});
