package in_memory

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	eth2keymanager "github.com/bloxapp/eth2-key-manager"
	"github.com/bloxapp/eth2-key-manager/core"
	"github.com/bloxapp/eth2-key-manager/stores"
)

func _byteArray(input string) []byte {
	res, _ := hex.DecodeString(input)
	return res
}

func getPopulatedWalletStorage() (core.Storage, []core.ValidatorAccount, error) {
	core.InitBLS()
	store := getStorage()

	// seed
	seed := _byteArray("0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1fff")

	options := &eth2keymanager.KeyVaultOptions{}
	options.SetStorage(store)
	vault, err := eth2keymanager.NewKeyVault(options)
	if err != nil {
		return nil, nil, err
	}

	wallet, err := vault.Wallet()
	if err != nil {
		return nil, nil, err
	}

	a1, err := wallet.CreateValidatorAccount(seed, nil)
	if err != nil {
		return nil, nil, err
	}
	a2, err := wallet.CreateValidatorAccount(seed, nil)
	if err != nil {
		return nil, nil, err
	}
	a3, err := wallet.CreateValidatorAccount(seed, nil)
	if err != nil {
		return nil, nil, err
	}
	a4, err := wallet.CreateValidatorAccount(seed, nil)
	if err != nil {
		return nil, nil, err
	}

	return store, []core.ValidatorAccount{a1, a2, a3, a4}, nil
}

func TestOpeningAccount(t *testing.T) {
	storage, accounts, err := getPopulatedWalletStorage()
	require.NoError(t, err)
	stores.TestingOpeningAccount(storage, accounts[0], t)
}

func TestAddingAccountsToWallet(t *testing.T) {
	storage, accounts, err := getPopulatedWalletStorage()
	require.NoError(t, err)
	stores.TestingSavingAccounts(storage, accounts, t)
}

func TestFetchingNonExistingAccount(t *testing.T) {
	storage, _, err := getPopulatedWalletStorage()
	require.NoError(t, err)
	stores.TestingFetchingNonExistingAccount(storage, t)
}

func TestListingAccounts(t *testing.T) {
	storage, _, err := getPopulatedWalletStorage()
	require.NoError(t, err)
	stores.TestingListingAccounts(storage, t)
}
