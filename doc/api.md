```
[GET] /api/v1/public/quote
```
This endpoint retrieves a 24-hour pricing of per GB storage cost

**Response is cached for:**
_1 second_

**Parameters:**
NONE

**Response:**
```json5
{
  "price": "1000", // 10^18 based number, not use decimal
}
```

```
[GET] /api/v1/public/status/:cid
```
This endpoint retrieves a 24-hour pricing of per GB storage cost

**Response is cached for:**
10 minutes

**Parameters:**
cid

**Response:**

Available statuses:
* `Status 200`
```json5
{
  "progress": 0-100, // progress%
  "desc": "Preparing to seal...." // message to describe current status
}
```


```
[POST] /api/v1/upload
```
This endpoint upload file to fileswan payment gateway server

**HTTP hearders:**
***Content-Type***
***Content-Length***


**Request BODY raw:**
```
file's data in bytes
```

**Response:**


Available statuses:
* `Status 200`
* `Status 550 if cannot find CID in smart contract`
* `Status 551 if locked token is too low to pay`
* `Status 552 if storage transaction passed deadline`
* `Status 553 if storage transaction is completed`
```json5
{
    "cid": "bafy2bzacedswv4asge2luw6xpaetzelmmekodmknbpazlszqb6r5rxjo4fthq"
}
```
