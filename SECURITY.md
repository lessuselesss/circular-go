# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |
| < Latest| :x:                |

## Reporting a Vulnerability

We take the security of the Circular Protocol Go SDK seriously. If you believe you have found a security vulnerability, please report it to us as described below.

### Where to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report them via email to:
- **Email**: security@circularlabs.io

If you prefer encrypted communication, please request our PGP key.

### What to Include

Please include the following information in your report:

- **Type of vulnerability** (e.g., cryptographic weakness, input validation issue, etc.)
- **Full paths of source file(s)** related to the manifestation of the vulnerability
- **Location of the affected source code** (tag/branch/commit or direct URL)
- **Step-by-step instructions** to reproduce the issue
- **Proof-of-concept or exploit code** (if possible)
- **Impact of the vulnerability**, including how an attacker might exploit it
- **Your contact information** for follow-up questions

### Response Timeline

- **Initial Response**: Within 48 hours of report submission
- **Vulnerability Assessment**: Within 5 business days
- **Fix Timeline**: Depends on severity and complexity
  - Critical: Within 7 days
  - High: Within 14 days
  - Medium: Within 30 days
  - Low: Next scheduled release

### Security Update Process

1. **Confirmation**: We confirm the vulnerability and determine its severity
2. **Fix Development**: We develop a fix in a private repository
3. **Testing**: Thorough testing of the fix
4. **Release**: Security patch released with credit to reporter (unless anonymity requested)
5. **Disclosure**: Public disclosure after patch is available

### Security Best Practices

When using the Circular Protocol Go SDK:

#### Private Key Management

**NEVER** store private keys in:
- Source code or version control
- Environment variables in public repositories
- Client-side code or frontend applications
- Log files or error messages
- Configuration files committed to git
- Global variables without protection

**DO** store private keys in:
- Secure key management systems (AWS Secrets Manager, HashiCorp Vault, etc.)
- Hardware security modules (HSMs)
- Encrypted environment variables (with restricted access)
- Secure secrets management services
- Encrypted files with proper file permissions (0600)

```go
package main

import (
    "fmt"
    "os"
)

// ❌ NEVER DO THIS
const privateKey = "c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3"

// ✅ DO THIS
func getPrivateKey() (string, error) {
    privateKey := os.Getenv("WALLET_PRIVATE_KEY")
    if privateKey == "" {
        return "", fmt.Errorf("private key not configured")
    }
    return privateKey, nil
}

// ✅ BETTER - Use a secrets manager
// import "github.com/aws/aws-sdk-go/service/secretsmanager"
// func getPrivateKey(ctx context.Context, sm *secretsmanager.SecretsManager) (string, error) {
//     result, err := sm.GetSecretValueWithContext(ctx, &secretsmanager.GetSecretValueInput{
//         SecretId: aws.String("wallet-private-key"),
//     })
//     if err != nil {
//         return "", err
//     }
//     return *result.SecretString, nil
// }
```

#### Input Validation

Always validate and sanitize inputs:

```go
package circular

import (
    "fmt"
    "regexp"
    "strings"
)

var addressRegex = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)

// IsValidAddress validates a Circular Protocol address
func IsValidAddress(address string) bool {
    // Remove 0x prefix if present
    cleanAddress := strings.TrimPrefix(address, "0x")
    return addressRegex.MatchString(cleanAddress)
}

// ValidateBlockchain validates blockchain identifier
func ValidateBlockchain(blockchain string) error {
    validBlockchains := map[string]bool{
        "Circular Main Public":      true,
        "Circular Secondary Public": true,
        "Circular Documark Public":  true,
        "Circular SandBox":          true,
    }
    
    if !validBlockchains[blockchain] {
        return fmt.Errorf("invalid blockchain identifier: %s", blockchain)
    }
    return nil
}

// Usage
if !IsValidAddress(address) {
    return fmt.Errorf("invalid address format: %s", address)
}

if err := ValidateBlockchain(blockchain); err != nil {
    return err
}
```

#### Transaction Verification

Always verify transaction parameters before signing:

```go
// Verify transaction before signing
nonce, err := api.GetWalletNonce(ctx, blockchain, fromAddress)
if err != nil {
    return fmt.Errorf("failed to get nonce: %w", err)
}

log.Printf("Transaction Details:")
log.Printf("From: %s", fromAddress)
log.Printf("To: %s", toAddress)
log.Printf("Amount: %s", amount)
log.Printf("Nonce: %d", nonce)

// Confirm before proceeding
signature, err := api.SignMessage(transactionHash, privateKey)
if err != nil {
    return fmt.Errorf("failed to sign: %w", err)
}
```

#### Network Security

- **Use HTTPS**: Always use HTTPS endpoints for NAG communication
- **Verify TLS Certificates**: Never disable certificate verification
- **Context with Timeout**: Always use context with timeout for HTTP requests
- **Rate Limiting**: Implement rate limiting using golang.org/x/time/rate

```go
import (
    "context"
    "net/http"
    "time"
)

// ✅ Good - uses HTTPS with timeout context
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, "GET", 
    "https://nag.circularlabs.io/NAG.php?cep=", nil)
if err != nil {
    return err
}

client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        // ✅ Keep TLS verification enabled
        TLSClientConfig: nil, // Uses default secure config
    },
}

resp, err := client.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()

// ❌ NEVER DO THIS
// transport := &http.Transport{
//     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// }
```

#### Error Handling

Never expose sensitive information in errors:

```go
// ❌ DON'T expose sensitive data
// return fmt.Errorf("failed to sign with key %s: %v", privateKey, err)

// ✅ DO use error wrapping without sensitive data
return fmt.Errorf("failed to sign transaction: %w", err)

// ✅ Log securely (not to stdout in production)
// Use structured logging
import "log/slog"

logger := slog.Default()
logger.Error("transaction signing failed",
    "error", err,
    "address", fromAddress, // OK to log
    // "private_key", key,  // NEVER log this
)
```

#### Dependency Security

- Keep dependencies up to date using `go get -u`
- Use `go mod tidy` to remove unused dependencies
- Check for known vulnerabilities with `govulncheck`
- Pin dependencies with `go.mod` and `go.sum`

```bash
# Check for known vulnerabilities
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# Update dependencies
go get -u ./...
go mod tidy

# Verify module checksums
go mod verify
```

#### Concurrency Safety

Go's concurrency features require careful handling:

```go
// ✅ Use sync.Mutex for shared state
type SafeWallet struct {
    mu      sync.RWMutex
    balance int64
}

func (w *SafeWallet) GetBalance() int64 {
    w.mu.RLock()
    defer w.mu.RUnlock()
    return w.balance
}

func (w *SafeWallet) SetBalance(balance int64) {
    w.mu.Lock()
    defer w.mu.Unlock()
    w.balance = balance
}

// ✅ Use channels for communication
// ✅ Run race detector in tests: go test -race ./...
```

## Known Security Considerations

### Cryptographic Operations

This SDK uses:
- **secp256k1** elliptic curve via `github.com/decred/dcrd/dcrec/secp256k1/v4`
- **Bitcoin chain hashing** from `github.com/btcsuite/btcd/chaincfg/chainhash`
- **Go's crypto package** for standard operations

These are industry-standard cryptographic primitives. However:

- **Randomness**: Go's `crypto/rand` provides cryptographically secure random numbers
- **Side-channel attacks**: Be cautious with timing attacks in comparisons
- **Memory**: Sensitive data may remain in memory; consider using secure memory handling

```go
import "crypto/subtle"

// ✅ Use constant-time comparison for sensitive data
func compareSignatures(a, b []byte) bool {
    return subtle.ConstantTimeCompare(a, b) == 1
}

// ❌ DON'T use regular comparison
// if bytes.Equal(expectedSig, actualSig) { ... }
```

### Server-Side Usage

**⚠️ This SDK is designed for server-side use**

Best practices:
- Use proper authentication middleware
- Implement rate limiting
- Use context for cancellation and timeouts
- Enable security headers in HTTP responses

### Go Version Compatibility

- **Minimum Go 1.23.1** is required
- Keep Go updated to receive security patches
- Use the latest stable Go version when possible

```bash
# Check Go version
go version

# Update Go (using official installer or go itself)
go install golang.org/dl/go1.23.1@latest
go1.23.1 download
```

### Memory Safety

Go provides memory safety, but be aware of:

```go
// ✅ Clear sensitive data when done
func clearPrivateKey(key []byte) {
    for i := range key {
        key[i] = 0
    }
}

privateKey := []byte("sensitive-key")
defer clearPrivateKey(privateKey)
// Use privateKey...
```

### Testing

- Always run tests with race detector: `go test -race ./...`
- Use fuzzing for input validation: `go test -fuzz=Fuzz ./...`
- Test error paths and edge cases
- Use table-driven tests for comprehensive coverage

```go
func TestIsValidAddress(t *testing.T) {
    tests := []struct {
        name    string
        address string
        want    bool
    }{
        {"valid address", "a1b2c3...", true},
        {"invalid address", "invalid", false},
        // Add more test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsValidAddress(tt.address); got != tt.want {
                t.Errorf("IsValidAddress() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Security Advisories

Security advisories will be published at:
- [GitHub Security Advisories](https://github.com/circular-protocol/circular-go/security/advisories)
- Release notes with `[SECURITY]` tag
- CHANGELOG.md with security section

## Bug Bounty Program

Currently, we do not have a formal bug bounty program. However:
- Significant vulnerabilities may be eligible for recognition
- Contributors will be credited in release notes (unless anonymity requested)
- We appreciate responsible disclosure

## Contact

For security concerns or questions:
- **Security Email**: security@circularlabs.io
- **General Support**: support@circularlabs.io
- **GitHub Issues**: For non-security bugs only

## Additional Resources

- [Go Security Best Practices](https://golang.org/doc/security)
- [Secure Go Development](https://go.dev/security/)
- [Circular Protocol Documentation](https://circular-protocol.gitbook.io/standard-apis)

## Acknowledgments

We would like to thank the following individuals for responsibly disclosing security vulnerabilities:

*No vulnerabilities reported yet*

---

**Last Updated**: 2025-12-13
