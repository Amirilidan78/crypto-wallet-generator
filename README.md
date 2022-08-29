### Setup 

https://developer.trustwallet.com/wallet-core/integration-guide/server-side

### Functions 

###### IsAddressValid
check wallet address is valid or not 
```
IsAddressValid(address string, coin Coin) bool
```

###### GenerateWallet
generate hd wallet for provided coin 
```
GenerateWallet(coin Coin) (string, string, string)
```

###### GetWallet
get address from `passphrase` and `seed`
```
GetWallet(passphrase string, seed string, coin Coin) string
```

### Coins 
check `coins.go` for supported blockchains


