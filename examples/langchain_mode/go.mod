module github.com/flaticols/cronscribe/examples/langchain_mode

go 1.24.0

require (
	github.com/flaticols/cronscribe/pkg/ai v0.0.0
	github.com/flaticols/cronscribe/pkg/core v0.0.0
	github.com/tmc/langchaingo v0.1.13
)

replace (
	github.com/flaticols/cronscribe/pkg/ai => ../../pkg/ai
	github.com/flaticols/cronscribe/pkg/core => ../../pkg/core
)