# AGENTS Guidelines for This Repository

This repository contains the Go SDK for the Circular Protocol API. When working on the project interactively with an agent (e.g., Codex CLI, Gemini CLI, Cursor, Claude Code, Open Code or other AI coding assistants), please follow the guidelines below to ensure a smooth development experience.

## 1. Development Environment Setup

* **Ensure Go 1.23.1 or higher** is installed:
  ```bash
  go version
  ```

* **Download dependencies** (Go modules will handle this automatically):
  ```bash
  go mod download
  ```

* **Tidy up dependencies** after changes:
  ```bash
  go mod tidy
  ```

## 2. Code Quality and Standards

* **Follow Go conventions**:
  - Use `gofmt` to format code (most editors do this automatically)
  - Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
  - Use meaningful package, function, and variable names

* **Format your code** before committing:
  ```bash
  go fmt ./...
  ```

* **Run the Go linter**:
  ```bash
  go vet ./...
  ```

* **Use `golangci-lint`** for comprehensive linting (recommended):
  ```bash
  golangci-lint run
  ```

* **Write documentation comments** - Every exported function, type, and package should have a comment:
  ```go
  // CircularAPI provides methods to interact with the Circular Protocol.
  type CircularAPI struct {
      // fields
  }

  // GetWallet retrieves wallet information from the blockchain.
  func (c *CircularAPI) GetWallet(blockchain, address string) (*Wallet, error) {
      // implementation
  }
  ```

## 3. Project Structure

* `circular_protocol_api/` - Main SDK package
* `utils/` - Utility functions and helpers
* `go.mod` - Go module definition
* `go.sum` - Dependency checksums

## 4. Testing and Verification

* **Write tests** for all new functionality - Test files should end with `_test.go`:
  ```go
  package circular_protocol_api

  import "testing"

  func TestGetWallet(t *testing.T) {
      // test implementation
  }
  ```

* **Run all tests**:
  ```bash
  go test ./...
  ```

* **Run tests with coverage**:
  ```bash
  go test -cover ./...
  ```

* **Detailed coverage report**:
  ```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out
  ```

* **Run specific tests**:
  ```bash
  go test -run TestFunctionName ./...
  ```

* **Run tests with race detection** (important for concurrent code):
  ```bash
  go test -race ./...
  ```

## 5. Dependencies

Current dependencies:
- `github.com/btcsuite/btcd/chaincfg/chainhash` (v1.0.1) - Bitcoin chain hashing
- `github.com/decred/dcrd/dcrec/secp256k1/v4` (v4.3.0) - Elliptic curve cryptography

Adding new dependencies:
```bash
go get github.com/some/package
go mod tidy
```

## 6. Useful Commands Recap

| Command                          | Purpose                                              |
| -------------------------------- | ---------------------------------------------------- |
| `go mod tidy`                    | Clean up and organize dependencies                   |
| `go mod download`                | Download module dependencies                         |
| `go fmt ./...`                   | Format all Go files                                  |
| `go vet ./...`                   | Run Go's built-in static analyzer                    |
| `go test ./...`                  | Run all tests                                        |
| `go test -cover ./...`           | Run tests with coverage report                       |
| `go test -race ./...`            | Run tests with race detector                         |
| `go build`                       | Build the package                                    |
| `go doc <package/function>`      | View documentation                                   |
| `golangci-lint run`              | Run comprehensive linter suite                       |

## 7. Building and Installing

* **Build the package** to verify it compiles:
  ```bash
  go build ./...
  ```

* **Install locally** (if creating a CLI):
  ```bash
  go install
  ```

## 8. Best Practices

* **Error handling** - Always handle errors explicitly:
  ```go
  result, err := someFunction()
  if err != nil {
      return nil, fmt.Errorf("failed to do something: %w", err)
  }
  ```

* **Context usage** - Use `context.Context` for cancellation and timeouts in long-running operations.

* **Avoid global state** - Prefer dependency injection over global variables.

* **Use interfaces** - Design with interfaces for better testability and flexibility.

## 9. API Documentation

* Refer to the [Circular Protocol Documentation](https://circular-protocol.gitbook.io/standard-apis) for API endpoints and usage patterns.
* Keep README.md examples up to date with any API changes.
* Use `go doc` to view package documentation locally.

## 10. CI/CD Considerations

Before pushing:
1. Format: `go fmt ./...`
2. Vet: `go vet ./...`
3. Test: `go test ./...`
4. Tidy: `go mod tidy`

---

Following these practices ensures that agent-assisted development remains efficient and maintains code quality. Go's tooling is excellent, so use it frequently to catch issues early!
