package main

import (
	"encoding/hex"
	"fmt"

	eth2keymanager "github.com/bloxapp/eth2-key-manager"
	"github.com/bloxapp/eth2-key-manager/core"
	"github.com/bloxapp/eth2-key-manager/stores/in_memory"
)

func main() {
	entropy, _ := core.GenerateNewEntropy()

	// print out mnemonic
	mnemonic, _ := core.EntropyToMnemonic(entropy)
	fmt.Printf("Generated mnemonic: %s\n", mnemonic)

	// generate seed
	seed, _ := core.SeedFromEntropy(entropy, "")

	// create storage
	store := in_memory.NewInMemStore(core.PyrmontNetwork)

	// create options
	options := &eth2keymanager.KeyVaultOptions{}
	options.SetStorage(store)

	// instantiate KeyVaul
	vault, _ := eth2keymanager.NewKeyVault(options)

	// create account
	wallet, _ := vault.Wallet()
	account, _ := wallet.CreateValidatorAccount(seed, nil)

	fmt.Printf("created validator account with pub key: %s\n", hex.EncodeToString(account.ValidatorPublicKey()))

}
