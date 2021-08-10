# payment-bridge


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