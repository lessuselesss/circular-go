/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'block.go', is part of the Circular Protocol Go SDK.
Module: blockchain/block
Purpose: Handles block operations.
*/

package blockchain

import (
	"strconv"

	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// GetBlockCount retrieves the current block count of a blockchain.
func (s *Service) GetBlockCount(blockchain string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_BLOCK_COUNT, s.nagURL)
}

// GetBlock retrieves a specific block by number.
func (s *Service) GetBlock(blockchain string, number int) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain":  utils.HexFix(blockchain),
		"BlockNumber": strconv.Itoa(number),
		"Version":     s.version,
	}
	return utils.SendRequest(data, utils.GET_BLOCK, s.nagURL)
}

// GetBlockRange retrieves a range of blocks.
func (s *Service) GetBlockRange(blockchain string, start int, end int) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Start":      strconv.Itoa(start),
		"End":        strconv.Itoa(end),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_BLOCK_RANGE, s.nagURL)
}
