## MCS Contracts

### FilinkConsumer.sol

Filink Consumer is a Chainlink Client Contract to make a GET request from a
public API. This API returns the Filecoin deal information given the `deal_id`
and `network`. MCS uses this API to get the storage price for the deal from the
Filecoin network
(Note: `PolygonFilink.sol` is the same logic, just uses a different Oracle and jobID for mainnet deployment).

- **`requestDealInfo(deal, network)`**  
  This function performs a GET request to get the deal information on the network
  (`filecoin_mainnet` or `filecoin_calibration`)

- **`getPrice(deal, network)`**
  After calling `requestDealInfo`, The chainlink Oracle will call our `fulfill`
  function, storing price in this contract's mapping. Users can retreive the price
  by calling this function.

### FilSwanOracle.sol

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

### PriceFeed.sol

This contract consults the price for wFIL-USDC by using a SushiSwap liquidity pool.

- **`consult(tokenInput, amount)`**
  1. find USDC price for 0.001 wFIL (bc LP is small)
  2. multiply by 10^12 because USDC has 6 decimals instead of 18
  3. now we have 0.001 wFIL in USDC, multiply by 10^3 to get USDC for 1 wFIL
  4. multiply by `amount` to get `amount` wFIL in USDC
  5. finally, divide by 10^18 to get readable USDC price

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

### MCSNFT.sol

MCSNFT is an ERC-721 contract, that allows any user to mint a new NFT for this
collection.

- **`mintData(minter, uri)`** mints a new token for `minter` and attaches the `uri` to the `tokenId`
- **`totalSupply()`** returns the total number of NFTs

### SwanFaucet.sol

This contract is deployed on Mumbai and Binance testnet only. This faucet allows
users to obtain our test USDC tokens and MATIC (on Mumbai). Interact with the UI
[here](https://calibration-faucet.filswan.com/#/dashboard) (Note: the faucet can
only be used once a day and the contract limits addresses every 24 hours).

- **`sendMultiTokens(tokenAddresses[], tokenAmounts[], address)`**  
  This function is only callable by the contract admins. It transfers funds from
  the contract address to `address`.
