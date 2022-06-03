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
- [Database](#Database)
- [Installation](#Installation)
- [After Installation](#After-Installation)
- [Configuration](#Configuration)
- [Work Process](#Work-Process)
- [Pay for Filecoin by Polygon](https://www.youtube.com/watch?v=c4Dvidz3plU)
- [License](https://github.com/filswan/multi-chain-storage/blob/main/LICENSE)

## Functions
- Make payment from multi chain for filecoin storage
- Backup user's file to filecoin network
- Supports payment with tokens such as USDC on polygon
- Currently, USDC is supported for payment.

## System Design

![MCS Desgin](https://user-images.githubusercontent.com/8363795/143811916-f051ccce-f9b2-49eb-99ab-8da1a0d9f2f2.png)


## Modules
* [Token Swap](#Token-Swap)
* [Payment Module](#Payment-Module)
* [Swan Client API](https://github.com/filswan/go-swan-client)
* [DAO Signature](#DAO-Signature)
* [Data DAO](https://github.com/filswan/flink)
* [IPFS](https://docs.ipfs.io/)
* [Filecoin Storage](https://lotus.filecoin.io/docs/set-up/install/)

### Token Swap
1. Users pay USDC or other tokens, which are called user tokens, when uploading a file.
2. MCS uses FIL, which is called wrapped token, to pay when store data to filecoin network.
3. User tokens should be changed to wrapped tokens by this module and this step is called token exchange(swap).
4. Token exchange(swap) is done through Sushi Swap which is a DEX.

### Payment Module
1. After a file is uploaded, the money to be paid is estimated based on: 
   1. the average price of all the miners on the entire network
   2. file size
   3. storage copy number
   4. duration
2. Then the estimated amount of money will be locked to the payment contract address, see [Configuration](#Configuration)
3. In unlock step, the amount pay to filcoin network by swan platform fil wallet, will be transfered to mcs payment receiver address, see [Configuration](#Configuration)
4. In refund step, the overpayment part that is locked will be returned to user wallet

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

## Database
- Please see schema create script in `../script/create_table.sql`
- Before installation, please create database and related tables using above script file

## Installation
### Option:one:  **Prebuilt package**: See [release assets](https://github.com/filswan/multi-chain-storage/releases)
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
- Before executing, you should check your configuration in `~/.swan/mcs/config.toml` to ensure it is right.
```shell
vi ~/.swan/mcs/config.toml
```
- Before executing, you should check your enviornment variable in `~/.swan/mcs/.env` to ensure it is right.
```shell
vi ~/.swan/mcs/.env
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

### config.toml
- **port**: Web api port
- **release**: When work in release mode: set this to true, otherwise to false and enviornment variable GIN_MODE not to release
- **swan_platform_fil_wallet**: The wallet address used to pay on the filecoin network
- **flink_url**: Deals data can be searched from here
- **filecoin_network**: filecoin_calibration or filecoin_mainnet

#### [database]
- **db_host**: Host to connect mcs database
- **db_port**: Port to connect mcs database
- **db_schema_name**: MCS database name, see [Database](#Database)
- **db_username**: Username to connect mcs database
- **db_password**: Password to connect mcs database
- **db_args**: Use default value `charset=utf8mb4&parseTime=True&loc=Local`

#### [swan_api]
- **api_url**: Swan API address: `https://go-swan-server.filswan.com`.
- :bangbang:**api_key**: Your Swan API key. Acquire from [Swan Platform](https://console.filswan.com/#/dashboard) -> "My Profile"->"Developer Settings".
- :bangbang:**access_token**: Your Swan API access token. Acquire from [Swan Platform](https://console.filswan.com/#/dashboard) -> "My Profile"->"Developer Settings".

#### [lotus]
- **client_api_url**:  Url of lotus client web api, such as: `http://[ip]:[port]/rpc/v0`, generally the `[port]` is `1234`.
- **client_access_token**:  Access token of lotus client web api with admin access right. Get it from lotus node by command `lotus auth create-token --perm admin`.

#### [ipfs_server]
- **download_url_prefix**: Ipfs server url prefix, such as: `http://[ip]:[port]`. Store car files for downloading by storage provider. Car file url will be `[download_url_prefix]/ipfs/[file_hash]`
- **upload_url_prefix**: Ipfs server url for uploading files, such as `http://[ip]:[port]`

#### [swan_task]
- **dir_deal**: Output directory for saving generated Car files and CSVs
- **description**: Task description
- **curated_dataset**: Task dataset
- **max_price**: Max price willing to pay per GiB/epoch for offline deal
- **expired_days**: Expected completion days for storage provider sealing data
- **verified_deal**: [true/false] Whether deals in this task are going to be sent as verified
- **fast_retrieval**: [true/false] Indicates that data should be available for fast retrieval
- **start_epoch_hours**: Start epoch for deals in hours from current time
- **min_file_size**: Source files size lower limit when merge them to a car file

[schedule_rule]
- **unlock_interval_second**: Job running interval, unit: second, default: 300
- **create_task_interval_second**: Job running interval, unit: second, default: 120
- **send_deal_interval_second**: Job running interval, unit: second, default: 180
- **scan_deal_status_interval_second**: Job running interval, unit: second, default: 300
- **scan_polygon_interval_second**: Job running interval, unit: second, default: 1

#### [polygon]
- **polygon_rpc_url**: Your polygon network rpc url
- **payment_contract_address**:  Swan payment gateway address on polygon to lock money
- **payment_recipient_address**:  MCS wallet address to receive money from unlock operation
- **dao_contract_address**:  Swan dao address on polygon, to receive dao signatures
- **mint_contract_address**:  Swan mint address on polygon
- **sushi_dex_address**:  Sushi address on polygon
- **usdc_wFil_pool_contract**:  Address to get exchange rate between uscs and wFil from sushi on polygon
- **gas_limit**: Gas limit for transaction
- **lock_time**: Lock days defined in smart contract, it is 6 now
- **pay_multiply_factor**:
- **unlock_interval_minute**: Unlock interval in minutes between 2 unlock operations, in cannot be less than 1
- **dao_unlock_interval_block**: Interval block between unlock & last dao signature after reaching dao threshold
- **scan_polygon_block_step**: How many blocks to scan each time, default 1000

### .env
- **privateKeyOnPolygon**: Private key of the wallet used to execute contract methods on the polygon network and pay for gas

## Work Process

1. Users upload a file they want to backup to filecoin network
2. User pay currencies we support to send tokens to our payment contract address, see [Configuration](#Configuration)
3. MCS writes the transaction info to our system
4. MCS scan those source files uploaded and paid but not yet created to car files, and then do the following steps:
   1. compute the max price for each source file, based on the source file size, token paid, and exchange rate betwee USDC and wFil
   2. if the scanned source file size sum is more than `[swan_task].min_file_size`, see [Configuration](#Configuration), or the earliest source file to be merged to car file is more 1 day ago, then MCS will do the following steps by calling Swan Client API, see [Swan Client](https://github.com/filswan/go-swan-client)
      1. create car files, use the minimum max price among the source files to be merged as the max price for the whole car file
      2. upload car files
      3. create task on swan platform
5. Market Matcher allocate miners for the car file created in last step
6. MCS send deals by calling Swan Client API, see [Swan Client](https://github.com/filswan/go-swan-client)
7. MCS Scan Scheduler module scan the deal info from lotus
8. When DAO organization find the deal succeeds, and then they will sign to agree to unlock the user's payment.
9. After more than half of the dao agree and after 1 minute later of the last DAO signature, MCS will unlock the user's payment, release the moeny spent on send deal to mcs payment receiver address, see [Swan Client](https://github.com/filswan/go-swan-client)
10. After all deals of a car file are unlocked, MCS refund the remaining money to user wallet address used when pay in step 2.

