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
6. DAO memebers check storage service status (this step is offline)
7. DAO memebers sign storage status info(dealid, mcp_payment_receiver_address) on-chain, by calling dao_contract
8. Once signatures from DAO memeber reach threshold,
9. mcp_platform calls unlockCarPayment method of mcp_payment_contract  
   9.1 mcp_payment_contract reads payment info from flink_service_contract,please make sure your caller wallet has engouh link for oracle service  
   9.2 mcp_payment_contract gets wFil/USDC price from DEX pool price_feed_address  
   9.3 mcp_payment_contract calculates cost of storage service and pays it to mcp_platform
10. mcp_platform calls refund method of mcp_payment_contract, and mcp_payment_contract returns remain tokens back to user.

#### How to caculate required locked tokens

After a file is uploaded, the cost of storage service is estimated based on following params

- average price of all the miners on the entire (filecoin?) network.
- file size
- number of copies
- duration

## Environment Info

### Mumbai Testnet Blockchain Network

RPC URL: **https://polygon-rpc.com/**  
Chain ID: **137**

#### Development Env multi miner version

##### Contract List

| Contract           | Address                                    |
| ------------------ | ------------------------------------------ |
| USDC               | 0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174 |
| ChainlinkPriceFeed | 0xFC8B846fEd57579F91973F0561a08a235A39a8dA |
| FilinkConsumer     | 0x2Bf5dBde4Fdd30de18b36405CF587044172ffD33 |
| FilswanOracle      | 0x2621BB3140D8914806E977F7e6035B468675304D |
| SwanPayment        | 0xA1f32c758c4324cC3070A3AA107C4dC7DdFe1a6f |
| SwanNFT            | 0xA6787587159c017AD83fe28e746FCFAE0DD91383 |

##### Wallet List

DAO addresses list  
0xd44CDe0f3BeEF47Af166fC763b52977A43bf8Fe7 (unlock wallet)  
0xa1c24757Da070E62b12bf2c762213c35FA30Fae5  
0x74D5Ef1C805EA2D42ee3163fA312c9e156635C68

## Development

#### generate go bind files

```
cd /on-chain/contracts
abigen --abi ./abi/FilswanOracle.json --pkg goBind --type FilswanOracle --out ../goBind/FilswanOracle.go
abigen --abi ./abi/SwanPayment.json --pkg goBind --type SwanPayment --out ../goBind/SwanPayment.go
```

## FAQ

...

## Glossary

### On Chain

Read/Write data onto blockchain network. Using [blockchain explorer](https://polygonscan.com/) to get on-chain data.
