# Circular Protocol - Go SDK

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

The **Circular Protocol Go SDK** is the official Go library for seamless integration with the Circular blockchain ecosystem. This open-source SDK provides a comprehensive suite of tools for efficient and secure interaction with blockchain networks, managing wallets, assets, smart contracts, and more.

## 🔥 Key Features

- **Blockchain Interaction**: Connect and interact with Circular's blockchain networks
- **Smart Contracts**: Deploy, test, and interact with smart contracts
- **Wallet Management**: Create, retrieve, and manage blockchain wallets with balance tracking
- **Asset Management**: Issue and manage assets, handle transfers, and retrieve supply information
- **Domain Management**: Resolve blockchain domain names to wallet addresses
- **Transaction Management**: Send transactions, track status, and search the blockchain
- **Analytics**: Access blockchain performance data and insights
- **Cryptographic Helpers**: Built-in utilities for key generation, signing, and hashing
- **Go Support**: Idiomatic Go implementation with strong typing

```bash
go get github.com/circular-protocol/circular-go
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/circular-protocol/circular-go/circular_protocol_api"
)

func main() {
    // Initialize the API client
    api := circular_protocol_api.NewClient(
        "https://nag.circularlabs.io/NAG.php?cep=", // NAG URL
    )

    // Method: Request object (Go-style, type-safe)
    result := api.CheckWallet(
        "MainNet",
        "0xd55872dbe508fd27445889b9d81bbc9411bb0f1353153a249f2fb34ef2690310",
    )

    fmt.Println("Wallet response:", result)
}
```

### Key Features

- **Modular Design**: Composed of specialized services (Wallet, Contract, Blockchain, Crypto)
- **Auto-Preprocessing**: Hex values automatically normalized ('0x' prefix optional)
- **Version Auto-Injection**: No need to specify version in requests

## 📜 API Reference

The Circular Protocol Go SDK provides **39 methods** across multiple categories for comprehensive blockchain interaction.

### Wallet Operations (5 methods)

- **`CheckWallet`** - Verify wallet existence on the blockchain
- **`GetWallet`** - Retrieve complete wallet details and metadata
- **`GetLatestTransaction`** - Get recent wallet activity and transaction history
- **`GetWalletBalance`** - Query current wallet balance across assets
- **`GetWalletNonce`** - Get transaction nonce for the wallet

### Transaction Operations (6 methods)

- **`SendTransaction`** - Submit new transaction to the blockchain
- **`GetPendingTransaction`** - Check transaction status in the mempool
- **`GetTransactionByID`** - Query transaction by unique identifier
- **`GetTransactionByNode`** - Query transactions by validator node
- **`GetTransactionByAddress`** - Query all transactions for a wallet address
- **`GetTransactionByDate`** - Query transactions within a date range

### Block Operations (4 methods)

- **`GetBlock`** - Retrieve block data by block number or hash
- **`GetBlockRange`** - Query multiple blocks within a range
- **`GetBlockCount`** - Get current blockchain height (latest block number)
- **`GetAnalytics`** - Retrieve blockchain performance metrics and analytics

### Contract Operations (2 methods)

- **`TestContract`** - Validate smart contract logic before deployment
- **`CallContract`** - Execute smart contract function call

### Asset Operations (4 methods)

- **`GetAssetList`** - List all available assets on the blockchain
- **`GetAsset`** - Get detailed asset information and metadata
- **`GetAssetSupply`** - Query total and circulating supply for an asset
- **`GetVoucher`** - Retrieve voucher data and redemption details

### Domain Operations (1 method)

- **`GetDomain`** - Query blockchain domain registry (resolve domain to address)

### Network Operations (1 method)

- **`GetBlockchains`** - List all supported blockchain networks

---

### Cryptographic Helpers (5 methods)

- **`SignMessage`** - Generate ECDSA secp256k1 signatures (DER format)
- **`VerifySignature`** - Verify message signatures against public keys
- **`GetPublicKey`** - Derive public key from private key (128 hex characters, uncompressed, no 0x04 prefix)
- **`GetKeysFromString`** - Generate key pair from seed phrase

**Implementation Details:**
- **Go**: `btcsuite/btcd/btcec/v2` secp256k1

---

### Advanced Helpers (3 methods)

- **`GetTransactionOutcome`** - Poll for transaction confirmation with automatic retries

**Transaction Polling Behavior:**
- Checks transaction status every **5 seconds** (configurable via `intervalSec`)
- Returns successfully when transaction has `BlockNumber > 0` (confirmed)
- Throws timeout error after **120 seconds** (configurable via `timeoutSec`)
- Handles "pending" status gracefully with automatic retries

---

### Convenience Methods (1 method)

- **`RegisterWallet`** - Simplified wallet registration (wraps `SendTransaction`)

**Implementation:**
- Automatically derives `From` and `To` addresses via `SHA256(publicKey)`
- Constructs transaction payload: `{"Action": "CP_WALLET", "PublicKey": "..."}`
- Sets default values: `Nonce="00000000"`, `Type="C"`, `Signature="0000..."`
- Calculates transaction ID as SHA-256 hash of transaction fields
- Returns same response structure as `SendTransaction`

---

## 📊 Total Methods: 39

- **23** API Endpoint Methods
- **5** Cryptographic Helpers
- **3** Configuration Methods (GetNAGURL, SetNAGURL, GetNAGKey, SetNAGKey)
- **1** Convenience Method

> **Note**: For detailed parameter types, response structures, and advanced usage examples, refer to the **[Go SDK Documentation](https://circular-protocol.gitbook.io/circular-sdk/api-docs/go)**.

## 🤝 Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](https://github.com/circular-protocol/circular-canonical/blob/main/CONTRIBUTING.md) file in the canonical repository for guidelines.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📚 Resources

- **[Go SDK Documentation](https://circular-protocol.gitbook.io/circular-sdk/api-docs/go)** - Complete API reference
- **[Circular Protocol Docs](https://circular-protocol.gitbook.io)** - Protocol documentation
- **[Circular Canonical](https://github.com/circular-protocol/circular-canonical)** - Single source of truth

## ℹ️ About

**Version**: 1.0.x
**License**: MIT
**Maintained**: Manually maintained to ensure compatibility with circular-js-npm while adding Go enhancements

---

© 2025 Circular Global Ledgers, Inc. - Open source for private and commercial use
