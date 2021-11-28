
const { ethers, upgrades } = require("hardhat");

const uniswapv2FactoryABI = require('@uniswap/v2-core/build/IUniswapV2Factory.json').abi;
const sushiswapFactoryAddress = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4";

const overrides = {
  gasLimit: 9999999
}


async function main() {

  const [deployer] = await ethers.getSigners();

  console.log("Deploying contracts with the account:", deployer.address);

  console.log("Account balance:", (await deployer.getBalance()).toString());

  const swanOracleContract = await ethers.getContractFactory("FilswanOracle");


  // threshold 2, meaning 2 votes are required to pass
  // const swanOracleInstance = await upgrades.deployProxy(swanOracleContract, [deployer.address, 2], overrides);
 
  // await swanOracleInstance.deployed();
  // console.log(`swanOracleInstance address: ${swanOracleInstance.address}`);

  const swanOracleAddress = "0xe3262c0848b0cc5cd43df7139103f1fbf26558cc";


  const usdcAddress = "0xe11A86849d99F524cAC3E7A0Ec1241828e332C62";
  const wFilAddress = "0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055";
  const pairAddress = "0x74038ed7D891A043d4aF41FeA242ED01914c2636";

  // const priceFeedContract = await ethers.getContractFactory("PriceFeed");
  // // todo: add tokenA and tokenB address
  // const priceOracleFeedInstance = await priceFeedContract.deploy(pairAddress, 0, overrides);
  // await priceOracleFeedInstance.deployed();
  // console.log(`priceOracleFeedInstance address: ${priceOracleFeedInstance.address}`);

  // const one = "1000000000000000000"; 

  // const price = await priceOracleFeedInstance.consult(usdcAddress, one);
  // console.log(`price: ${price}`);


  const priceFeedOracleAddress = "0xe8a67994c114e0c17E1c135d0CB599a2394f1505";
  const filinkOracleAddress = "0xcE9A9e594db39dCD449E392d68F60959533c0D75";

  const swanPaymentContract = await ethers.getContractFactory("SwanPayment");

  const swanPaymentInstance = await upgrades.deployProxy(swanPaymentContract, [deployer.address, usdcAddress, swanOracleAddress, priceFeedOracleAddress, filinkOracleAddress], overrides);
  await swanPaymentInstance.deployed();
  console.log(`swanPaymentInstance address: ${swanPaymentInstance.address}`)

  // const swanPaymentAddress  = "0xABeAAb124e6b52afFF504DB71bbF08D0A768D053";

  // swanPaymentInstance address: 0xABeAAb124e6b52afFF504DB71bbF08D0A768D053

  // const contract = await ethers.getContractFactory("SwanPayment");
  // const paymentInstance = await contract.deploy(deployer.address);
  // console.log("paymentInstance address:", paymentInstance.address);
  // await paymentInstance.deployed();

  // console.log('deploy oracle instance')
  // const oracleContract = await ethers.getContractFactory("FilecoinOracle");
  // oracleInstance = await oracleContract.deploy(deployer.address);
  // await oracleInstance.deployed();
  // console.log("oracleInstance address:", oracleInstance.address);
  // await oracleInstance.setDAOUsers([deployer.address]);

  // console.log('set oracle')
  // await paymentInstance.setOracle(oracleInstance.address);




}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });