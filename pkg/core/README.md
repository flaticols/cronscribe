# CronScribe Core Package

This package provides the core functionality for CronScribe, allowing conversion of human-readable schedule descriptions to cron expressions using YAML-based rules.

## Features

- YAML-based rule system for defining patterns and translations
- Support for multiple languages
- Flexible and extensible rule definitions
- Minimal dependencies

## Installation

```bash
go get github.com/flaticols/cronscribe/pkg/core
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/flaticols/cronscribe/pkg/core"
)

func main() {
    // Create a new CronScribe instance with rules from a directory
    cs, err := core.New("./rules")
    if err != nil {
        log.Fatalf("Failed to create CronScribe: %v", err)
    }

    // Convert a human-readable expression to cron
    cronExpr, err := cs.Convert("every day at noon")
    if err != nil {
        log.Fatalf("Conversion error: %v", err)
    }

    fmt.Printf("Cron expression: %s\n", cronExpr)
}
```

## Rules Directory Structure

The rules directory should contain YAML files with rule definitions for different languages. Each file should follow this structure:

```yaml
language: en
rules:
  - name: daily-at-time
    pattern: every day at (\d+)(?::(\d+))?\s*(am|pm)?
    variables:
      hour: 1
      minute: 2
      ampm: 3
    dictionaries:
      hour: hour_values
      ampm: ampm_values
    format: "%minute %hour * * *"
    default_values:
      minute: "0"
# ... more rules

dictionaries:
  ampm_values:
    am: ""
    pm: "+12"
  # ... more dictionaries
```

See the examples in the `rules` directory for complete reference implementations.

## Dependencies

- `gopkg.in/yaml.v3`: For parsing YAML rule files