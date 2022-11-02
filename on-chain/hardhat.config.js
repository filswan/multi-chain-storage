require('@openzeppelin/hardhat-upgrades')
require('@nomiclabs/hardhat-waffle')
require('dotenv').config()

require('hardhat-abi-exporter')
require('hardhat-gas-reporter')

require('@nomiclabs/hardhat-etherscan')

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task('accounts', 'Prints the list of accounts', async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners()

  for (const account of accounts) {
    console.log(account.address)
  }
})

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

module.exports = {
  defaultNetwork: 'hardhat',
  networks: {
    polygon: {
      url: 'https://polygon-rpc.com',
      accounts: [process.env.PERSONAL_KEY],
    },
    mumbai: {
      url: process.env.MUMBAI_RPC,
      accounts: [process.env.PRIVATE_KEY],
    },
    tbnc: {
      url: 'https://data-seed-prebsc-1-s1.binance.org:8545/',
      accounts: [process.env.PRIVATE_KEY],
    },
  },
  solidity: {
    compilers: [
      { version: '0.8.2' },
      {
        version: '0.8.7',
      },
      {
        version: '0.6.6',
      },
      {
        version: '0.8.4',
        settings: {
          optimizer: {
            enabled: true,
            runs: 20000,
          },
        },
      },
    ],
    settings: {
      optimizer: {
        enabled: true,
        runs: 20000,
      },
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS ? true : false,
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
  etherscan: { apiKey: process.env.POLYGONSCAN_API_KEY },
}
