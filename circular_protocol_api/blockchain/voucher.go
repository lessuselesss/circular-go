/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'voucher.go', is part of the Circular Protocol Go SDK.
Module: blockchain/voucher
Purpose: Handles voucher operations.
*/

package blockchain

import (
	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// GetVoucher retrieves voucher information by code.
func (s *Service) GetVoucher(blockchain string, code string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Code":       utils.HexFix(code),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_VOUCHER, s.nagURL)
}
