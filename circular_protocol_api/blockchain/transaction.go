/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'transaction.go', is part of the Circular Protocol Go SDK.
Module: blockchain/transaction
Purpose: Handles transaction operations.
*/

package blockchain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// SendTransaction sends a transaction to the blockchain.
func (s *Service) SendTransaction(id string, sender string, to string, timestamp string, transactionType string, payload string, nonce string, signature string, blockchain string) map[string]interface{} {
	data := map[string]interface{}{
		"ID":         utils.HexFix(id),
		"From":       utils.HexFix(sender),
		"To":         utils.HexFix(to),
		"Timestamp":  timestamp,
		"Type":       transactionType,
		"Payload":    payload,
		"Nonce":      nonce,
		"Signature":  signature,
		"Blockchain": utils.HexFix(blockchain),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.SEND_TRANSACTION, s.nagURL)
}

// GetPendingTransaction retrieves a pending transaction by ID.
func (s *Service) GetPendingTransaction(blockchain string, TxID string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"ID":         utils.HexFix(TxID),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_PENDING_TRANSACTION, s.nagURL)
}

// GetTransactionByID retrieves a transaction by ID within a block range.
// If end = 0 then start is the number of blocks from the last one minted.
func (s *Service) GetTransactionByID(blockchain string, TxID string, start int, end int) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"ID":         utils.HexFix(TxID),
		"Start":      strconv.Itoa(start),
		"End":        strconv.Itoa(end),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_TRANSACTION_BY_ID, s.nagURL)
}

// GetTransactionByNode retrieves transactions by node ID within a block range.
// If end = 0 then start is the number of blocks from the last one minted.
func (s *Service) GetTransactionByNode(blockchain string, node string, start int, end int) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"NodeID":     utils.HexFix(node),
		"Start":      strconv.Itoa(start),
		"End":        strconv.Itoa(end),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_TRANSACTION_BY_NODE, s.nagURL)
}

// GetTransactionByAddress retrieves all transactions involving a specific address.
// If end = 0 then start is the number of blocks from the last one minted.
func (s *Service) GetTransactionByAddress(blockchain string, address string, start int, end int) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"Start":      strconv.Itoa(start),
		"End":        strconv.Itoa(end),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_TRANSACTIONS_BY_ADDRESS, s.nagURL)
}

// GetTransactionByDate retrieves transactions involving an address within a date range.
func (s *Service) GetTransactionByDate(blockchain string, address string, startDate string, endDate string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Address":    utils.HexFix(address),
		"StartDate":  startDate,
		"EndDate":    endDate,
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_TRANSACTION_BY_DATE, s.nagURL)
}

// GetTransactionOutcome polls for transaction finality with a timeout.
func (s *Service) GetTransactionOutcome(blockchain string, TxID string, timeoutSec int) string {
	startTime := time.Now()
	interval := time.Duration(s.intervalSec) * time.Second
	timeout := time.Duration(timeoutSec) * time.Second

	var checkTransaction func() string
	checkTransaction = func() string {
		elapsedTime := time.Since(startTime)
		fmt.Println("Checking transaction...x", elapsedTime, timeout)
		if elapsedTime > timeout {
			fmt.Println("Timeout exceeded")
			return "Timeout exceeded"
		}
		data := s.GetTransactionByID(blockchain, TxID, 0, 10)

		fmt.Println("Data received:", data)

		if result, ok := data["Result"].(float64); ok && result == 200 {
			response, ok := data["Response"].(map[string]interface{})
			if ok && response["Status"] != "Pending" && response["Status"] != "Transaction Not Found" {
				fmt.Println("OK!")
				return response["Status"].(string)
			}
		}
		fmt.Println("Transaction not yet confirmed or not found, polling again...")
		time.Sleep(interval)
		return checkTransaction() // Recursive call
	}
	return checkTransaction()
}
