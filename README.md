#Payment-Bridge

Scan the blockchain, find the event log of the block where the swan payment contract has been executed,  <br>
and save to the database

## Prerequisite
- golang1.13 (minimum version)

## Getting Started

1 &ensp;  clone code to $GOPATH/src
```console
git clone https://github.com/filswan/payment-bridge
```

2 &ensp;  Enter payment-bridge/off-chain/config directory, <br>
&ensp;  &ensp;     change config.toml.example to config.toml , <br>
&ensp;  &ensp; input your database parameters and blockchain node address
```console
cd $GOPATH/src/payment-bridge/off-chain/config
mv config.toml.example config.toml
```

3 &ensp;  Enter payment-bridge directory, <br>
&ensp;  &ensp; and execute the make command
```console
cd $GOPATH/src/payment-bridge/
go mod download
go mod tidy
GO111MODULE=on make
```

4 &ensp;  Enter payment-bridge/off-chain/build/bin directory, <br>  
&ensp;  &ensp;  execute the binary file of the payment-bridge project
```console
cd $GOPATH/src/payment-bridge/off-chain/build/bin
chmod +x payment-bridge
./payment_bridge
```

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