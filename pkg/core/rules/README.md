# CronScribe Rules Documentation

CronScribe translates natural language expressions into cron syntax using rules defined in YAML files. This comprehensive guide explains how to create, customize, and optimize rules for various languages and scheduling patterns.

## Introduction to Cron Syntax

Before creating rules, it's important to understand standard cron syntax:

```
┌───────────── minute (0-59)
│ ┌───────────── hour (0-23)
│ │ ┌───────────── day of the month (1-31)
│ │ │ ┌───────────── month (1-12)
│ │ │ │ ┌───────────── day of the week (0-6) (Sunday to Saturday)
│ │ │ │ │
* * * * *
```

CronScribe supports extended cron syntax, including:
- `L` - Last day of month or last specific weekday (`5L` = last Friday)
- `#` - Nth occurrence of a weekday (`1#3` = third Monday)
- `W` - Nearest weekday (`15W` = weekday nearest to the 15th)
- `/` - Step values (`*/5` = every 5 minutes)

## Rule File Structure

Each language has its own YAML file (e.g., `en.yaml`, `nl.yaml`, `ru.yaml`) with the following structure:

```yaml
language: en  # Language code (ISO 639-1)
rules:
  # Array of rule definitions
  - name: rule_name
    # Rule properties...
  - name: another_rule
    # More rule properties...
dictionaries:
  # Dictionary definitions for this language
  dictionary_name:
    key: "value"
```

Files must be named with the language code followed by `.yaml` extension.

## Rule Components in Detail

### Basic Rule Properties

Every rule requires these components:

```yaml
- name: daily_at_time  # Unique identifier for the rule
  pattern: '(?i)(?:each|every)\s+day\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?'  # Regex pattern
  variables:  # Maps regex capture groups to named variables
    hour: 1   # '1' refers to the first capture group in the pattern
    minute: 2 # '2' refers to the second capture group
    ampm: 3   # '3' refers to the third capture group
  dictionaries:  # Maps variables to dictionary lookups for translation
    ampm: time_ampm  # 'ampm' variable will be looked up in the 'time_ampm' dictionary
  format: "%minute %hour * * *"  # Cron format template with variables
  default_values:  # Fallback values for optional or missing variables
    minute: "0"    # If 'minute' is not captured, use "0"
```

#### Pattern Design (Regex)

The `pattern` field uses regular expressions to match natural language:

- `(?i)` flag for case-insensitivity
- Non-capturing groups `(?:...)` for grouping alternatives
- Capturing groups `(...)` only for values to extract
- `\s+` for flexible whitespace
- `?` for optional components

Common patterns:
- `(\d+)` - One or more digits
- `(word1|word2|word3)` - One of several options
- `(?:\s+at\s+(\d+))` - Optional time specification

#### Variables

The `variables` map connects regex capture groups to named variables:

```yaml
variables:
  hour: 1   # First capture group → hour variable
  minute: 2 # Second capture group → minute variable
```

Use consistent variable names across all rules (e.g., always use `hour` not `hr`).

#### Format String

The `format` string is a template for the cron expression:

```yaml
format: "%minute %hour * * %weekday"
```

Each `%variable` is replaced with its value. Fixed values (like `*`) are written directly.

### Advanced Rule Properties

#### Dictionaries

Dictionaries translate human terms to cron values:

```yaml
dictionaries:
  weekday: weekdays  # Use 'weekdays' dictionary for 'weekday' variable
```

At the file level, define your dictionaries:

```yaml
dictionaries:
  weekdays:
    sunday: "0"
    monday: "1"
    # ...
  
  months:
    january: "1"
    february: "2"
    # ...
```

Dictionary lookups are case-insensitive and normalize input.

#### Transformations

Transformations modify variable values based on conditions:

```yaml
transformations:
  hour:  # Transform the 'hour' variable
    - condition: "ampm == 'pm' && hour < 12"  # Condition to evaluate
      operation: "hour + 12"  # Mathematical operation to perform
    - condition: "ampm == 'am' && hour == 12"
      operation: "0"
```

##### Condition Syntax

Supported operators in conditions:
- `==`, `!=` - Equality checks
- `<`, `>`, `<=`, `>=` - Numeric comparisons
- `&&` - Logical AND
- `||` - Logical OR

Variables in conditions are replaced with their actual values before evaluation.

##### Operation Syntax

Operations can include:
- Arithmetic: `hour + 12`, `minute - 30`
- String literals: `"0"`, `"L"`
- Variable references: `weekday`

Transformations are processed sequentially; only the first matching transformation for each variable is applied.

#### Special Cases

Special cases provide alternative formats based on conditions:

```yaml
special_cases:
  - condition: "ordinal == 'last'"  # If the ordinal is "last"
    format: "0 0 * * %weekdayL"  # Use L suffix for last occurrence
```

If the condition is true, the alternative format replaces the main format.

#### Default Values

Default values provide fallbacks for missing components:

```yaml
default_values:
  minute: "0"  # If minute isn't specified, use 0
  hour: "0"    # If hour isn't specified, use 0
```

Default values are applied before transformations and special cases.

## Detailed Examples with Explanations

### Example 1: Daily Schedule

```yaml
- name: daily_at_time
  pattern: '(?i)(?:each|every)\s+day\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?'
  variables:
    hour: 1
    minute: 2
    ampm: 3
  dictionaries:
    ampm: time_ampm
  format: "%minute %hour * * *"
  default_values:
    minute: "0"
  transformations:
    hour:
      - condition: "ampm == 'pm' && hour < 12"
        operation: "hour + 12"
      - condition: "ampm == 'am' && hour == 12"
        operation: "0"
```

**Step-by-step breakdown:**

1. **Pattern breakdown**:
   - `(?i)` - Case insensitive matching
   - `(?:each|every)` - Match "each" or "every" without capturing
   - `\s+day\s+at\s+` - Match "day at" with flexible spaces
   - `(\d+)` - Capture hour digits (group 1)
   - `(?::(\d+))?` - Optionally capture minutes after colon (group 2)
   - `\s*(am|pm)?` - Optionally capture AM/PM (group 3)

2. **Processing flow**:
   - Extract variables from regex groups
   - Apply default values for missing variables
   - Apply transformations based on conditions
   - Format the cron string with variables

3. **Transformation logic**:
   - For PM times (1pm-11pm): add 12 to hour (1pm → 13)
   - For 12am (midnight): convert to 0
   - For AM times and 12pm: no transformation needed

**Examples with detailed conversion:**

- Input: "every day at 9am"
  - Matches: `hour=9, minute=null, ampm=am`
  - Default values: `minute=0`
  - No transformations apply
  - Result: "0 9 * * *" (At 9:00 AM, every day)

- Input: "each day at 3pm"
  - Matches: `hour=3, minute=null, ampm=pm`
  - Default values: `minute=0`
  - Transformation: `hour=3+12=15` (pm and hour<12)
  - Result: "0 15 * * *" (At 3:00 PM, every day)

- Input: "every day at 12:30am"
  - Matches: `hour=12, minute=30, ampm=am`
  - Transformation: `hour=0` (am and hour=12)
  - Result: "30 0 * * *" (At 12:30 AM, every day)

### Example 2: Weekly Schedule

```yaml
- name: weekly_day_at_time
  pattern: '(?i)(?:each|every)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?'
  variables:
    weekday: 1
    hour: 2
    minute: 3
    ampm: 4
  dictionaries:
    weekday: weekdays
    ampm: time_ampm
  format: "0 %minute %hour * * %weekday"
  default_values:
    minute: "0"
    hour: "0"
  transformations:
    hour:
      - condition: "ampm == 'pm' && hour < 12"
        operation: "hour + 12"
      - condition: "ampm == 'am' && hour == 12"
        operation: "0"
```

**Step-by-step breakdown:**

1. **Pattern purpose**: Match expressions specifying a day of the week with optional time
2. **Dictionary usage**: Translate day names to numbers (monday → 1)
3. **Default behavior**: If time isn't specified, defaults to midnight (0:00)

**Examples with detailed conversion:**

- Input: "every Monday"
  - Matches: `weekday=monday, hour=null, minute=null, ampm=null`
  - Dictionary lookup: `weekday=1`
  - Default values: `hour=0, minute=0`
  - Result: "0 0 0 * * 1" (At midnight on Monday)

- Input: "every Tuesday at 3:45pm"
  - Matches: `weekday=tuesday, hour=3, minute=45, ampm=pm`
  - Dictionary lookups: `weekday=2`
  - Transformation: `hour=3+12=15` (pm and hour<12)
  - Result: "0 45 15 * * 2" (At 3:45 PM on Tuesday)

### Example 3: Monthly Schedule with Special Case

```yaml
- name: nth_weekday_of_month
  pattern: '(?i)(?:each|every)\s+(first|second|third|fourth|fifth|last)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)(?:\s+of\s+(?:the\s+)?month)?'
  variables:
    ordinal: 1
    weekday: 2
  dictionaries:
    ordinal: ordinals
    weekday: weekdays
  format: "0 0 * * %weekday#%ordinal"
  special_cases:
    - condition: "ordinal == 'last'"
      format: "0 0 * * %weekdayL"
```

**Step-by-step breakdown:**

1. **Pattern purpose**: Match expressions for nth occurrence of a weekday in a month
2. **Special case handling**: 
   - Normal ordinals use `#` syntax (3rd Monday = 1#3)
   - "Last" uses the `L` suffix instead (last Friday = 5L)

**Examples with detailed conversion:**

- Input: "every third Wednesday of the month"
  - Matches: `ordinal=third, weekday=wednesday`
  - Dictionary lookups: `ordinal=3, weekday=3`
  - No special case applies
  - Result: "0 0 * * 3#3" (At midnight on the third Wednesday of every month)

- Input: "every last Friday"
  - Matches: `ordinal=last, weekday=friday`
  - Dictionary lookups: `ordinal=L, weekday=5`
  - Special case applies (ordinal == 'last')
  - Result: "0 0 * * 5L" (At midnight on the last Friday of every month)

### Example 4: Specific Day of Month

```yaml
- name: specific_day_of_month
  pattern: '(?i)(?:each|every)\s+(\d+)(?:st|nd|rd|th)?\s+(?:day\s+)?of\s+(?:the\s+)?month(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?'
  variables:
    day: 1
    hour: 2
    minute: 3
    ampm: 4
  dictionaries:
    ampm: time_ampm
  format: "%minute %hour %day * *"
  default_values:
    minute: "0"
    hour: "0"
  transformations:
    hour:
      - condition: "ampm == 'pm' && hour < 12"
        operation: "hour + 12"
      - condition: "ampm == 'am' && hour == 12"
        operation: "0"
```

**Examples:**
- "every 15th of the month" → "0 0 15 * *"
- "each 1st of month at 2:30pm" → "30 14 1 * *"

### Example 5: Last Day of Month

```yaml
- name: last_day_of_month
  pattern: '(?i)(?:each|every|the)\s+last\s+day\s+of\s+(?:the\s+)?month(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?'
  variables:
    hour: 1
    minute: 2
    ampm: 3
  dictionaries:
    ampm: time_ampm
  format: "%minute %hour L * *"
  default_values:
    minute: "0"
    hour: "0"
  transformations:
    hour:
      - condition: "ampm == 'pm' && hour < 12"
        operation: "hour + 12"
      - condition: "ampm == 'am' && hour == 12"
        operation: "0"
```

**Examples:**
- "every last day of month" → "0 0 L * *"
- "the last day of the month at 6pm" → "0 18 L * *"

## Advanced Pattern Techniques

### Capture Groups vs. Non-Capturing Groups

- **Capturing groups** `(pattern)` store matches for variables
- **Non-capturing groups** `(?:pattern)` group alternatives without capturing

```yaml
pattern: '(?i)(?:each|every)\s+(monday|tuesday)'
#        ^     ^            ^
#        |     |            |
#        |     |            Capturing group (becomes a variable)
#        |     Non-capturing group (won't be captured)
#        Case-insensitive flag
```

### Optional Components

Use question marks to make parts optional:

```yaml
pattern: '(?i)(?:each|every)\s+day(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?'
#                                  ^                                     ^
#                                  Optional time specification
```

### Flexible Whitespace

Use `\s+` for flexible whitespace matching:

```yaml
pattern: '(?i)(?:each|every)\s+day\s+at\s+(\d+)'
#                            ^     ^    ^
#                            One or more whitespace characters
```

## Best Practices for Rule Creation

### 1. Rule Ordering

- **Most specific first**: Place more specific rules before general ones
- **Order matters**: The first matching rule is used, so prioritize carefully
- **Test overlapping rules**: Ensure specific cases match correctly before general cases

### 2. Pattern Design

- **Be specific**: Make patterns as specific as possible to avoid false matches
- **Case flexibility**: Use `(?i)` for case-insensitive matching
- **Whitespace tolerance**: Use `\s+` to handle variable spacing between words
- **Optional components**: Use `?` to make parts optional
- **Capture only what's needed**: Use non-capturing groups for alternatives

### 3. Variable Naming

- **Consistency**: Use the same variable names across similar rules
- **Descriptive names**: Choose clear names like `hour`, `minute`, `weekday`
- **Documentation**: Add comments explaining non-obvious variables

### 4. Error Handling and Defaults

- **Reasonable defaults**: Set sensible default values for optional components
- **Validate with transformations**: Use transformations to correct invalid values
- **Handle edge cases**: Account for special cases like 12am/12pm conversion

### 5. Testing Strategies

- **Multiple phrasings**: Test with different ways to express the same schedule
- **Edge cases**: Test boundary conditions and special values
- **Missing components**: Test with optional parts omitted
- **Case variations**: Test with mixed case input

## Common Cron Patterns Reference

| Description | Cron Format | Explanation |
|-------------|------------|-------------|
| Every minute | `* * * * *` | Run every minute of every hour, every day |
| Every hour | `0 * * * *` | Run at minute 0 of every hour, every day |
| Every day at midnight | `0 0 * * *` | Run at 00:00 every day |
| Every Monday | `0 0 * * 1` | Run at midnight every Monday |
| First Monday of month | `0 0 * * 1#1` | Run at midnight on first Monday of every month |
| Last day of month | `0 0 L * *` | Run at midnight on the last day of every month |
| Last Friday of month | `0 0 * * 5L` | Run at midnight on the last Friday of every month |
| Weekday nearest 15th | `0 0 15W * *` | Run at midnight on the weekday closest to the 15th |
| Every 5 minutes | `*/5 * * * *` | Run every 5 minutes |
| Every 2 hours | `0 */2 * * *` | Run every 2 hours, on the hour |

## Troubleshooting Rules

### Common Issues and Solutions

1. **Rule not matching expected input**
   - Check pattern for typos or incorrect capturing groups
   - Verify case sensitivity with (?i) flag
   - Test pattern with regex testing tools

2. **Wrong values in output**
   - Check variable mappings to capture groups
   - Verify dictionary lookups
   - Check transformation conditions and operations

3. **Pattern too restrictive**
   - Add more alternatives with (option1|option2)
   - Make more components optional with ? suffix
   - Use \s+ instead of specific whitespace requirements

4. **Pattern too permissive**
   - Add more specific constraints
   - Check order of rules (more specific rules first)
   - Add boundary markers (^ for start, $ for end)

5. **Unexpected format results**
   - Check format string variable references
   - Verify special case conditions
   - Ensure transformations are processing correctly
