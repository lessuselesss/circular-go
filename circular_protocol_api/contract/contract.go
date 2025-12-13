/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'contract/contract.go', defines the ContractService struct and implements
all smart contract-specific API functions for the Circular Protocol Go SDK. This
includes testing and calling smart contracts deployed on the Circular blockchain.
It is an internal service package used by the main Client Facade to interact with
the Circular Layer 1 Blockchain Protocol.

Package: github.com/circular-protocol/circular-go/circular_protocol_api/contract
Purpose: Smart contract operations for the Circular Protocol SDK
*/
package contract

import (
	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// Service provides smart contract operations for the Circular Protocol.
type Service struct {
	nagURL  string
	version string
}

// NewService creates a new ContractService instance.
func NewService(nagURL string, version string) *Service {
	return &Service{
		nagURL:  nagURL,
		version: version,
	}
}

// TestContract tests the execution of a smart contract.
func (s *Service) TestContract(blockchain string, sender string, project string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"From":       utils.HexFix(sender),
		"Timestamp":  utils.GetFormattedTimestamp(),
		"Project":    project,
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.TEST_CONTRACT, s.nagURL)
}

// CallContract calls a function on a smart contract.
func (s *Service) CallContract(blockchain string, from string, project string, request string) map[string]interface{} {
	data := map[string]interface{}{
		"Address":    utils.HexFix(project),
		"Blockchain": utils.HexFix(blockchain),
		"From":       utils.HexFix(from),
		"Timestamp":  utils.GetFormattedTimestamp(),
		"Request":    utils.StringToHex(request),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.CALL_CONTRACT, s.nagURL)
}
