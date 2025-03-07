# CronScribe Development Guidelines

## Commands
- **Build**: `go build`
- **Test all**: `go test ./...`
- **Test file**: `go test ./file_test.go`
- **Test single**: `go test -run TestName`
- **Test verbose**: `go test -v ./...` 
- **Lint**: `golangci-lint run`
- **Format**: `gofmt -w .`

## Code Style
- **Imports**: Standard library first, then external packages
- **Naming**: 
  - CamelCase for exported items
  - lowerCamelCase for unexported items
  - Descriptive function names that indicate functionality
- **Error handling**: Return errors as last value, use `fmt.Errorf` with `%w` for wrapping
- **Testing**: Table-driven tests with clear test cases
- **Comments**: Document exported functions, types and methods
- **Project structure**:
  - Core code in root directory
  - Configuration files in `rules/` directory
  - Test files with `_test.go` suffix
- **Dependencies**: Minimal external dependencies (currently only `gopkg.in/yaml.v3`)