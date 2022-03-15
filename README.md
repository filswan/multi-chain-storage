# Multi Chain Storage Guide
[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on discord](https://img.shields.io/badge/join%20-discord-brightgreen.svg)](https://discord.com/invite/KKGhy8ZqzK)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)

- Join us on our [public discord channel](https://discord.com/invite/KKGhy8ZqzK) for news, discussions, and status updates. 
- [Check out our medium](https://filswan.medium.com) for the latest posts and announcements.
## Table of Contents
- [Functions](#Functions)
- [System Design](#System-Design)
- [Modules](#Modules)
- [Prerequisites](#Prerequisites)
- [Installation](#Installation)
- [After Installation](#After-Installation)
- [Configuration](#Configuration)
- [Payment Process](#Payment-Process)
- [Database Table Introduction](#Database-Table-Introduction)
- [Pay for Filecoin by Polygon](https://www.youtube.com/watch?v=c4Dvidz3plU)
- [License](#License)

## Functions
- Make payment from multi chain for filecoin storage
- Backup user's file to filecoin network
- Supports payment with tokens such as USDC on polygon
- Currently, USDC is supported for payment.

## System Design

![MCP-MCP Desgin](https://user-images.githubusercontent.com/8363795/143811916-f051ccce-f9b2-49eb-99ab-8da1a0d9f2f2.png)

## Modules
* [Token Swap](#Token-Swap)
* [Payment Module](#Payment-Module)
* [Swan Client API](https://github.com/filswan/go-swan-client)
* [DAO Signature](#DAO-Signature)
* [Data DAO](https://github.com/filswan/flink)
* [IPFS](https://docs.ipfs.io/)
* [Filecoin Storage](https://lotus.filecoin.io/docs/set-up/install/)
* [FLink](https://github.com/filswan/flink)

### Token Swap
1. Users pay USDC or other tokens, which are called user tokens, when uploading a file.
2. MCS uses FIL, which is called wrapped token, to pay when store data to filecoin network.
3. User tokens should be changed to wrapped tokens by this module and this step is called token exchange(swap).
4. Token exchange(swap) is done through Sushi Swap which is a DEX.

### Payment Module
1. After a file is uploaded, the money to be paid is estimated based on: 
   1. the average price of all the miners on the entire network
   2. file size
   3. duration
2. Then the estimated amount of money will be locked to the payment contract address, see [Configuration](#Configuration)
3. In unlock step, the amount paid to filcoin network by swan platform fil wallet, will be transfered to MCS payment receiver address, see [Configuration](#Configuration), and the overpayment part that is locked will be returned to user wallet

### DAO Signature
- If DAO detects that the file uploaded has been chained, it will trigger a signature operation

## Prerequisites
- OS: Ubuntu 20.04 LTS
- Mysql5.5+
- [Lotus Node](#Lotus-Node)
- [IPFS Client](https://docs.ipfs.io/install/)

### Lotus Node
- Lotus node is used for making car files and sending offline deals
- Install lotus node or lotus lite node in the same machine as MCS
- Lotus lite node is preferred since lotus full node is too heavy compared with lotus lite node 
- Lotus lite node depends on a lotus node, so ensure that a lotus node exists somewhere when using lotus lite node
#### Option:one: [install a lotus full node](https://lotus.filecoin.io/docs/set-up/install/)
#### Option:two: [install a lotus lite node](https://lotus.filecoin.io/docs/set-up/lotus-lite/#amd-and-intel-based-computers)

## Installation
### Option:one:  **Prebuilt package**: See [release assets](https://github.com/filswan/multi-chain-payment/releases)
```shell
wget https://github.com/filswan/multi-chain-storage/releases/tag/v1.0.1/install.sh
./install.sh
```

### Option:two:  Source Code
:bell:**go 1.16+** is required
```shell
git clone https://github.com/filswan/multi-chain-storage.git
cd multi-chain-storage
git checkout <release_branch>
./build_from_source.sh
```

## After Installation
- Before executing, you should check your configuration in `./build/config/config.toml` to ensure it is right.
```shell
vi ./build/config/config.toml
```
- Before executing, you should check your polygon configuration in `./build/config/polygon/config_polygon.toml` to ensure it is right.
```shell
vi ./build/config/polygon/config_polygon.toml
```
- Before executing, you should check your enviornment variable in `./.env` to ensure it is right.
```shell
vi ./.env
```
- After set your config and env variable in the related files, you can run `multi-chain-storage` in `./build` directory
```shell
./build/multi-chain-storage
```
### Note
- Logs are in directory `./logs`
- You can add `nohup` before `./multi-chain-storage` to ignore the HUP (hangup) signal and therefore avoid stop when you log out.
- You can add `>> mcs.log` in the command to let all the logs output to `mcs.log`.
- You can add `&` at the end of the command to let the program run in background.
- Such as:
```shell
nohup ./multi-chain-storage-0.2.1-rc1-unix >> mcs.log &   #After installation from Option 1
nohup ./build/multi-chain-storage >> ./build/mcs.log &    #After installation from Option 2
```



## Configuration

- **admin_wallet_on_polygon**: The wallet address used to execute contract methods on the polygon network and pay for gas

- **swan_payment_address_on_polygon**: The contract address of the swan payment gateway, used for lock and unlock user fees on polygon

- **file_coin_wallet**: The wallet to pay file storage on filecoin network

#### [swan_api]

swan_api section defines the token used for connecting with Swan platform.

- **api_url**: Default: "https://api.filswan.com"
- **api_key & access_token**: Acquire from [Filswan](https://www.filswan.com) -> "My Profile"->"Developer Settings". You
  can also check the [Guide](https://nebulaai.medium.com/how-to-use-api-key-in-swan-a2ebdb005aa4)

#### [lotus]

- **api_url**: Url of lotus client web api, such as: **http://[ip]:[port]/rpc/v0**, generally the [port] is **1234**.
  See [Lotus API](https://docs.filecoin.io/reference/lotus-api/)
- **access_token:** Access token of lotus market web api. When market and miner are not separate, it is also the access
  token of miner access token. See [Obtaining Tokens](https://docs.filecoin.io/build/lotus/api-tokens/#obtaining-tokens)

#### [ipfs_server]

ipfs-server is used to upload and download generated Car files.
- **upstream_url**: upload api url of IPFS Server
- **download_stream_url**: download api url of IPFS Server

#### [swan_task]

- **dir_deal**: Output directory for saving generated Car files and CSVs
- **verified_deal**: [true/false] Whether deals in this task are going to be sent as verified
- **fast_retrieval**: [true/false] Indicates that data should be available for fast retrieval
- **start_epoch_hours**: start_epoch for deals in hours from current time
- **expired_days**: expected completion days for storage provider sealing data
- **max_price**: Max price willing to pay per GiB/epoch for offline deal
- **min_price**: Min price willing to pay per GiB/epoch for offline deal
- **generate_md5**: [true/false] Whether to generate md5 for each car file, note: this is a resource consuming action
- **relative_epoch_from_main_network**:  The difference between the block height of filecoin's mainnet and testnet. If
  the filecoin testnet is linked, this configuration item is set to -858481, and if the link is the mainnet, it is set
  to 0

### config_polygon.toml

Currently, USDC is supported for payment. Take polygon network as an example to introduce configuration items

#### [polygon_mainnet_node]

- **rpc_url**: the polygon network rpc url
- **payment_contract_address**:  swan payment gateway address on polygon
- **contract_function_signature**:  swan payment gateway's lock payment event's function signature on polygon
- **dao_swan_oracle_address**:  swan dao address on polygon
- **dao_event_function_signature**:  swan dao's signature event's function signature on polygon
- **scan_step**:  the number of blocks scanned per scan
- **start_from_blockNo**:  scan data from this block number
- **start_from_blockNo**:  the time between each scan of the blockchain

### .env
- **privateKeyOnPolygon**: private key of the wallet used to execute contract methods on the polygon network and pay for gas

## The Payment Process

1. Users upload a file they want to backup to filecoin network
2. MCS create car file for the uploaded source file, and create task on swan platform, by calling Swan Client API, see [Swan Client](https://github.com/filswan/go-swan-client)
3. User pay currencies we support to send tokens to our payment contract address, see [Configuration](#Configuration)
4. Market Matcher allocate miners for the created car file
5. MCS send deals by calling Swan Client API, see [Swan Client](https://github.com/filswan/go-swan-client)
6. MCS Scan Scheduler module scan the deal info from lotus, and update status locally
7. When DAO organization find the deal succeeds, and then they will sign to agree to unlock the user's payment.
8. After more than half of the dao agree, MCS will unlock the user's payment, release the moeny spent on send deal to MCS payment receiver address, see [Swan Client](https://github.com/filswan/go-swan-client), and return the remaining money to user wallet address used when pay in step 2.

## Database table description
You can get db table ddl sql script in $GOPATH/src/payment-bridge/script/dbschema.sql <br>
There are two tables you need to initialize the data before you can use it<br>

### system_config_param
|column                 |description       |
|param_key              |param_value       |
|-----------------------|------------------|
|SWAN_PAYMENT_CONTRACT_ADDRESS     |swan payment gateway contract address                     |
|LOCK_TIME                         |time that user's token will be locked time                |
|RECIPIENT                         |admin wallet address, used to receive tokens paid by users|
|PAY_GAS_LIMIT                     |max gas limit                                             |
|USDC_ADDRESS                      |usdc address                                              |
|MINT_CONTRACT                     |mint contract address                                     |

### dao_info
### system_config_param
|column                 |description       |
|-----------------------|------------------|
|ID                     |primary key of table   |
|DAO_NAME               |dao's name,required    |
|DAO_ADDRESS            |dao's address,required |
|ORDER_INDEX            |dao display order      |
|DESCRIPTION            |description            |
|CREATE_AT              |create time            |

### Run Payment Bridge as system service
Before running the playbook to start payment bridge as a system service, please first change the following line in `$GOPATH/src/payment-bridge/script/run_services/payment_bridge.service` to the actual path of payment bridge executable:
```
ExecStart=/home/filswan/payment-bridge/build/payment-bridge
```
Now we can run Payment Bridge as a system service by executing the following command in shell and entering sudo password when prompt:
```bash
cd $GOPATH/src/payment-bridge/script/run_services
ansible-playbook run_payment_bridge_service.yaml --ask-become-pass -vvv
```

## License

[MIT](https://github.com/filswan/multi-chain-payment/blob/main/LICENSE)
