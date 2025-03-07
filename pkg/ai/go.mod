module github.com/flaticols/cronscribe/pkg/ai

go 1.24.0

require (
	github.com/flaticols/cronscribe/pkg/core v0.0.0
	github.com/dlclark/regexp2 v1.10.0
	github.com/google/uuid v1.6.0
	github.com/pkoukk/tiktoken-go v0.1.6
	github.com/tmc/langchaingo v0.1.13
)

replace github.com/flaticols/cronscribe/pkg/core => ../core