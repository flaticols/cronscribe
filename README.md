# CronScribe

CronScribe is a Go library that converts human-readable schedule descriptions to cron expressions using a rule-based system with multilingual support.

## Features

- Rule-based pattern matching with regex
- Dictionary-based translation system
- Support for multiple languages (English, Dutch, Russian)
- Flexible and extensible rule definitions
- Transformation rules for time formats
- Minimal dependencies

## Installation

```bash
go get github.com/flaticols/cronscribe
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/flaticols/cronscribe"
)

func main() {
    // Create a new CronScribe instance
    cs, err := cronscribe.New()
    if err != nil {
        log.Fatalf("Failed to create CronScribe: %v", err)
    }

    // Convert a human-readable expression to cron
    cronExpr, err := cs.Convert("every day at noon")
    if err != nil {
        log.Fatalf("Conversion error: %v", err)
    }

    fmt.Printf("Cron expression: %s\n", cronExpr)
    
    // You can also specify the language
    cronExpr, err = cs.ConvertWithLang("elke dag om 12 uur", "nl")
    if err != nil {
        log.Fatalf("Conversion error: %v", err)
    }
    
    fmt.Printf("Dutch cron expression: %s\n", cronExpr)
}
```

## Project Structure

```
cronscribe/
├── cronscribe.go       # Main package entry point
├── mapper.go           # Language mapping functionality
├── translator.go       # Rule-based translation core
├── pkg/
│   ├── langs/          # Language-specific implementations
│   │   ├── en/         # English language support
│   │   ├── nl/         # Dutch language support
│   │   └── ru/         # Russian language support
│   └── rules/          # Core rule system
│       ├── dictionaries.go  # Constants and dictionary definitions
│       └── rules.go    # Rule processing engine
```

## Constants System

CronScribe uses a constants-based approach to ensure consistency and maintainability across the codebase. This approach eliminates string literal duplication, reduces the risk of typos, and makes the code more robust. The constants system helps with:

1. **Consistency**: Using the same identifiers for dictionary keys throughout the codebase
2. **Maintainability**: Changing a value in one place affects all usages
3. **Type Safety**: Compiler can catch typos and misuses
4. **Readability**: Self-documenting code with descriptive constant names
5. **Internationalization**: Clear organization of language-specific terms

Constants are organized into logical groups:

### Core Constants (in pkg/rules/dictionaries.go)

```go
// Dictionary names
const (
    DictWeekdays  = "weekdays"
    DictOrdinals  = "ordinals"
    DictTimeAmPm  = "time_ampm"
    DictMonths    = "months"
)

// Variable names
const (
    VarHour     = "hour"
    VarMinute   = "minute"
    VarDay      = "day"
    VarWeekday  = "weekday"
    VarMonth    = "month"
    VarAmPm     = "ampm"
    VarOrdinal  = "ordinal"
    // ...more variable names
)

// Time period values
const (
    TimeAm = "am"
    TimePm = "pm"
)

// Special values
const (
    OrdinalLast = "last"
)
```

### Language Identifier Constants (in mapper.go)

```go
const (
    LangEN = "en" // English
    LangNL = "nl" // Dutch
    LangRU = "ru" // Russian
)
```

### Language-Specific Constants

Each language module can define its own constants for language-specific terms, such as:

- Dutch (nl) - Time periods and ordinals
- Russian (ru) - Grammatical cases and gender variations

These constants ensure consistency across the codebase, eliminate string literal duplication, and make the code more maintainable.

## Rule System

See the detailed documentation in [pkg/rules/README.md](pkg/rules/README.md) for information on:

- Rule structure and components
- Pattern design with regular expressions
- Variable mapping and transformations
- Special case handling
- Dictionary lookups

## Adding New Languages

To add a new language:

1. Create a new package under `pkg/langs/[language-code]/`
2. Define language-specific constants for special cases:
   ```go
   // constants for the new language
   const (
       // Time period formats
       TimeCustomFormat1 = "value1"
       TimeCustomFormat2 = "value2"
       
       // Ordinals specific to this language
       OrdinalCustom1 = "value3"
       OrdinalCustom2 = "value4"
   )
   ```
3. Create rule sets following the pattern in existing language modules
4. Add the language to the mapper in `mapper.go`:
   ```go
   const (
       // Existing language constants
       LangEN = "en" // English
       LangNL = "nl" // Dutch
       LangRU = "ru" // Russian
       
       // Add your new language
       LangXX = "xx" // Your language name
   )
   
   func LoadDefaultRules() (map[string]*rules.RuleSet, error) {
       // ...
       allRules[LangXX] = xx.YourLanguageRuleSet
       // ...
   }
   ```
5. Use the existing constants from `pkg/rules/dictionaries.go` for standard dictionary keys and variable names
6. Only add language-specific constants for terms that are unique to the new language

## Dependencies

- `gopkg.in/yaml.v3`: For parsing rule configurations