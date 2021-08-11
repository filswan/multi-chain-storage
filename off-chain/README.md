#Payment-Bridge

Scan the blockchain, find the event log of the block where the swan payment contract has been executed,  <br>
and save to the database

## Getting Started

1 &ensp;  clone code to $GOPATH/src
```console
#git clone https://github.com/filswan/payment-bridge
```

2 &ensp;  Enter payment-bridge/off-chain/config directory, <br>
&ensp;  &ensp;     change config.toml.example to config.toml , <br>
&ensp;  &ensp; input your database parameters and blockchain node address
```console
#cd $GOPATH/src/payment-bridge/off-chain/config
#mv config.toml.example config.toml
```

3 &ensp;  Enter payment-bridge directory, <br>
&ensp;  &ensp; and execute the make command 
```console
#cd $GOPATH/src/payment-bridge/
#make all
```

4 &ensp;  Enter payment-bridge/off-chain/build/bin directory, <br>  
&ensp;  &ensp;  execute the binary file of the payment-bridge project
```console
#cd $GOPATH/src/payment-bridge/off-chain/build/bin
#chmod +x payment-bridge
#./payment_bridge
```

## Versioning and Releases

Feature branches and master are designated as **unstable** which are internal-only development builds.

Periodically a build will be designated as **stable** and will be assigned a version number by tagging the repository
using Semantic Versioning in the following format: `vMajor.Minor.Patch`.