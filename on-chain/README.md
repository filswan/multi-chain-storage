[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on Slack](https://img.shields.io/badge/slack-filswan.slack.com-green.svg)](https://filswan.slack.com)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)

# MCS Contract

## Architecture Overview

![Architecture Overview!](./docs/image/architecture.png 'Contract Architecture Overview')

## Main Work Flow

### Business Logic Flow

1. user uploads file,
2. user_wallet calls lockTokenPayment method of mcp_payment_contract, locking usdc token into contract
3. usdc token deposit event is generated [on-chain](#On-Chain)
4. mcp_platform gets deposit info from [on-chain](#On-Chain)
5. mcp_platform packs multiple files as a car file and calls Storage Service

#### How to caculate required locked tokens

After a file is uploaded, the cost of storage service is estimated based on following params

- average price of all the miners on the entire (filecoin?) network.
- file size
- number of copies
- duration

## Environment Info

### Goerli Arbitrum Testnet Blockchain Network

RPC URL: **https://goerli-rollup.arbitrum.io/rpc**  
Chain ID: **421613**

##### Contract List

| Contract                     | Address                                    |
| ---------------------------- | ------------------------------------------ |
| usdc token address           | 0x1FdDc28136a57CF1713E5Fc416953687Fe2Ba339 |
| renFil token address         | 0x2da6871c3Fd3598Bc2249901df01139bDfFe815e |
| usdc/renFil DEX pair address | 0xFF2278d5C5C5761d951Fd87bA077947Aa31d737f |
| nft contract                 | 0xE9CDfa20366233443eabd8B8061E7dF4b4bD2DcB |
| price feed contract          | 0x044CEB98dB2993779b9dD71afA4D12B378102f04 |
| faucet contract              | 0x6df9600d9e41068cA6E5F7D3AC729609B4eCf7f5 |
| payment contract             | 0xb157536Aa52a2bfb76b1B75969E571846EF7A009 |  |

## Development

#### generate go bind files

```
cd /on-chain/contracts
abigen --abi ./abi/SwanPayment.json --pkg goBind --type SwanPayment --out ../goBind/SwanPayment.go
```

## FAQ

...

## Glossary

### On Chain

Read/Write data onto blockchain network. Using [blockchain explorer](https://goerli.arbiscan.io/) to get on-chain data.
