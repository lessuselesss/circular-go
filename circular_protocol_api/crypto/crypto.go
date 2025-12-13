/*
LLMs.txt/AGENTS.md Metadata:
Circular Protocol - Standard APIs SDK
This file, 'crypto/crypto.go', defines the CryptoService struct and implements
all cryptographic operations for the Circular Protocol Go SDK. This includes
key generation, message signing, and signature verification using secp256k1
elliptic curve cryptography. It is an internal service package used by the main
Client Facade to interact with the Circular Layer 1 Blockchain Protocol.

Package: github.com/circular-protocol/circular-go/circular_protocol_api/crypto
Purpose: Cryptographic operations for the Circular Protocol SDK
*/
package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// ECDSASignature defines the structure for DER encoded signature
type ECDSASignature struct {
	R, S *big.Int
}

// Service provides cryptographic operations for the Circular Protocol.
type Service struct {
	version string
}

// NewService creates a new CryptoService instance.
func NewService(version string) *Service {
	return &Service{
		version: version,
	}
}

// SignMessage signs a message with a private key in hex format and returns the signature in hex format.
func (s *Service) SignMessage(message string, privateKey string) map[string]interface{} {
	bytesPrivateKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return map[string]interface{}{
			"Result":   "500",
			"Response": "Error during the decoding of the private key",
		}
	}

	privKey := secp256k1.PrivKeyFromBytes(bytesPrivateKey)

	messageHash := chainhash.HashB([]byte(message))
	r, sigS, err := ecdsa.Sign(rand.Reader, privKey.ToECDSA(), messageHash)
	if err != nil {
		return map[string]interface{}{
			"Result":   "500",
			"Response": "Error during the signing of the message",
		}
	}

	derSignature, err := asn1.Marshal(ECDSASignature{R: r, S: sigS})
	if err != nil {
		return map[string]interface{}{
			"Result":   "500",
			"Response": "Error during the encoding of the signature",
		}
	}

	stringDERSignature := hex.EncodeToString(derSignature)
	return map[string]interface{}{"Signature": stringDERSignature, "R": r, "S": sigS}
}

// VerifySignature verifies a signature against a message using the public key.
func (s *Service) VerifySignature(publicKey string, message string, signature string) bool {
	if len(publicKey) != 130 {
		fmt.Println("Invalid public key length")
		return false
	}

	// Remove the 0x prefix from the public key
	remove0x := hexFix(publicKey)

	// Hash the message to verify with SHA256
	msgHash := sha256.Sum256([]byte(message))

	// Decode the public key from hex
	publicKeyBytes, err := hex.DecodeString(remove0x)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Parse the public key
	pubKey, err := secp256k1.ParsePubKey(publicKeyBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Convert btcec.PublicKey to ecdsa.PublicKey
	ecdsaPubKey := ecdsa.PublicKey{
		Curve: secp256k1.S256(),
		X:     pubKey.X(),
		Y:     pubKey.Y(),
	}

	// Decode the signature from hex
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Parse the signature from DER format
	var ecsig ECDSASignature
	_, err = asn1.Unmarshal(signatureBytes, &ecsig)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Verify the signature
	return ecdsa.Verify(&ecdsaPubKey, msgHash[:], ecsig.R, ecsig.S)
}

// GetPublicKey derives the public key from a private key.
func (s *Service) GetPublicKey(privateKey string) string {
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return ""
	}

	privKey := secp256k1.PrivKeyFromBytes(privKeyBytes)
	pubKey := privKey.PubKey()

	return hex.EncodeToString(pubKey.SerializeUncompressed())
}

// GetKeysFromString generates a key pair from a seed phrase.
func (s *Service) GetKeysFromString(seedPhrase string) (map[string]string, error) {
	// Generate private key using SHA256 hash of seed phrase
	hashHex := sha256Hex(seedPhrase)
	hash, err := hex.DecodeString(hashHex)
	if err != nil {
		return nil, err
	}

	privateKey := secp256k1.PrivKeyFromBytes(hash)
	publicKey := privateKey.PubKey()

	// Encode private key in DER format
	privateKeyDER := privateKey.Serialize()

	// Encode public key in DER format
	publicKeyDER := publicKey.SerializeUncompressed()

	// Generate address from public key
	addressDER := sha256Hex(bytesToHex(publicKey.SerializeUncompressed()))

	// Create a map to hold the results
	result := map[string]string{
		"privateKey": bytesToHex(privateKeyDER),
		"publicKey":  bytesToHex(publicKeyDER),
		"address":    addressDER,
		"seedPhrase": seedPhrase,
	}

	return result, nil
}

// Helper functions

func hexFix(word interface{}) string {
	switch v := word.(type) {
	case int:
		return fmt.Sprintf("%x", v)
	case string:
		if len(v) >= 2 && v[:2] == "0x" {
			return v[2:]
		}
		return v
	default:
		return ""
	}
}

func sha256Hex(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func bytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}
