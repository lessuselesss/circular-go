/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'asset.go', is part of the Circular Protocol Go SDK.
Module: blockchain/asset
Purpose: Handles asset operations.
*/

package blockchain

import (
	"github.com/circular-protocol/circular-go/circular_protocol_api/utils"
)

// GetAssetList retrieves the list of assets on a blockchain.
func (s *Service) GetAssetList(blockchain string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_ASSET_LIST, s.nagURL)
}

// GetAsset retrieves asset description.
func (s *Service) GetAsset(blockchain string, asset string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"AssetName":  utils.HexFix(asset),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_ASSET, s.nagURL)
}

// GetAssetSupply retrieves the supply of an asset.
func (s *Service) GetAssetSupply(blockchain string, asset string) map[string]interface{} {
	data := map[string]interface{}{
		"Blockchain": utils.HexFix(blockchain),
		"AssetName":  utils.HexFix(asset),
		"Version":    s.version,
	}
	return utils.SendRequest(data, utils.GET_ASSET_SUPPLY, s.nagURL)
}
