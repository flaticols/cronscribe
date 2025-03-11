package nl

import "github.com/flaticols/cronscribe/pkg/rules"

// Dutch language constants
const (
	TimeVM = "vm" // voormiddag (morning)
	TimeNM = "nm" // namiddag (afternoon)

	OrdinalLaatste = "laatste" // last
)

// Dutch specific dictionary values
var DutchWeekdayValues = map[string]string{
	"zondag":    "0",
	"maandag":   "1",
	"dinsdag":   "2",
	"woensdag":  "3",
	"donderdag": "4",
	"vrijdag":   "5",
	"zaterdag":  "6",
}

var DutchOrdinalValues = map[string]string{
	"eerste":       "1",
	"tweede":       "2",
	"derde":        "3",
	"vierde":       "4",
	"vijfde":       "5",
	OrdinalLaatste: "L",
}

var DutchTimeAmPmValues = map[string]string{
	TimeVM: rules.TimeAm,
	TimeNM: rules.TimePm,
}

var DutchMonthValues = map[string]string{
	"januari":   "1",
	"februari":  "2",
	"maart":     "3",
	"april":     "4",
	"mei":       "5",
	"juni":      "6",
	"juli":      "7",
	"augustus":  "8",
	"september": "9",
	"oktober":   "10",
	"november":  "11",
	"december":  "12",
}

// RuleSet contains Dutch language rules and dictionaries
var RuleSet = &rules.RuleSet{
	Language: "nl",
	Dictionaries: map[string]rules.Dictionary{
		rules.DictWeekdays: DutchWeekdayValues,
		rules.DictOrdinals: DutchOrdinalValues,
		rules.DictTimeAmPm: DutchTimeAmPmValues,
		rules.DictMonths:   DutchMonthValues,
	},
	Rules: []rules.Rule{
		{
			Name:    "nth_weekday_of_month",
			Pattern: `(?i)(?:elke|iedere)\s+(eerste|tweede|derde|vierde|vijfde|laatste)\s+(maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag)(?:\s+van\s+de\s+maand)?`,
			Variables: map[string]int{
				rules.VarOrdinal: 1,
				rules.VarWeekday: 2,
			},
			Dictionaries: map[string]string{
				rules.VarOrdinal: rules.DictOrdinals,
				rules.VarWeekday: rules.DictWeekdays,
			},
			Format: "0 0 * * %weekday#%ordinal",
			SpecialCases: []rules.SpecialCase{
				{
					Condition: rules.VarOrdinal + " == '" + OrdinalLaatste + "'",
					Format:    "0 0 * * %weekdayL",
				},
			},
		},
		{
			Name:    "weekly_day_at_time",
			Pattern: `(?i)(?:elke|iedere)\s+(maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag)(?:\s+om\s+(\d+)(?::(\d+))?\s*(vm|nm)?)?`,
			Variables: map[string]int{
				rules.VarWeekday: 1,
				rules.VarHour:    2,
				rules.VarMinute:  3,
				rules.VarAmPm:    4,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday: rules.DictWeekdays,
				rules.VarAmPm:    rules.DictTimeAmPm,
			},
			Format: "0 %minute %hour * * %weekday",
			DefaultValues: map[string]string{
				rules.VarMinute: "0",
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarHour: {
					{
						Condition: rules.VarAmPm + " == '" + rules.TimePm + "' && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: rules.VarAmPm + " == '" + rules.TimeAm + "' && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "daily_at_time",
			Pattern: `(?i)(?:elke|iedere)\s+dag\s+om\s+(\d+)(?::(\d+))?\s*(vm|nm)?`,
			Variables: map[string]int{
				rules.VarHour:   1,
				rules.VarMinute: 2,
				rules.VarAmPm:   3,
			},
			Dictionaries: map[string]string{
				rules.VarAmPm: rules.DictTimeAmPm,
			},
			Format: "%minute %hour * * *",
			DefaultValues: map[string]string{
				rules.VarMinute: "0",
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarHour: {
					{
						Condition: rules.VarAmPm + " == '" + rules.TimePm + "' && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: rules.VarAmPm + " == '" + rules.TimeAm + "' && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "hourly",
			Pattern: `(?i)(?:elk|ieder)\s+uur`,
			Format:  "0 * * * *",
		},
		{
			Name:    "every_n_minutes",
			Pattern: `(?i)(?:elke|iedere)\s+(\d+)\s+min(?:u(?:ut|ten))?`,
			Variables: map[string]int{
				rules.VarMinutes: 1,
			},
			Format: "*/%minutes * * * *",
		},
		{
			Name:    "every_n_hours",
			Pattern: `(?i)(?:elke|iedere)\s+(\d+)\s+(?:uur|uren)`,
			Variables: map[string]int{
				rules.VarHours: 1,
			},
			Format: "0 */%hours * * *",
		},
		{
			Name:    "specific_day_of_month",
			Pattern: `(?i)(?:elke|iedere)\s+(\d+)(?:e|de|ste)?\s+(?:dag\s+)?van\s+de\s+maand(?:\s+om\s+(\d+)(?::(\d+))?\s*(vm|nm)?)?`,
			Variables: map[string]int{
				rules.VarDay:    1,
				rules.VarHour:   2,
				rules.VarMinute: 3,
				rules.VarAmPm:   4,
			},
			Dictionaries: map[string]string{
				rules.VarAmPm: rules.DictTimeAmPm,
			},
			Format: "%minute %hour %day * *",
			DefaultValues: map[string]string{
				rules.VarMinute: "0",
				rules.VarHour:   "0",
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarHour: {
					{
						Condition: rules.VarAmPm + " == '" + rules.TimePm + "' && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: rules.VarAmPm + " == '" + rules.TimeAm + "' && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "specific_month_day",
			Pattern: `(?i)(?:elke|iedere)\s+(januari|februari|maart|april|mei|juni|juli|augustus|september|oktober|november|december)\s+(\d+)(?:e|de|ste)?(?:\s+om\s+(\d+)(?::(\d+))?\s*(vm|nm)?)?`,
			Variables: map[string]int{
				rules.VarMonth:  1,
				rules.VarDay:    2,
				rules.VarHour:   3,
				rules.VarMinute: 4,
				rules.VarAmPm:   5,
			},
			Dictionaries: map[string]string{
				rules.VarMonth: rules.DictMonths,
				rules.VarAmPm:  rules.DictTimeAmPm,
			},
			Format: "%minute %hour %day %month *",
			DefaultValues: map[string]string{
				rules.VarMinute: "0",
				rules.VarHour:   "0",
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarHour: {
					{
						Condition: rules.VarAmPm + " == '" + rules.TimePm + "' && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: rules.VarAmPm + " == '" + rules.TimeAm + "' && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "last_day_of_month",
			Pattern: `(?i)(?:elke|iedere|de)\s+laatste\s+dag\s+van\s+de\s+maand(?:\s+om\s+(\d+)(?::(\d+))?\s*(vm|nm)?)?`,
			Variables: map[string]int{
				rules.VarHour:   1,
				rules.VarMinute: 2,
				rules.VarAmPm:   3,
			},
			Dictionaries: map[string]string{
				rules.VarAmPm: rules.DictTimeAmPm,
			},
			Format: "%minute %hour L * *",
			DefaultValues: map[string]string{
				rules.VarMinute: "0",
				rules.VarHour:   "0",
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarHour: {
					{
						Condition: rules.VarAmPm + " == '" + rules.TimePm + "' && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: rules.VarAmPm + " == '" + rules.TimeAm + "' && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "weekday_nearest_day",
			Pattern: `(?i)(?:elke|iedere|de)\s+(maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag)\s+(?:het\s+)?dichtstbij\s+(?:de\s+)?(\d+)(?:e|de|ste)?(?:\s+om\s+(\d+)(?::(\d+))?\s*(vm|nm)?)?`,
			Variables: map[string]int{
				rules.VarWeekday: 1,
				rules.VarDay:     2,
				rules.VarHour:    3,
				rules.VarMinute:  4,
				rules.VarAmPm:    5,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday: rules.DictWeekdays,
				rules.VarAmPm:    rules.DictTimeAmPm,
			},
			Format: "%minute %hour %dayW %month %weekday",
			DefaultValues: map[string]string{
				rules.VarMinute: "0",
				rules.VarHour:   "0",
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarHour: {
					{
						Condition: rules.VarAmPm + " == '" + rules.TimePm + "' && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: rules.VarAmPm + " == '" + rules.TimeAm + "' && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
	},
}

func init() {
	// Compile the patterns for all Dutch rules
	for i := range RuleSet.Rules {
		_ = RuleSet.Rules[i].CompilePattern()
	}
}
