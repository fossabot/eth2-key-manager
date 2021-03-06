package core

import "github.com/tyler-smith/go-bip39"

// GenerateNewEntropy generates a new entropy
func GenerateNewEntropy() ([]byte, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}

// Given an entropy, create the mnemonic passphrase.
func EntropyToMnemonic(entropy []byte) (string, error) {
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

// SeedFromMnemonic generates a new seed.
// The seed is the product of applying a key derivation algo (PBKDF2) on the mnemonic (as the entropy)
// and the password as salt.
// Please see https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
func SeedFromMnemonic(mnemonic string, password string) ([]byte, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return seed, nil
}

// the seed is the product of applying a key derivation algo (PBKDF2) on the mnemonic (as the entropy)
// and the password as salt.
// please see https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
func SeedFromEntropy(entropy []byte, password string) ([]byte, error) {
	mnemonic, err := EntropyToMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	return bip39.NewSeed(mnemonic, password), nil
}
