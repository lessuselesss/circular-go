package utils

import (
	newutils "github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// Type Aliases
type ECSignature = newutils.ECSignature
type ECDSASignature = newutils.ECDSASignature

// Constants
const (
	TEST_CONTRACT                 = newutils.TEST_CONTRACT
	CALL_CONTRACT                 = newutils.CALL_CONTRACT
	CHECK_WALLET                  = newutils.CHECK_WALLET
	GET_WALLET                    = newutils.GET_WALLET
	GET_LATEST_TRANSACTIONS       = newutils.GET_LATEST_TRANSACTIONS
	GET_WALLET_BALANCE            = newutils.GET_WALLET_BALANCE
	REGISTER_WALLET               = newutils.REGISTER_WALLET
	GET_DOMAIN                    = newutils.GET_DOMAIN
	GET_ASSET_LIST                = newutils.GET_ASSET_LIST
	GET_ASSET                     = newutils.GET_ASSET
	GET_ASSET_SUPPLY              = newutils.GET_ASSET_SUPPLY
	GET_VOUCHER                   = newutils.GET_VOUCHER
	GET_BLOCK_RANGE               = newutils.GET_BLOCK_RANGE
	GET_BLOCK                     = newutils.GET_BLOCK
	GET_BLOCK_COUNT               = newutils.GET_BLOCK_COUNT
	GET_ANALYTICS                 = newutils.GET_ANALYTICS
	GET_BLOCKCHAINS               = newutils.GET_BLOCKCHAINS
	GET_PENDING_TRANSACTION       = newutils.GET_PENDING_TRANSACTION
	GET_TRANSACTION_BY_ID         = newutils.GET_TRANSACTION_BY_ID
	GET_TRANSACTION_BY_NODE       = newutils.GET_TRANSACTION_BY_NODE
	GET_TRANSACTIONS_BY_ADDRESS   = newutils.GET_TRANSACTIONS_BY_ADDRESS
	GET_TRANSACTION_BY_DATE       = newutils.GET_TRANSACTION_BY_DATE
	SEND_TRANSACTION              = newutils.SEND_TRANSACTION
	GET_WALLET_NONCE              = newutils.GET_WALLET_NONCE
)

// Function Wrappers

func SendRequest(data interface{}, nagFunction string, nagURL string) map[string]interface{} {
	return newutils.SendRequest(data, nagFunction, nagURL)
}

func PadNumber(number int) string {
	return newutils.PadNumber(number)
}

func GetFormattedTimestamp() string {
	return newutils.GetFormattedTimestamp()
}

func SignMessage(message string, privateKey string) map[string]interface{} {
	return newutils.SignMessage(message, privateKey)
}

func StringToHex(str string) string {
	return newutils.StringToHex(str)
}

func HexToString(hexStr string) (string, error) {
	return newutils.HexToString(hexStr)
}

func HexFix(word interface{}) string {
	return newutils.HexFix(word)
}

func Sha256(data string) string {
	return newutils.Sha256(data)
}

func VerifySignature(publicKey string, message string, signature string) bool {
	return newutils.VerifySignature(publicKey, message, signature)
}

func GetPublicKey(privateKey string) string {
	return newutils.GetPublicKey(privateKey)
}

func GetKeysFromString(seedPhrase string) (map[string]string, error) {
	return newutils.GetKeysFromString(seedPhrase)
}
