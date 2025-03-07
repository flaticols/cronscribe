module github.com/flaticols/cronscribe

go 1.24.0

require (
	github.com/flaticols/cronscribe/pkg/core v0.0.0
	github.com/flaticols/cronscribe/pkg/ai v0.0.0
	gopkg.in/yaml.v3 v3.0.1
)

replace (
	github.com/flaticols/cronscribe/pkg/core => ./pkg/core
	github.com/flaticols/cronscribe/pkg/ai => ./pkg/ai
)