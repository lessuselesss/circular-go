/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'exports.go', provides backward-compatible package-level function exports
for the Circular Protocol Go SDK. It maintains the original API by delegating all
function calls to a default Client instance, ensuring existing code that uses
package-level functions (e.g., circular_protocol_api.CheckWallet) continues to work
without modification.

Package: github.com/circular-protocol/circular-go/circular_protocol_api
Purpose: Backward-compatible package-level exports for the Circular Protocol SDK
*/
package circular_protocol_api

// Default client instance for backward compatibility
var defaultClient *Client

// Package-level configuration variables (preserved for backward compatibility)
var __NAG_URL__ = defaultNAGURL
var __NAG_KEY__ = ""
var __lastError__ = ""

func init() {
	defaultClient = NewClient(__NAG_URL__)
}

// ==================== CONFIGURATION FUNCTIONS ====================

// SetNAGKey sets the NAG key for the default client.
func SetNAGKey(key string) {
	__NAG_KEY__ = key
	defaultClient.SetNAGKey(key)
}

// GetNAGKey returns the current NAG key.
func GetNAGKey() string {
	return __NAG_KEY__
}

// SetNAGURL sets the NAG URL and recreates the default client.
func SetNAGURL(url string) {
	__NAG_URL__ = url
	defaultClient = NewClient(url)
}

// GetNAGURL returns the current NAG URL.
func GetNAGURL() string {
	return __NAG_URL__
}

// GetVersion returns the current SDK version.
func GetVersion() string {
	return version
}

// ==================== SMART CONTRACT FUNCTIONS ====================

// TestContract tests the execution of a smart contract.
func TestContract(blockchain string, sender string, project string) map[string]interface{} {
	return defaultClient.TestContract(blockchain, sender, project)
}

// CallContract calls a function on a smart contract.
func CallContract(blockchain string, from string, project string, request string) map[string]interface{} {
	return defaultClient.CallContract(blockchain, from, project, request)
}

// ==================== WALLET FUNCTIONS ====================

// CheckWallet checks if a wallet is valid/registered on the blockchain.
func CheckWallet(blockchain string, address string) map[string]interface{} {
	return defaultClient.CheckWallet(blockchain, address)
}

// GetWallet retrieves information about a wallet.
func GetWallet(blockchain string, address string) map[string]interface{} {
	return defaultClient.GetWallet(blockchain, address)
}

// GetWalletNonce retrieves the nonce of a wallet.
func GetWalletNonce(blockchain string, address string) map[string]interface{} {
	return defaultClient.GetWalletNonce(blockchain, address)
}

// GetLatestTransaction retrieves the latest transactions of a wallet.
func GetLatestTransaction(blockchain string, address string) map[string]interface{} {
	return defaultClient.GetLatestTransaction(blockchain, address)
}

// GetWalletBalance retrieves the balance of a specific asset for the given wallet.
func GetWalletBalance(blockchain string, address string, asset string) map[string]interface{} {
	return defaultClient.GetWalletBalance(blockchain, address, asset)
}

// RegisterWallet registers a new wallet on the blockchain.
func RegisterWallet(blockchain string, publicKey string) map[string]interface{} {
	return defaultClient.RegisterWallet(blockchain, publicKey)
}

// ==================== DOMAIN MANAGEMENT ====================

// GetDomain resolves a domain name to a wallet address.
func GetDomain(blockchain string, name string) map[string]interface{} {
	return defaultClient.GetDomain(blockchain, name)
}

// ==================== ASSET MANAGEMENT ====================

// GetAssetList retrieves the list of assets on a blockchain.
func GetAssetList(blockchain string) map[string]interface{} {
	return defaultClient.GetAssetList(blockchain)
}

// GetAsset retrieves asset description.
func GetAsset(blockchain string, asset string) map[string]interface{} {
	return defaultClient.GetAsset(blockchain, asset)
}

// GetAssetSupply retrieves the supply of an asset.
func GetAssetSupply(blockchain string, asset string) map[string]interface{} {
	return defaultClient.GetAssetSupply(blockchain, asset)
}

// GetVoucher retrieves voucher information by code.
func GetVoucher(blockchain string, code string) map[string]interface{} {
	return defaultClient.GetVoucher(blockchain, code)
}

// ==================== BLOCKCHAIN MANAGEMENT ====================

// GetBlockRange retrieves a range of blocks.
func GetBlockRange(blockchain string, start int, end int) map[string]interface{} {
	return defaultClient.GetBlockRange(blockchain, start, end)
}

// GetBlock retrieves a specific block by number.
func GetBlock(blockchain string, number int) map[string]interface{} {
	return defaultClient.GetBlock(blockchain, number)
}

// GetBlockCount retrieves the current block count of a blockchain.
func GetBlockCount(blockchain string) map[string]interface{} {
	return defaultClient.GetBlockCount(blockchain)
}

// GetAnalytics retrieves analytics data for a blockchain.
func GetAnalytics(blockchain string) map[string]interface{} {
	return defaultClient.GetAnalytics(blockchain)
}

// GetBlockchains retrieves the list of available blockchains.
func GetBlockchains() map[string]interface{} {
	return defaultClient.GetBlockchains()
}

// ==================== TRANSACTION MANAGEMENT ====================

// GetPendingTransaction retrieves a pending transaction by ID.
func GetPendingTransaction(blockchain string, TxID string) map[string]interface{} {
	return defaultClient.GetPendingTransaction(blockchain, TxID)
}

// GetTransactionByID retrieves a transaction by ID within a block range.
func GetTransactionByID(blockchain string, TxID string, start int, end int) map[string]interface{} {
	return defaultClient.GetTransactionByID(blockchain, TxID, start, end)
}

// GetTransactionByNode retrieves transactions by node ID within a block range.
func GetTransactionByNode(blockchain string, node string, start int, end int) map[string]interface{} {
	return defaultClient.GetTransactionByNode(blockchain, node, start, end)
}

// GetTransactionByAddress retrieves all transactions involving a specific address.
func GetTransactionByAddress(blockchain string, address string, start int, end int) map[string]interface{} {
	return defaultClient.GetTransactionByAddress(blockchain, address, start, end)
}

// GetTransactionByDate retrieves transactions involving an address within a date range.
func GetTransactionByDate(blockchain string, address string, startDate string, endDate string) map[string]interface{} {
	return defaultClient.GetTransactionByDate(blockchain, address, startDate, endDate)
}

// SendTransaction sends a transaction to the blockchain.
func SendTransaction(id string, sender string, to string, timestamp string, transactionType string, payload string, nonce string, signature string, blockchain string) map[string]interface{} {
	return defaultClient.SendTransaction(id, sender, to, timestamp, transactionType, payload, nonce, signature, blockchain)
}

// GetTransactionOutcome polls for transaction finality with a timeout.
func GetTransactionOutcome(blockchain string, TxID string, timeoutSec int) string {
	return defaultClient.GetTransactionOutcome(blockchain, TxID, timeoutSec)
}
