[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on Slack](https://img.shields.io/badge/slack-filswan.slack.com-green.svg)](https://filswan.slack.com)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)

# MCP Contract


## Architecture Overview

![Architecture Overview!](./docs/image/architecture.png "Contract Architecture Overview")


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
1.  mcp_platform calls refund method of mcp_payment_contract, and mcp_payment_contract returns remain tokens back to user.  




#### How to caculate required locked tokens
After a file is uploaded, the cost of storage service is estimated based on following params   
   - average price of all the miners on the entire (filecoin?) network.  
   - file size  
   - number of copies
   - duration  


## Environment Info

### Mumbai Testnet Blockchain Network
RPC URL: **https://matic-mumbai.chainstacklabs.com**      
Chain ID: **80001**
#### Stable Env
##### Contract List
|Contract   |  Address |
|---|---|
| mcp_payment_contract  | 0x12EDC75CE16d778Dc450960d5f1a744477ee49a0  |
| dao_contract  | 0xe3262c0848b0cc5cd43df7139103f1fbf26558cc  |
| flink_service_contract  | 0xcE9A9e594db39dCD449E392d68F60959533c0D75  |
| DEX pool price_feed_address  | 0xe8a67994c114e0c17E1c135d0CB599a2394f1505  |
| usdc token address  | 0xe11A86849d99F524cAC3E7A0Ec1241828e332C62  |
| wFil token address  | 0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055  |

##### Wallet List
... add wallet addresses


#### Development Env
##### Contract List

|Contract   |  Address |
|---|---|
| mcp_payment_contract  | 0x24B9c56BB6419f4c5AE6a63Fd64dE0dCFA1841F1  |
| dao_contract  | 0x9208C2B417Ec2699454843A06A5E49fA6dd88422  |
| flink_service_contract  | 0xB6312D719F6B496647703c81F6965EF38bF58B8D  |
| DEX pool price_feed_address  | 0xe8a67994c114e0c17E1c135d0CB599a2394f1505  |
| usdc token address  | 0xe11A86849d99F524cAC3E7A0Ec1241828e332C62  |
| wFil token address  | 0x97916e6CC8DD75c6E6982FFd949Fc1768CF8c055  |

DEX pool, usdc, wFil addresses are the same, since we are using the same network
##### Wallet List
DAO addresses list  
0x6d2e5279b106843f6E924194401B50e6e27FE12a(unlock wallet)  
0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5  
0xeA2bf08288bbfB0d3DBf534f35af32bF2c6E5e45  

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
Read/Write data onto blockchain network. Using [blockchain explorer](https://mumbai.polygonscan.com/) to get on-chain data.
