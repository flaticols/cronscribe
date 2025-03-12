package rules

// Dictionary key constants
const (
	// Dictionary names
	DictWeekdays     = "weekdays"
	DictOrdinals     = "ordinals"
	DictTimeAmPm     = "time_ampm"
	DictMonths       = "months"
	DictTimePeriods  = "time_periods"  // For morning, afternoon, etc.
	DictTimeSpecific = "time_specific" // For noon, midnight, etc.

	// Variable names
	VarHour       = "hour"
	VarMinute     = "minute"
	VarDay        = "day"
	VarWeekday    = "weekday"
	VarMonth      = "month"
	VarAmPm       = "ampm"
	VarOrdinal    = "ordinal"
	VarMinutes    = "minutes"
	VarHours      = "hours"
	VarTimePeriod = "timeperiod"   // For time periods (morning, afternoon, etc.)
	VarTimePoint  = "timepoint"    // For specific time points (noon, midnight)

	// Time period values
	TimeAm = "am"
	TimePm = "pm"

	// Special values
	OrdinalLast = "last"
)

// RuleSet represents a complete set of rules and dictionaries for a language
type RuleSet struct {
	Language        string                `yaml:"language"`
	Rules           []Rule                `yaml:"rules"`
	Dictionaries    map[string]Dictionary `yaml:"dictionaries"`
	SpecialTestCases map[string]string    `yaml:"special_test_cases,omitempty"`
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

// TimeSpecificValues contains standard time point mappings (English)
var TimeSpecificValues = map[string]string{
	"midnight": "0",
	"noon":     "12",
}

// TimePeriodValues contains standard time period mappings (English)
var TimePeriodValues = map[string]string{
	"morning":   "5-11",  // 5:00 AM to 11:59 AM
	"afternoon": "12-17", // 12:00 PM to 5:59 PM
	"evening":   "18-21", // 6:00 PM to 9:59 PM
	"night":     "22-4",  // 10:00 PM to 4:59 AM
}

// TimeSpecificValuesRU contains time point mappings for Russian
var TimeSpecificValuesRU = map[string]string{
	"полночь": "0",
	"полдень": "12",
}

// TimePeriodValuesRU contains time period mappings for Russian
var TimePeriodValuesRU = map[string]string{
	"утро":    "5-11",  // 5:00 AM to 11:59 AM
	"день":    "12-17", // 12:00 PM to 5:59 PM
	"вечер":   "18-21", // 6:00 PM to 9:59 PM
	"ночь":    "22-4",  // 10:00 PM to 4:59 AM
	"утром":   "5-11",  // Instrumental case for "in the morning"
	"днем":    "12-17", // Instrumental case for "in the afternoon"
	"вечером": "18-21", // Instrumental case for "in the evening"
	"ночью":   "22-4",  // Instrumental case for "at night"
}

// TimeSpecificValuesNL contains time point mappings for Dutch
var TimeSpecificValuesNL = map[string]string{
	"middernacht": "0",
	"middag":      "12",
}

// TimePeriodValuesNL contains time period mappings for Dutch
var TimePeriodValuesNL = map[string]string{
	"ochtend":  "5-11",  // 5:00 AM to 11:59 AM
	"namiddag": "12-17", // 12:00 PM to 5:59 PM
	"avond":    "18-21", // 6:00 PM to 9:59 PM
	"nacht":    "22-4",  // 10:00 PM to 4:59 AM
}
