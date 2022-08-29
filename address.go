package crypto_wallet_generator

// #include <TrustWalletCore/TWAnyAddress.h>
import "C"

// IsAddressValid returns address valid or not
func IsAddressValid(address string, coin Coin) bool {
	result := C.TWAnyAddressIsValid(TWStringCreateWithGoString(address), coin)
	return bool(result)
}

// GenerateWallet returns address, passphrase ,seed
func GenerateWallet(coin Coin) (string, string, string) {

	passphrase := generateRandomString(16)

	seed := generateSeed(passphrase)

	return GetWallet(passphrase, seed, coin), passphrase, seed
}

// GetWallet returns address
func GetWallet(passphrase string, seed string, coin Coin) string {
	cSeed := TWStringCreateWithGoString(seed)
	cPassphrase := TWStringCreateWithGoString(passphrase)
	defer C.TWStringDelete(cSeed)
	defer C.TWStringDelete(cPassphrase)
	wallet := C.TWHDWalletCreateWithMnemonic(cSeed, cPassphrase)
	defer C.TWHDWalletDelete(wallet)
	address := C.TWHDWalletGetAddressForCoin(wallet, coin)
	return TWStringGoString(address)
}

// generateSeed returns seed
func generateSeed(passphrase string) string {
	cPassphrase := TWStringCreateWithGoString(passphrase)
	defer C.TWStringDelete(cPassphrase)
	wallet := C.TWHDWalletCreate(128, cPassphrase)
	defer C.TWHDWalletDelete(wallet)
	return TWStringGoString(C.TWHDWalletMnemonic(wallet))
}
