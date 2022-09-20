require('@openzeppelin/hardhat-upgrades');
require("@nomiclabs/hardhat-waffle");
require('dotenv').config();

require('hardhat-abi-exporter');
require("hardhat-gas-reporter");


// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */


module.exports = {
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
      forking: {
        url: "https://eth-mainnet.alchemyapi.io/v2/FjSVplPnI6oBEN8dUcwVfr_sMZTpOuUx",
        blockNumber: 13247494
      }
    },

    // hardhat: {
    //   forking:{
    //     url: "https://polygon-mumbai.g.alchemy.com/v2/shs801xBZD6sqSpQMf7-UE_35ZicdRyL",
    //     blockNumber: 17357506
    //   }
    // },

    localhost: {
      url: "http://localhost:8545",
      accounts: [
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
        "0xde9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0",
        "0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e",
        "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
        "0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
        "bb2ab0f910c23824b521af2920a7779077bd261133de5993802d22e9a2cfbba4"],
    },
    ethereum: {
      url: `https://rpc.flashbots.net`,
      accounts: [process.env.ownerPK, process.env.key2, process.env.key3,process.env.key4,process.env.key5]
    },
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.infuraKey}`,
      accounts: [process.env.ownerPK, process.env.key2, process.env.key3,process.env.key4]
    },
    kovan: {
      url: `https://kovan.infura.io/v3/${process.env.infuraKey}`,
      accounts: [process.env.ownerPK]
    },
    matictest: {
      url: "https://polygon-mumbai.g.alchemy.com/v2/h0hrnucL6avUmswe31DKRNhOwI14SRbs", // https://forum.openzeppelin.com/t/hardhat-upgrade-failing-with-hh110-on-matic-testnet/12911
      accounts: [process.env.ownerPK, process.env.key2, process.env.key3,process.env.key4],
      gasPrice: 10000000000,
      gasLimit: 9999999
    },
    bsctest: {
      url: "https://data-seed-prebsc-1-s1.binance.org:8545",
      accounts: [process.env.ownerPK]
    },
    polygon: {
      url: `https://polygon-mainnet.g.alchemy.com/v2/${process.env.alchemyKey}`,
      accounts: [process.env.ownerPK, process.env.key2, process.env.key3,process.env.key4,process.env.key5]
    },

    
  },
  solidity: {
    compilers: [
      {
        version: "0.6.6",
      },
      {
        version: "0.8.4",
        settings: {
          optimizer: {
            enabled: true,
            runs: 20000
          }
        }
      },
    ],
    settings: {
      optimizer: {
        enabled: true,
        runs: 20000
      }
    }

  },
  gasReporter: {
    enabled: (process.env.REPORT_GAS) ? true : false
  },
  abiExporter: {
    path: './contracts/plain-abi',
    runOnCompile: true,
    clear: true,
    flat: true,
    only: [':FilswanOracle$'],
    spacing: 2,
    pretty: true,
  },
};
