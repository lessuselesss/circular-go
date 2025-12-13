/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'blockchain/blockchain.go', defines the BlockchainService struct.
The implementation of methods has been split into separate files:
- block.go
- transaction.go
- asset.go
- domain.go
- voucher.go
- analytics.go

Package: github.com/circular-protocol/circular-go/circular_protocol_api/blockchain
Purpose: Blockchain service struct definition
*/
package blockchain

// Service provides blockchain, transaction, and asset operations for the Circular Protocol.
type Service struct {
	nagURL      string
	version     string
	intervalSec int
}

// NewService creates a new BlockchainService instance.
func NewService(nagURL string, version string) *Service {
	return &Service{
		nagURL:      nagURL,
		version:     version,
		intervalSec: 5,
	}
}
