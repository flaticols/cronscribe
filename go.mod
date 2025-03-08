module github.com/flaticols/cronscribe

go 1.24.0

require (
	github.com/flaticols/cronscribe/pkg/ai v0.0.0
	github.com/flaticols/cronscribe/pkg/core v0.0.0
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/flaticols/cronscribe/pkg/ai => ./pkg/ai
	github.com/flaticols/cronscribe/pkg/core => ./pkg/core
)
