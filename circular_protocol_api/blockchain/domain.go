/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'domain.go', is part of the Circular Protocol Go SDK.
Module: blockchain/domain
Purpose: Handles domain resolution operations.
*/

package blockchain

import (
	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// GetDomain resolves a domain name to a wallet address.
func (s *Service) GetDomain(blockchain string, name string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Name":       utils.HexFix(name),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_DOMAIN, s.nagURL)
}
