/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'analytics.go', is part of the Circular Protocol Go SDK.
Module: blockchain/analytics
Purpose: Handles analytics operations.
*/

package blockchain

import (
	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// GetBlockchains retrieves the list of available blockchains.
func (s *Service) GetBlockchains() map[string]interface{} {
	data := map[string]interface{}{}
	return utils.SendRequest(data, utils.GET_BLOCKCHAINS, s.nagURL)
}

// GetAnalytics retrieves analytics data for a blockchain.
func (s *Service) GetAnalytics(blockchain string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_ANALYTICS, s.nagURL)
}
