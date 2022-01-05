# Multi Chain Payment Guide
[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on Slack](https://img.shields.io/badge/slack-filswan.slack.com-green.svg)](https://filswan.slack.com)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)

- Join us on our [public Slack channel](https://www.filswan.com/) for news, discussions, and status updates.
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
- [Other Topics](#Other-Topics)
- [License](#License)

## Functions
- Make payment from multi chain for filecoin storage
- Backup user's file to filecoin network
- Supports payment with tokens such as USDC on polygon

## System Design

![MCP-MCP Desgin](https://user-images.githubusercontent.com/8363795/143811916-f051ccce-f9b2-49eb-99ab-8da1a0d9f2f2.png)

## Modules

* [Token Swap](#Token-Swap)
* [Payment Lock](#Payment-Lock)
* [DAO Organization](#DAO-Organization)
* [DAO Signature](#DAO-Signature)
* [Unlock Payment](#Unlock-Payment)
* [Data DAO](https://github.com/filswan/flink)
* IPFS/Filecoin Storage

### Token Swap
- Users pay USDC or other tokens, which are called user tokens, when uploading a file.
- MCP uses FIL, which is called wrapped token, to store data to filecoin network.
- User tokens should be changed to wrapped token by this module and this step is called token exchange(swap).
- Token exchange(swap) is done through Sushi Swap which is a DEX.

### Payment Lock
- After a file is uploaded, the money to be paid is estimated based on the average price of all the miners on the entire network.
- Then the estimated amount of money will be locked.
- The overpayment part that is locked will be returned through the unlock operation later

### DAO Organization
- If Dao detects that the file uploaded by the user has been chained, it will trigger a signature operation

### DAO Signature
- If Dao detects that the file uploaded by the user has been chained, it will trigger a signature operation

### Unlock Payment
- When more than half of the dao have been signed, the unlock operation will be triggered
- The part that needs to be paid will be deducted from the locked token
- The remaining token will be returned to the user

## Prerequisites
- OS: Ubuntu 20.04 LTS
- Ansible: Version 2.12+
- Python 3.6+
- Golang1.16 (minimum version)
- Mysql5.5
- [Lotus Node](#Lotus-Node)
- [IPFS Client](https://docs.ipfs.io/install/)

If python3 or ansible is not yet installed in the system, please run `install_pre-requisite.sh` script first
You can get this script from our payment-bridge source code which will be introduced below 

### Lotus Node
- Lotus node is used for making car files and sending offline deals
- Install lotus node or louts lite node in the same machine as MCP
- Lotus full node is too heavy compared with lotus lite node, so lotus lite node is preferred
- Lotus lite node depends on a lotus node, so ensure that a lotus node exists somewhere when using lotus lite node
#### Option:one: [install a lotus full node](https://lotus.filecoin.io/docs/set-up/install/)
#### Option:two: [install a lotus lite node](https://lotus.filecoin.io/docs/set-up/lotus-lite/#amd-and-intel-based-computers)

## Installation
### Option:one:  **Prebuilt package**: See [release assets](https://github.com/filswan/multi-chain-payment/releases)
```shell
wget https://github.com/filswan/multi-chain-payment/releases/tag/v1.0.1/install.sh
./install.sh
```

### Option:two:  Source Code
:bell:**go 1.16+** is required
```shell
git clone https://github.com/filswan/multi-chain-payment.git
cd multi-chain-payment
git checkout <release_branch>
./build_from_source.sh
```

## After Installation
- Before executing, you should check your configuration in `~/.swan/mcp/config.toml` to ensure it is right.
```shell
vi ~/.swan/mcp/config.toml
```
- To pay on polygon network, you need to check your configuration in `~/.swan/mcp/config_polygon.toml` to ensure it is right.
```shell
vi ~/.swan/mcp/config_polygon.toml
```
- To pay on goerli network, you need to check your configuration in `~/.swan/mcp/config_goerli.toml` to ensure it is right.
```shell
vi ~/.swan/mcp/config_goerli.toml
```
- After set your config in the related config files, you can run `multi-chain-payment` in `./build`
```shell
cd build
./multi-chain-payment
```
### Note
- Logs are in directory ./logs
- You can add `nohup` before `./multi-chain-payment` to ignore the HUP (hangup) signal and therefore avoid stop when you log out.
- You can add `>> mcp.log` in the command to let all the logs output to `mcp.log`.
- You can add `&` at the end of the command to let the program run in background.
- Such as:
```shell
nohup ./multi-chain-payment-0.2.1-rc1-unix >> mcp.log &   #After installation from Option 1
nohup ./build/multi-chain-payment >> mcp.log &            #After installation from Option 2
```

## Configuration

### config.toml
- **admin_wallet_on_polygon**: The wallet address used to execute contract methods on the polygon network and pay for gas
- **swan_payment_address_on_polygon**: The contract address of the swan payment gateway, used for lock and unlock user fees on polygon
- **file_coin_wallet**: The wallet address of the user's paying for storage file to the filecoin network
#### [lotus]
- **client_api_url**:  Url of lotus client web api, such as: `http://[ip]:[port]/rpc/v0`, generally the `[port]` is `1234`. See [Lotus API](https://docs.filecoin.io/reference/lotus-api/#features)
- **client_access_token**:  Access token of lotus client web api. It should have admin access right. You can get it from your lotus node machine using command `lotus auth create-token --perm admin`. See [Obtaining Tokens](https://docs.filecoin.io/build/lotus/api-tokens/#obtaining-tokens)
#### [ipfs_server]
- **download_url_prefix**: Ipfs server url prefix, such as: `http://[ip]:[port]`. Store car files for downloading by storage provider. Car file url will be `[download_url_prefix]/ipfs/[file_hash]`
- **upload_url_prefix**: Ipfs server url for uploading files, such as `http://[ip]:[port]`
#### [swan_task]
- **dir_deal:** Output directory for saving generated Car files and CSVs
- **verified_deal:** [true/false] Whether deals in this task are going to be sent as verified
- **fast_retrieval:** [true/false] Indicates that data should be available for fast retrieval
- **start_epoch_hours:** start_epoch for deals in hours from current time
- **expired_days:** expected completion days for storage provider sealing data
- **max_price:** Max price willing to pay per GiB/epoch for offline deal
- **min_price:** Min price willing to pay per GiB/epoch for offline deal
- **generate_md5:** [true/false] Whether to generate md5 for each car file, note: this is a resource consuming action

### config_polygon.toml

Currently, USDC is supported for payment. Take polygon network as an example to introduce configuration items
#### polygon_mainnet_node

- **rpc_url:** the polygon network rpc url
- **payment_contract_address:**  swan payment gateway address on polygon
- **contract_function_signature:**  swan payment gateway's lock payment event's function signature on polygon
- **dao_swan_oracle_address:**  swan dao address on polygon
- **dao_event_function_signature:**  swan dao's signature event's function signature on polygon
- **scan_step:**  the number of blocks scanned per scan
- **start_from_blockNo:**  scan data from this block number
- **start_from_blockNo:**  the time between each scan of the blockchain

## Payment Process

**Step:one:** Users upload a file they want to backup to filecoin network, then use the currencies we support to send tokens
to our contract address.<br>
**Step:two:** MCP scans the events of the above transactions<br>
**Step:three:** When the event data that get in second step meet the conditions, and then the user can perform the filecoin
network storage function<br>
**Step:four:** when the user's storage is successful, it will be scanned by the dao organization, and then dao signed to
agree to unlock the user's payment.<br>
**Step:five:** If more than half of the dao agree, the payment bridge will unlock the user's payment, deduct the user's
storage fee, and the remaining locked virtual currency Is returned to the customer's wallet<br>

## Database Table Introduction
- You can get db table ddl sql script in `[mcp-source-file-path]/script/dbschema.sql`
- Two tables should be initialized before it can be used

### system_config_param
|column                 |description       |
|param_key              |param_value       |
|-----------------------|------------------|
|SWAN_PAYMENT_CONTRACT_ADDRESS     |swan payment gateway contract address                     |
|LOCK_TIME                         |time that user's token will be locked time                |
|RECIPIENT                         |admin wallet address, used to receive tokens paid by users|
|PAY_GAS_LIMIT                     |max gas limit                                             |
|USDC_ADDRESS                      |usdc address                                              |

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



## Other Topics

- [how to pay for filecoin network storage with polygon](https://www.youtube.com/watch?v=c4Dvidz3plU)

## License

[MIT](https://github.com/filswan/multi-chain-payment/blob/main/LICENSE)
