[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on Slack](https://img.shields.io/badge/slack-filswan.slack.com-green.svg)](https://filswan.slack.com)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)

# Payment-Bridge

Payment bridge is designed for make payment from multi chain for filecoin storage,and backup user's file to filecoin
network. Now supports payment with tokens such as USDC on polygon

# System Design

![MCP-MCP Desgin](https://user-images.githubusercontent.com/8363795/143811916-f051ccce-f9b2-49eb-99ab-8da1a0d9f2f2.png)

# Modules

* [Token Swap](#Token Swap)
* Data DAO : [Flink](https://github.com/filswan/flink)
* IPFS/Filecoin Storage

## Token Swap

Token Swap module is in charge of swap the user token to wrapped token, it can be USDC or other tokens.

### Scan Module

Scan the blockchain, find the event log of the block where the swan payment contract has been executed,and save to the
database

### Prerequisite

- Golang1.16 (minimum version)
- Mysql5.5
- Lotus lite node (Please make sure that lotus lite node and payment bridge project are running on the same server)

### Install lotus node

Lotus node is Used for making car files and sending offline deals <br>
You can use lotus full node, but the full node is too heavy, you can use lotus lite node instead <br>
If you deploy the payment-bridge project, you need to deploy the payment-bridge and lotus full nodes on the same server <br>
If you don’t want to deploy payment-bridge and lotus full node on the same server,you can also choose to deploy a lotus lite node<br>
Please make sure that you already have a lotus full node, because lotus lite node needs to communicate with a full node  <br>
#### Option one: **install a lotus full node**: See [lotus full node official installation document](https://lotus.filecoin.io/docs/set-up/install/)
#### Option two: **install a lotus lite node**: See [lotus lite node official installation document](https://lotus.filecoin.io/docs/set-up/lotus-lite/#amd-and-intel-based-computers)


## Getting Started with PaymentBridge

### 1 Clone code to $GOPATH/src

```console
git clone https://github.com/filswan/payment-bridge
```

### 2 Pull-in the submodules:

```console
cd $GOPATH/src/payment-bridge/
git submodule update --init --recursive
make ffi
```

### 3 Build project

Enter payment-bridge directory,and execute the make command  <br>
You can get a runnable binary file named payment_bridge and config file in $GOPATH/src/payment-bridge/config/ <br>
We support multi-chain payment,for example<br>
If you want to pay on polygon network,you also need to edit $GOPATH/src/payment-bridge/config/polygon/config_polygon.toml<br>
If you want to pay on goerli network,you also need to edit$GOPATH/src/payment-bridge/config/goerli/config_goerli.toml<br>

```console
cd $GOPATH/src/payment-bridge/
GO111MODULE=on make
```

### 4 Run the project

Enter payment-bridge/build/ directory, and execute the binary file of the payment-bridge project <br>
Before actually running the payment bridege project, you need to read the section about config introduction below <br>
and then fill in the configuration items as required, and then run this project<br>

```console
cd $GOPATH/src/payment-bridge/build/`
chmod +x payment-bridge
./payment_bridge
```

##### Note:
you need to edit the config file and input your config params, the configuration items will be introduced below<br>
Take payment on the polygon network as an example, before running the ./payment-bridge command<br>
You need to edit two config file: 
- $GOPATH/src/payment-bridge/build/config/config.toml
- $GOPATH/src/payment-bridge/build/config/polygon/config_polygon.toml
## Configuration

### config.toml

```
admin_wallet_on_polygon=""
swan_payment_address_on_polygon=""
file_coin_wallet=""

[database]
db_host="localhost"
db_port="3306"
db_schema_name="payment_bridge"
db_username="root"
db_pwd="123456"
db_args="charset=utf8mb4&parseTime=True&loc=Local"

[swan_api]
api_url = "http://xxxxxxxx:5002"
api_key = ""
access_token = ""

[lotus]
api_url="http://127.0.0.1:1234/rpc/v0"   # Url of lotus web api
access_token=".."   # Access token of lotus web api

[ipfs_server]
download_url_prefix = "http://192.168.xx.xx:5050"
upload_url = "http://192.168.xx.xx:5001"

[swan_task]
min_price = 0.00000001
max_price = 0.00005
expire_days = 4
verified_deal = false
fast_retrieval = true
start_epoch_hours = 96
get_should_send_task_url_suffix = "paymentgateway/deals?offset=0&limit=10&source_id=4&task_status=Assigned&is_public=1"


[schedule_rule]
unlock_payment_rule = "0 */5 * * * ?"  #every 3 minute
send_deal_rule = "0 */3 * * * ?"  #every 3 minute
create_task_rule = "0 */2 * * * ?"
scan_deal_status_rule = "0 */4 * * * ?"
```

#### admin_wallet_on_polygon

The wallet address used to execute contract methods on the polygon network and pay for gas

#### swan_payment_address_on_polygon

The contract address of the swan payment gateway, used for lock and unlock user fees on polygon

#### file_coin_wallet

The wallet address of the user's paying for storage file to the filecoin network

#### [swan_api]

swan_api section defines the token used for connecting with Swan platform.

- **api_key & access_token:** Acquire from [Filswan](https://www.filswan.com) -> "My Profile"->"Developer Settings". You
  can also check the [Guide](https://nebulaai.medium.com/how-to-use-api-key-in-swan-a2ebdb005aa4)
- **api_url:** Default: "https://api.filswan.com"

#### [lotus]

- **api_url:** Url of lotus client web api, such as: **http://[ip]:[port]/rpc/v0**, generally the [port] is **1234**.
  See [Lotus API](https://docs.filecoin.io/reference/lotus-api/)
- **access_token:** Access token of lotus market web api. When market and miner are not separate, it is also the access
  token of miner access token. See [Obtaining Tokens](https://docs.filecoin.io/build/lotus/api-tokens/#obtaining-tokens)

#### [ipfs_server]

ipfs-server is used to upload and download generated Car files. You can upload generated Car files via `upstream_url`
and storage provider will download Car files from this ipfs-server using `download_stream_url`. The downloadable URL in
the CSV file is built with the following format: host+port+ipfs+hash, e.g. http://host:
port/ipfs/QmPrQPfGCAHwYXDZDdmLXieoxZP5JtwQuZMUEGuspKFZKQ

#### [swan_task]

- **dir_deal:** Output directory for saving generated Car files and CSVs
- **verified_deal:** [true/false] Whether deals in this task are going to be sent as verified
- **fast_retrieval:** [true/false] Indicates that data should be available for fast retrieval
- **start_epoch_hours:** start_epoch for deals in hours from current time
- **expired_days:** expected completion days for storage provider sealing data
- **max_price:** Max price willing to pay per GiB/epoch for offline deal
- **min_price:** Min price willing to pay per GiB/epoch for offline deal
- **generate_md5:** [true/false] Whether to generate md5 for each car file, note: this is a resource consuming action
- **relative_epoch_from_main_network:**  The difference between the block height of filecoin's mainnet and testnet. If
  the filecoin testnet is linked, this configuration item is set to -858481, and if the link is the mainnet, it is set
  to 0

### config_polygon.toml

Currently, USDC is supported for payment. Take polygon network as an example to introduce configuration items

```console
[polygon_mainnet_node]
#rpcUrl="https://rpc-mumbai.maticvigil.com"
rpc_url = "https://polygon-mumbai.g.alchemy.com/v2/xxx"
payment_contract_address = ""
contract_function_signature = ""
dao_swan_oracle_address = ""
dao_event_function_signature = ""
scan_step = 1000
gas_limit = 8000000
start_from_blockNo = 17590780
cycle_time_interval = 10 #unit:second
```

#### polygon_mainnet_node

- **rpc_url:** the polygon network rpc url
- **payment_contract_address:**  swan payment gateway address on polygon
- **contract_function_signature:**  swan payment gateway's lock payment event's function signature on polygon
- **dao_swan_oracle_address:**  swan dao address on polygon
- **dao_event_function_signature:**  swan dao's signature event's function signature on polygon
- **scan_step:**  the number of blocks scanned per scan
- **start_from_blockNo:**  scan data from this block number
- **start_from_blockNo:**  the time between each scan of the blockchain

## The Payment Process

- First, users upload a file they want to backup to filecoin network, then use the currencies we support to send tokens
  to our contract address. <br>
- Second, payment-biridge scans the events of the above transactions
- Third, when the event data that get in second step meet the conditions, and then the user can perform the filecoin
  network storage function
- Fourth, when the user's storage is successful, it will be scanned by the dao organization, and then dao signed to
  agree to unlock the user's payment.
- Fifth, If more than half of the dao agree, the payment bridge will unlock the user's payment, deduct the user's
  storage fee, and the remaining locked virtual currency Is returned to the customer's wallet

## Database table description
You can get db table ddl sql script in $GOPATH/src/payment-bridge/script/dbschema.sql <br>
There are two tables you need to initialize the data before you can use it<br>

###system_config_param
|column                 |description       |
|param_key              |param_value       |
|-----------------------|------------------|
|SWAN_PAYMENT_CONTRACT_ADDRESS     |swan payment gateway contract address                     |
|LOCK_TIME                         |time that user's token will be locked time                |
|RECIPIENT                         |admin wallet address, used to receive tokens paid by users|
|PAY_GAS_LIMIT                     |max gas limit                                             |
|USDC_ADDRESS                      |usdc address                                              |

###dao_info
###system_config_param
|column                 |description       |
|-----------------------|------------------|
|ID                     |primary key of table   |
|DAO_NAME               |dao's name,required    |
|DAO_ADDRESS            |dao's address,required |
|ORDER_INDEX            |dao display order      |
|DESCRIPTION            |description            |
|CREATE_AT              |create time            |



## Other Topics

- [how to pay for filecoin network storage with polygon](https://www.youtube.com/watch?v=c4Dvidz3plU)

## Versioning and Releases

Feature branches and master are designated as **unstable** which are internal-only development builds.

Periodically a build will be designated as **stable** and will be assigned a version number by tagging the repository
using Semantic Versioning in the following format: `vMajor.Minor.Patch`.

