# payment-bridge


script to generate golang bind

```
abigen --abi ./contracts/abi/SwanPayment.json --pkg fileswan_payment --type PaymentGateway --out goBind/PaymentGateway.go
```


// todo: implement
add event emitter
returns err code of unlock payment call