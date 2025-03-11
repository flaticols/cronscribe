package rules

// Dictionary key constants
const (
	// Dictionary names
	DictWeekdays = "weekdays"
	DictOrdinals = "ordinals"
	DictTimeAmPm = "time_ampm"
	DictMonths   = "months"

	// Variable names
	VarHour    = "hour"
	VarMinute  = "minute"
	VarDay     = "day"
	VarWeekday = "weekday"
	VarMonth   = "month"
	VarAmPm    = "ampm"
	VarOrdinal = "ordinal"
	VarMinutes = "minutes"
	VarHours   = "hours"

	// Time period values
	TimeAm = "am"
	TimePm = "pm"

	// Special values
	OrdinalLast = "last"
)

// RuleSet represents a complete set of rules and dictionaries for a language
type RuleSet struct {
	Language     string                `yaml:"language"`
	Rules        []Rule                `yaml:"rules"`
	Dictionaries map[string]Dictionary `yaml:"dictionaries"`
}

// Dictionary represents a mapping of keys to values
type Dictionary map[string]string

// WeekdayValues contains standard weekday mappings
var WeekdayValues = map[string]string{
	"sunday":    "0",
	"monday":    "1",
	"tuesday":   "2",
	"wednesday": "3",
	"thursday":  "4",
	"friday":    "5",
	"saturday":  "6",
}

// OrdinalValues contains standard ordinal mappings
var OrdinalValues = map[string]string{
	"first":     "1",
	"second":    "2",
	"third":     "3",
	"fourth":    "4",
	"fifth":     "5",
	OrdinalLast: "L",
}

// TimeAmPmValues contains standard time format mappings
var TimeAmPmValues = map[string]string{
	TimeAm: TimeAm,
	TimePm: TimePm,
}

// MonthValues contains standard month name mappings
var MonthValues = map[string]string{
	"january":   "1",
	"february":  "2",
	"march":     "3",
	"april":     "4",
	"may":       "5",
	"june":      "6",
	"july":      "7",
	"august":    "8",
	"september": "9",
	"october":   "10",
	"november":  "11",
	"december":  "12",
}
