# MCS Contracts

1. `SwanPayment` is used to lock an estimated amount of USDC for filecoin storage
2. `FilswanOracle` verifies the deal is on chain, signing transactions until the treshold of 3 is met
3. `FilinkConsumer` is used to get the storage price of a deal in FIL
4. `PriceFeed` converts this amount to USDC
5. `SwanPayment` will unlock any over-estimated amount back to the user.

## Functions

### SwanNFT.sol

MCSNFT is an ERC-1155 contract, that allows any user to mint a new NFT for this
collection.

- **`mintUnique(minter, uri)`** mints a new unique token for `minter` and attaches the `uri` to the `tokenId`
- **`totalSupply()`** returns the total number of NFTs

### FilinkConsumer.sol

Filink Consumer is a Chainlink Client Contract to make a GET request from a
public API. This API returns the Filecoin deal information given the `deal_id`
and `network`. MCS uses this API to get the storage price for the deal from the
Filecoin network

- **`requestDealInfo(deal, network)`**  
  This function performs a GET request to get the deal information on the network
  (`filecoin_mainnet` or `filecoin_calibration`)

- **`getPrice(deal, network)`**
  After calling `requestDealInfo`, The chainlink Oracle will call our `fulfill`
  function, storing price in this contract's mapping. Users can retreive the price
  by calling this function. Returns the storage_price in wFIL (in wei)

### PriceFeed.sol

This contract uses Chainlink USD/FIL price data feed to get the lastest price

- **`getLatestPrice()`**
  This function gets the lastest price of 1 FIL in USD, (8 demicals)

- **`consult(tokenInput, amount)`** gets price of `amount` FIL in USD

  1. `getLatestPrice() * amount / 10^8` to get price in wei
  2. `price in wei / 10^(18 - tokenInput decimals)` to get price in `tokenInput` decimals
  3. Return `amount` FIL in USD (`tokenInput` decimal)

### SwanPayment.sol

This contract handles the payment and unlock of funds for files stored on the
Filecoin network. Users pay in USDC and this contract consults `PriceFeed.sol`
to get the price in wFIL. With enough signatures in `FilSwanOracle.sol`, this
contract's `unlock` function will be called, to return excess payment to the user.

- **`setOracle(oracle)`** sets `FilSwanOracle` address
- **`setChainlinkOracle(chainlinkOracle)`** sets `FilinkConsumer` address
- **`setPriceFeed(priceFeed)`** sets `PriceFeed` address
- **`getLockedPaymentInfo(cId)`** get the payment information for the CID
- **`lockTokenPayment(param)`**
  - sets payment information for this CID
  - transfer `param.amount` of ERC20 tokens to this contract
- **`unlockTokenPayment(param)`** refund the locked amount to user after the deadline has passed.
- **`unlockCarPayment(dealId, network, recipient)`**
  - get `serviceCost` from `FilinkConsumer` (in FIL)
  - get `cidList` for this `dealId` and `network` from `FilSwanOracle`
  - consult `PriceFeed` to get `serviceCost` in USDC (or another ERC-20 token)
  - get the total size for files in `cidList`, `unitPrice = serviceCost(USDC) / total size`
  - cost of each file is `size * unitPrice`, subtract this cost from the file's `lockedFee`
  - transfer the `serviceCost in USDC` to `recipient`
- **`refund(cidList[])`** return any excess `lockedFee` to its owner

### FilswanOracle.sol

This is the DAO contract to sign transactions. Once a deal has enough signatures
(depending on `threshold`), the locked funds can be unlocked to the owner.

- **`updateThreshold(threshold)`** change the threshold for signatures
- **`setFilinkOracle(filinkAddress)`** change `FilinkConsumer` contract address
- **`setDAOUsers(daoUsers[])`** change the list of DAO users

- **`signCarTransaction(cidList[], dealId, network, recipient)`**  
  signs a [CAR file](https://car.ipfs.io/). Once the number of signatures reaches
  the `threshold` amount, this contract will call the `requestDealInfo` from
  `FilinkConsumer`

- **`isCarPaymentAvailable(dealId, network, recipient)`**
  returns `number of signatures >= threshold`
- **`getCarPaymentVotes(dealId, network, recipient)`**
  returns number of signatures
- **`getThreshold()`** returns `threshold`
- **`getCidList(dealId, network)`** returns `cidList` for a deal
- **`getSignatureList(dealId, network)`** returns the list of signatures for a deal

- **`preSign(dealId, network, recipient, batchCount)`**

- **`sign(dealId, network, cidList[], batchNo)`**

- **`signHash(dealId, network, recipient, voteKey)`**
- **`getHashKey(dealId, network, recipient, cidList[])`**
