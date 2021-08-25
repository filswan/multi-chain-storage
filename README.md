# Payment-Bridge
Payment bridge is desgined for make payment from multichain for filecoin storage.
## Scan Module 
Scan the blockchain, find the event log of the block where the swan payment contract has been executed,and save to the database

## Prerequisite
- golang1.13 (minimum version)

## Getting Started

1 &ensp;  clone code to $GOPATH/src
```console
git clone https://github.com/filswan/payment-bridge
```

2 &ensp;  Enter payment-bridge/scan/config directory <br>
&ensp;  &ensp;     change config.toml.example to config.toml  <br>
&ensp;  &ensp; input your database parameters and blockchain node address
```console
cd $GOPATH/src/payment-bridge/scan/config
mv config.toml.example config.toml
```

3 &ensp; Enter payment-bridge directory <br>
&ensp; &ensp;  and execute the make command
```console
cd $GOPATH/src/payment-bridge/
GO111MODULE=on make
```

## Run the project   
Enter payment-bridge/scan/build/bin directory <br> 
execute the binary file of the payment-bridge project
```console
cd $GOPATH/src/payment-bridge/scan/build/bin
chmod +x payment-bridge
./payment_bridge
```

## Call the api to check your scan data from chain <br>
The default port number of the project is 8888, you can modify it to the port you want in config.toml
Access the application at the address [http://localhost:8888/api/v1/events/logs/:contractAddress/:payloadCid](http://localhost:8888/api/v1/events/logs/:contractAddress/:payloadCid) <br>
Replace: contractAddress and: payloadCid with your own values.

the return value like this below:
```json
{
  "status": "success",
  "code": "200",
  "data": {
    "goerli": [
      {
        "id": 17,
        "tx_hash": "0xee80d76a0d273ac8620d837b2acbc15a322eeded28b2fa61e1479d85cf38755a",
        "event_name": "",
        "indexed_data": "",
        "contract_name": "SwanPayment",
        "contract_address": "0x5210ED929B5BEdBFFBA2F6b9A0b1B608eEAb3aa0",
        "locked_fee": "50000000000000000",
        "deadline": "1629149936",
        "payload_cid": "bafk2bzaceajwszlar4obpnmifoydefxlhfd7npbnmq3onkfzkincyy4fdj5xk",
        "block_no": 17685850,
        "miner_address": "0xE53AEd6DEA9e44116D4551a93eEeE28bC8684916"
      }
    ],
    "polygon": [
      {
        "id": 15,
        "tx_hash": "0x10bee0dd4907015fefa813263b1302a2cb0e8d2e8feb18b0551a12d26f24ab61",
        "event_name": "",
        "indexed_data": "",
        "contract_name": "SwanPayment",
        "contract_address": "0x5210ED929B5BEdBFFBA2F6b9A0b1B608eEAb3aa0",
        "locked_fee": "50000000000000000",
        "deadline": "1629142492",
        "payload_cid": "bafk2bzaceajwszlar4obpnmifoydefxlhfd7npbnmq3onkfzkincyy4fdj5xk",
        "block_no": 17682240,
        "miner_address": "0xE53AEd6DEA9e44116D4551a93eEeE28bC8684916"
      }
    ]
  }
}
```

## Database table description

|table                 |description       |
|----------------------|------------------|
|block_scan_record     |record the block number that has been scanned to the blockchain|
|event_goerli          |record eligible data on goerli            |
|event_polygon         |record eligible data on polygon   |


## Other Topics
- [how to pay for filecoin network storage with polygon](https://www.youtube.com/watch?v=c4Dvidz3plU)




## Versioning and Releases

Feature branches and master are designated as **unstable** which are internal-only development builds.

Periodically a build will be designated as **stable** and will be assigned a version number by tagging the repository
using Semantic Versioning in the following format: `vMajor.Minor.Patch`.



#Need lao liu to update
script to generate golang bind

```
abigen --abi ./contracts/abi/SwanPayment.json --pkg fileswan_payment --type PaymentGateway --out goBind/PaymentGateway.go
```


// todo: implement
add event emitter
returns err code of unlock payment call

goerli payment contract address:
0xad8cE271beE7b917F2a1870C8b64EDfF6aAF3342

https://goerli.etherscan.io/address/0xad8cE271beE7b917F2a1870C8b64EDfF6aAF3342

oracle contract address:
0x940695A2B36084aD0552C84af72c859e8C6b0b38

goerli contract owner: 0xE41c53Eb9fce0AC9D204d4F361e28a8f28559D54  
goerli oracle contract owner: 0xE41c53Eb9fce0AC9D204d4F361e28a8f28559D54
