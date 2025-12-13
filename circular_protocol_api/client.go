/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'client.go', defines the public Client struct, which acts as the main
non-breaking Facade for the Circular Protocol Go SDK. It manages configuration
and uses struct composition to delegate all public method calls to internal,
modular service packages (Wallet, Contract, Blockchain, Crypto), ensuring backward
compatibility for external users interacting with the Circular Layer 1
Blockchain Protocol.

Package: github.com/circular-protocol/circular-go/circular_protocol_api
Purpose: Main Facade client for the Circular Protocol SDK
*/
package circular_protocol_api

import (
	"github.com/circular-protocol/circular-go/circular_protocol_api/blockchain"
	"github.com/circular-protocol/circular-go/circular_protocol_api/contract"
	"github.com/circular-protocol/circular-go/circular_protocol_api/crypto"
	"github.com/circular-protocol/circular-go/circular_protocol_api/wallet"
)

const version = "1.0.8"
const defaultNAGURL = "https://nag.circularlabs.io/NAG.php?cep="

// Client is the main Facade struct for interacting with the Circular Protocol API.
// It composes internal service structs to handle specific functional domains.
type Client struct {
	nagURL string
	nagKey string

	// Service Delegates (Composed Services)
	// These fields hold the implementation logic and can be accessed directly
	// or via the Client's delegating public methods.
	Wallet     *wallet.Service
	Contract   *contract.Service
	Blockchain *blockchain.Service
	Crypto     *crypto.Service
}

// NewClient is the public constructor for the Circular Protocol API Client.
func NewClient(nagURL string) *Client {
	c := &Client{
		nagURL: nagURL,
	}

	// Initialize all service delegates
	c.Wallet = wallet.NewService(nagURL, version)
	c.Contract = contract.NewService(nagURL, version)
	c.Blockchain = blockchain.NewService(nagURL, version)
	c.Crypto = crypto.NewService(version)

	return c
}

// GetVersion returns the current SDK version.
func (c *Client) GetVersion() string {
	return version
}

// GetNAGURL returns the current NAG URL.
func (c *Client) GetNAGURL() string {
	return c.nagURL
}

// SetNAGKey sets the NAG key for the client.
func (c *Client) SetNAGKey(key string) {
	c.nagKey = key
}

// GetNAGKey returns the current NAG key.
func (c *Client) GetNAGKey() string {
	return c.nagKey
}

// ==================== WALLET FUNCTIONS (DELEGATED) ====================

// CheckWallet checks if a wallet is valid/registered on the blockchain.
func (c *Client) CheckWallet(blockchain string, address string) map[string]interface{} {
	return c.Wallet.CheckWallet(blockchain, address)
}

// GetWallet retrieves information about a wallet.
func (c *Client) GetWallet(blockchain string, address string) map[string]interface{} {
	return c.Wallet.GetWallet(blockchain, address)
}

// GetWalletNonce retrieves the nonce of a wallet.
func (c *Client) GetWalletNonce(blockchain string, address string) map[string]interface{} {
	return c.Wallet.GetWalletNonce(blockchain, address)
}

// GetLatestTransaction retrieves the latest transactions of a wallet.
func (c *Client) GetLatestTransaction(blockchain string, address string) map[string]interface{} {
	return c.Wallet.GetLatestTransaction(blockchain, address)
}

// GetWalletBalance retrieves the balance of a specific asset for the given wallet.
func (c *Client) GetWalletBalance(blockchain string, address string, asset string) map[string]interface{} {
	return c.Wallet.GetWalletBalance(blockchain, address, asset)
}

// RegisterWallet registers a new wallet on the blockchain.
func (c *Client) RegisterWallet(blockchain string, publicKey string) map[string]interface{} {
	return c.Wallet.RegisterWallet(blockchain, publicKey, c.Blockchain.SendTransaction)
}

// ==================== CONTRACT FUNCTIONS (DELEGATED) ====================

// TestContract tests the execution of a smart contract.
func (c *Client) TestContract(blockchain string, sender string, project string) map[string]interface{} {
	return c.Contract.TestContract(blockchain, sender, project)
}

// CallContract calls a function on a smart contract.
func (c *Client) CallContract(blockchain string, from string, project string, request string) map[string]interface{} {
	return c.Contract.CallContract(blockchain, from, project, request)
}

// ==================== BLOCKCHAIN FUNCTIONS (DELEGATED) ====================

// GetBlockCount retrieves the current block count of a blockchain.
func (c *Client) GetBlockCount(blockchain string) map[string]interface{} {
	return c.Blockchain.GetBlockCount(blockchain)
}

// GetBlock retrieves a specific block by number.
func (c *Client) GetBlock(blockchain string, number int) map[string]interface{} {
	return c.Blockchain.GetBlock(blockchain, number)
}

// GetBlockRange retrieves a range of blocks.
func (c *Client) GetBlockRange(blockchain string, start int, end int) map[string]interface{} {
	return c.Blockchain.GetBlockRange(blockchain, start, end)
}

// GetBlockchains retrieves the list of available blockchains.
func (c *Client) GetBlockchains() map[string]interface{} {
	return c.Blockchain.GetBlockchains()
}

// GetAnalytics retrieves analytics data for a blockchain.
func (c *Client) GetAnalytics(blockchain string) map[string]interface{} {
	return c.Blockchain.GetAnalytics(blockchain)
}

// ==================== TRANSACTION FUNCTIONS (DELEGATED) ====================

// SendTransaction sends a transaction to the blockchain.
func (c *Client) SendTransaction(id string, sender string, to string, timestamp string, transactionType string, payload string, nonce string, signature string, blockchain string) map[string]interface{} {
	return c.Blockchain.SendTransaction(id, sender, to, timestamp, transactionType, payload, nonce, signature, blockchain)
}

// GetPendingTransaction retrieves a pending transaction by ID.
func (c *Client) GetPendingTransaction(blockchain string, TxID string) map[string]interface{} {
	return c.Blockchain.GetPendingTransaction(blockchain, TxID)
}

// GetTransactionByID retrieves a transaction by ID within a block range.
func (c *Client) GetTransactionByID(blockchain string, TxID string, start int, end int) map[string]interface{} {
	return c.Blockchain.GetTransactionByID(blockchain, TxID, start, end)
}

// GetTransactionByNode retrieves transactions by node ID within a block range.
func (c *Client) GetTransactionByNode(blockchain string, node string, start int, end int) map[string]interface{} {
	return c.Blockchain.GetTransactionByNode(blockchain, node, start, end)
}

// GetTransactionByAddress retrieves all transactions involving a specific address.
func (c *Client) GetTransactionByAddress(blockchain string, address string, start int, end int) map[string]interface{} {
	return c.Blockchain.GetTransactionByAddress(blockchain, address, start, end)
}

// GetTransactionByDate retrieves transactions involving an address within a date range.
func (c *Client) GetTransactionByDate(blockchain string, address string, startDate string, endDate string) map[string]interface{} {
	return c.Blockchain.GetTransactionByDate(blockchain, address, startDate, endDate)
}

// GetTransactionOutcome polls for transaction finality with a timeout.
func (c *Client) GetTransactionOutcome(blockchain string, TxID string, timeoutSec int) string {
	return c.Blockchain.GetTransactionOutcome(blockchain, TxID, timeoutSec)
}

// ==================== ASSET FUNCTIONS (DELEGATED) ====================

// GetAssetList retrieves the list of assets on a blockchain.
func (c *Client) GetAssetList(blockchain string) map[string]interface{} {
	return c.Blockchain.GetAssetList(blockchain)
}

// GetAsset retrieves asset description.
func (c *Client) GetAsset(blockchain string, asset string) map[string]interface{} {
	return c.Blockchain.GetAsset(blockchain, asset)
}

// GetAssetSupply retrieves the supply of an asset.
func (c *Client) GetAssetSupply(blockchain string, asset string) map[string]interface{} {
	return c.Blockchain.GetAssetSupply(blockchain, asset)
}

// ==================== DOMAIN FUNCTIONS (DELEGATED) ====================

// GetDomain resolves a domain name to a wallet address.
func (c *Client) GetDomain(blockchain string, name string) map[string]interface{} {
	return c.Blockchain.GetDomain(blockchain, name)
}

// GetVoucher retrieves voucher information by code.
func (c *Client) GetVoucher(blockchain string, code string) map[string]interface{} {
	return c.Blockchain.GetVoucher(blockchain, code)
}

// ==================== CRYPTO FUNCTIONS (DELEGATED) ====================

// SignMessage signs a message with a private key.
func (c *Client) SignMessage(message string, privateKey string) map[string]interface{} {
	return c.Crypto.SignMessage(message, privateKey)
}

// VerifySignature verifies a signature against a message using the public key.
func (c *Client) VerifySignature(publicKey string, message string, signature string) bool {
	return c.Crypto.VerifySignature(publicKey, message, signature)
}

// GetPublicKey derives the public key from a private key.
func (c *Client) GetPublicKey(privateKey string) string {
	return c.Crypto.GetPublicKey(privateKey)
}

// GetKeysFromString generates a key pair from a seed phrase.
func (c *Client) GetKeysFromString(seedPhrase string) (map[string]string, error) {
	return c.Crypto.GetKeysFromString(seedPhrase)
}
