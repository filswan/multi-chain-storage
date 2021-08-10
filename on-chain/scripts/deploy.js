async function main() {
    const [deployer] = await ethers.getSigners();
  
    console.log("Deploying contracts with the account:", deployer.address);
  
    console.log("Account balance:", (await deployer.getBalance()).toString());
  
    const contract = await ethers.getContractFactory("SwanPayment");
    const paymentInstance = await contract.deploy(deployer.address);
    console.log("paymentInstance address:", paymentInstance.address);
    await paymentInstance.deployed();

    console.log('deploy oracle instance')
    const oracleContract = await ethers.getContractFactory("FilecoinOracle");
    oracleInstance = await oracleContract.deploy(deployer.address);
    await oracleInstance.deployed();
    console.log("oracleInstance address:", oracleInstance.address);
    await oracleInstance.setDAOUsers([deployer.address]);

    console.log('set oracle')
    await paymentInstance.setOracle(oracleInstance.address);

  }
  
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });