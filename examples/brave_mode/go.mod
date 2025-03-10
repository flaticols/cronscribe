module github.com/flaticols/cronscribe/examples/brave_mode

go 1.24.0

require (
	github.com/flaticols/cronscribe/pkg/ai v0.0.0
	github.com/flaticols/cronscribe/pkg/core v0.0.0
	github.com/sashabaranov/go-openai v1.38.0
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/flaticols/cronscribe/pkg/ai => ../../pkg/ai
	github.com/flaticols/cronscribe/pkg/core => ../../pkg/core
)