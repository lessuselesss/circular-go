/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'wallet/wallet.go', defines the WalletService struct and implements
all wallet-specific API functions for the Circular Protocol Go SDK. This includes
wallet registration, balance checking, nonce retrieval, and transaction history.
It is an internal service package used by the main Client Facade to interact with
the Circular Layer 1 Blockchain Protocol.

Package: github.com/circular-protocol/circular-go/circular_protocol_api/wallet
Purpose: Wallet operations for the Circular Protocol SDK
*/
package wallet

import (
	"encoding/hex"
	"encoding/json"

	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// Service provides wallet-related operations for the Circular Protocol.
type Service struct {
	nagURL  string
	version string
}

// NewService creates a new WalletService instance.
func NewService(nagURL string, version string) *Service {
	return &Service{
		nagURL:  nagURL,
		version: version,
	}
}

// CheckWallet checks if a wallet is valid/registered on the blockchain.
func (s *Service) CheckWallet(blockchain string, address string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.CHECK_WALLET, s.nagURL)
}

// GetWallet retrieves information about a wallet.
func (s *Service) GetWallet(blockchain string, address string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_WALLET, s.nagURL)
}

// GetWalletNonce retrieves the nonce of a wallet.
func (s *Service) GetWalletNonce(blockchain string, address string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_WALLET_NONCE, s.nagURL)
}

// GetLatestTransaction retrieves the latest transactions of a wallet.
func (s *Service) GetLatestTransaction(blockchain string, address string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_LATEST_TRANSACTIONS, s.nagURL)
}

// GetWalletBalance retrieves the balance of a specific asset for the given wallet.
func (s *Service) GetWalletBalance(blockchain string, address string, asset string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"Asset":      utils.HexFix(asset),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_WALLET_BALANCE, s.nagURL)
}

// RegisterWallet registers a new wallet on the blockchain.
// This is a helper that creates and sends the registration transaction.
func (s *Service) RegisterWallet(blockchain string, publicKey string, sendTransaction func(id, sender, to, timestamp, transactionType, payload, nonce, signature, blockchain string) map[string]interface{}) map[string]interface{} {
	blockchain = utils.HexFix(blockchain)
	publicKey = utils.HexFix(publicKey)

	sender := utils.Sha256(publicKey)
	to := sender
	nonce := "0"
	txType := "C_TYPE_REGISTERWALLET"
	payloadObj := map[string]interface{}{
		"Action":    "CP_REGISTERWALLET",
		"PublicKey": publicKey,
	}

	jsonData, err := json.Marshal(payloadObj)
	if err != nil {
		return map[string]interface{}{
			"Error": "Wrong payload format",
		}
	}

	payload := hex.EncodeToString(jsonData)
	timestamp := utils.GetFormattedTimestamp()
	dataToHash := blockchain + sender + to + payload + nonce + timestamp
	id := utils.Sha256(dataToHash)
	signature := ""

	return sendTransaction(id, sender, to, timestamp, txType, payload, nonce, signature, blockchain)
}
