package en

import "github.com/flaticols/cronscribe/pkg/rules"

// RuleSet contains English language rules and dictionaries
var RuleSet = &rules.RuleSet{
	Language: "en",
	Dictionaries: map[string]rules.Dictionary{
		rules.DictWeekdays:     rules.WeekdayValues,
		rules.DictOrdinals:     rules.OrdinalValues,
		rules.DictTimeAmPm:     rules.TimeAmPmValues,
		rules.DictMonths:       rules.MonthValues,
		rules.DictTimePeriods:  rules.TimePeriodValues,
		rules.DictTimeSpecific: rules.TimeSpecificValues,
	},
	Rules: []rules.Rule{
		{
			Name:    "nth_weekday_of_month",
			Pattern: `(?i)(?:each|every)\s+(first|second|third|fourth|fifth|last)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)(?:\s+of\s+(?:the\s+)?month)?`,
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
					Condition: rules.VarOrdinal + " == '" + rules.OrdinalLast + "'",
					Format:    "0 0 * * %weekdayL",
				},
			},
		},
		{
			Name:    "weekly_day_at_time",
			Pattern: `(?i)(?:each|every)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?`,
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
			Pattern: `(?i)(?:each|every)\s+day\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?`,
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
			Name:    "daily_at_24h_time",
			Pattern: `(?i)(?:each|every)\s+day\s+at\s+(\d{1,2}):(\d{2})(?:\s*(?:hours?|hrs?))?`,
			Variables: map[string]int{
				rules.VarHour:   1,
				rules.VarMinute: 2,
			},
			Format: "%minute %hour * * *",
		},
		{
			Name:    "daily_at_specific_time",
			Pattern: `(?i)(?:each|every)\s+day\s+at\s+(noon|midnight)`,
			Variables: map[string]int{
				rules.VarTimePoint: 1,
			},
			Dictionaries: map[string]string{
				rules.VarTimePoint: rules.DictTimeSpecific,
			},
			Format: "0 %timepoint * * *",
		},
		{
			Name:    "daily_in_time_period",
			Pattern: `(?i)(?:each|every)\s+day\s+(?:in|during)\s+(?:the\s+)?(morning|afternoon|evening|night)`,
			Variables: map[string]int{
				rules.VarTimePeriod: 1,
			},
			Dictionaries: map[string]string{
				rules.VarTimePeriod: rules.DictTimePeriods,
			},
			Format: "0 %timeperiod * * *",
		},
		{
			Name:    "weekly_day_at_specific_time",
			Pattern: `(?i)(?:each|every)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+at\s+(noon|midnight)`,
			Variables: map[string]int{
				rules.VarWeekday:   1,
				rules.VarTimePoint: 2,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday:   rules.DictWeekdays,
				rules.VarTimePoint: rules.DictTimeSpecific,
			},
			Format: "0 %timepoint * * %weekday",
		},
		{
			Name:    "weekly_day_in_time_period",
			Pattern: `(?i)(?:each|every)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+(?:in|during)\s+(?:the\s+)?(morning|afternoon|evening|night)`,
			Variables: map[string]int{
				rules.VarWeekday:    1,
				rules.VarTimePeriod: 2,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday:    rules.DictWeekdays,
				rules.VarTimePeriod: rules.DictTimePeriods,
			},
			Format: "0 %timeperiod * * %weekday",
		},
		{
			Name:    "weekly_day_at_24h_time",
			Pattern: `(?i)(?:each|every)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+at\s+(\d{1,2}):(\d{2})(?:\s*(?:hours?|hrs?))?`,
			Variables: map[string]int{
				rules.VarWeekday: 1,
				rules.VarHour:    2,
				rules.VarMinute:  3,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday: rules.DictWeekdays,
			},
			Format: "%minute %hour * * %weekday",
		},
		{
			Name:    "hourly",
			Pattern: `(?i)(?:each|every)\s+hour`,
			Format:  "0 * * * *",
		},
		{
			Name:    "every_n_minutes",
			Pattern: `(?i)(?:each|every)\s+(\d+)\s+minutes?`,
			Variables: map[string]int{
				rules.VarMinutes: 1,
			},
			Format: "*/%minutes * * * *",
		},
		{
			Name:    "every_n_hours",
			Pattern: `(?i)(?:each|every)\s+(\d+)\s+hours?`,
			Variables: map[string]int{
				rules.VarHours: 1,
			},
			Format: "0 */%hours * * *",
		},
		{
			Name:    "specific_day_of_month",
			Pattern: `(?i)(?:each|every)\s+(\d+)(?:st|nd|rd|th)?\s+(?:day\s+)?of\s+(?:the\s+)?month(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?`,
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
			Name:    "specific_day_of_month_at_24h_time",
			Pattern: `(?i)(?:each|every)\s+(\d+)(?:st|nd|rd|th)?\s+(?:day\s+)?of\s+(?:the\s+)?month\s+at\s+(\d{1,2}):(\d{2})(?:\s*(?:hours?|hrs?))?`,
			Variables: map[string]int{
				rules.VarDay:    1,
				rules.VarHour:   2,
				rules.VarMinute: 3,
			},
			Format: "%minute %hour %day * *",
		},
		{
			Name:    "specific_day_of_month_at_specific_time",
			Pattern: `(?i)(?:each|every)\s+(\d+)(?:st|nd|rd|th)?\s+(?:day\s+)?of\s+(?:the\s+)?month\s+at\s+(noon|midnight)`,
			Variables: map[string]int{
				rules.VarDay:       1,
				rules.VarTimePoint: 2,
			},
			Dictionaries: map[string]string{
				rules.VarTimePoint: rules.DictTimeSpecific,
			},
			Format: "0 %timepoint %day * *",
		},
		{
			Name:    "specific_day_of_month_in_time_period",
			Pattern: `(?i)(?:each|every)\s+(\d+)(?:st|nd|rd|th)?\s+(?:day\s+)?of\s+(?:the\s+)?month\s+(?:in|during)\s+(?:the\s+)?(morning|afternoon|evening|night)`,
			Variables: map[string]int{
				rules.VarDay:        1,
				rules.VarTimePeriod: 2,
			},
			Dictionaries: map[string]string{
				rules.VarTimePeriod: rules.DictTimePeriods,
			},
			Format: "0 %timeperiod %day * *",
		},
		{
			Name:    "specific_month_day",
			Pattern: `(?i)(?:each|every)\s+(january|february|march|april|may|june|july|august|september|october|november|december)\s+(\d+)(?:st|nd|rd|th)?(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?`,
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
			Name:    "specific_month_day_at_24h_time",
			Pattern: `(?i)(?:each|every)\s+(january|february|march|april|may|june|july|august|september|october|november|december)\s+(\d+)(?:st|nd|rd|th)?\s+at\s+(\d{1,2}):(\d{2})(?:\s*(?:hours?|hrs?))?`,
			Variables: map[string]int{
				rules.VarMonth:  1,
				rules.VarDay:    2,
				rules.VarHour:   3,
				rules.VarMinute: 4,
			},
			Dictionaries: map[string]string{
				rules.VarMonth: rules.DictMonths,
			},
			Format: "%minute %hour %day %month *",
		},
		{
			Name:    "specific_month_day_at_specific_time",
			Pattern: `(?i)(?:each|every)\s+(january|february|march|april|may|june|july|august|september|october|november|december)\s+(\d+)(?:st|nd|rd|th)?\s+at\s+(noon|midnight)`,
			Variables: map[string]int{
				rules.VarMonth:     1,
				rules.VarDay:       2,
				rules.VarTimePoint: 3,
			},
			Dictionaries: map[string]string{
				rules.VarMonth:     rules.DictMonths,
				rules.VarTimePoint: rules.DictTimeSpecific,
			},
			Format: "0 %timepoint %day %month *",
		},
		{
			Name:    "specific_month_day_in_time_period",
			Pattern: `(?i)(?:each|every)\s+(january|february|march|april|may|june|july|august|september|october|november|december)\s+(\d+)(?:st|nd|rd|th)?\s+(?:in|during)\s+(?:the\s+)?(morning|afternoon|evening|night)`,
			Variables: map[string]int{
				rules.VarMonth:      1,
				rules.VarDay:        2,
				rules.VarTimePeriod: 3,
			},
			Dictionaries: map[string]string{
				rules.VarMonth:      rules.DictMonths,
				rules.VarTimePeriod: rules.DictTimePeriods,
			},
			Format: "0 %timeperiod %day %month *",
		},
		{
			Name:    "last_day_of_month",
			Pattern: `(?i)(?:each|every|the)\s+last\s+day\s+of\s+(?:the\s+)?month(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?`,
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
			Name:    "last_day_of_month_at_24h_time",
			Pattern: `(?i)(?:each|every|the)\s+last\s+day\s+of\s+(?:the\s+)?month\s+at\s+(\d{1,2}):(\d{2})(?:\s*(?:hours?|hrs?))?`,
			Variables: map[string]int{
				rules.VarHour:   1,
				rules.VarMinute: 2,
			},
			Format: "%minute %hour L * *",
		},
		{
			Name:    "last_day_of_month_at_specific_time",
			Pattern: `(?i)(?:each|every|the)\s+last\s+day\s+of\s+(?:the\s+)?month\s+at\s+(noon|midnight)`,
			Variables: map[string]int{
				rules.VarTimePoint: 1,
			},
			Dictionaries: map[string]string{
				rules.VarTimePoint: rules.DictTimeSpecific,
			},
			Format: "0 %timepoint L * *",
		},
		{
			Name:    "last_day_of_month_in_time_period",
			Pattern: `(?i)(?:each|every|the)\s+last\s+day\s+of\s+(?:the\s+)?month\s+(?:in|during)\s+(?:the\s+)?(morning|afternoon|evening|night)`,
			Variables: map[string]int{
				rules.VarTimePeriod: 1,
			},
			Dictionaries: map[string]string{
				rules.VarTimePeriod: rules.DictTimePeriods,
			},
			Format: "0 %timeperiod L * *",
		},
		{
			Name:    "weekday_nearest_day",
			Pattern: `(?i)(?:each|every|the)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+nearest\s+(?:to\s+)?(?:the\s+)?(\d+)(?:st|nd|rd|th)?(?:\s+at\s+(\d+)(?::(\d+))?\s*(am|pm)?)?`,
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
		{
			Name:    "weekday_nearest_day_at_24h_time",
			Pattern: `(?i)(?:each|every|the)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+nearest\s+(?:to\s+)?(?:the\s+)?(\d+)(?:st|nd|rd|th)?\s+at\s+(\d{1,2}):(\d{2})(?:\s*(?:hours?|hrs?))?`,
			Variables: map[string]int{
				rules.VarWeekday: 1,
				rules.VarDay:     2,
				rules.VarHour:    3,
				rules.VarMinute:  4,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday: rules.DictWeekdays,
			},
			Format: "%minute %hour %dayW %month %weekday",
		},
		{
			Name:    "weekday_nearest_day_at_specific_time",
			Pattern: `(?i)(?:each|every|the)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+nearest\s+(?:to\s+)?(?:the\s+)?(\d+)(?:st|nd|rd|th)?\s+at\s+(noon|midnight)`,
			Variables: map[string]int{
				rules.VarWeekday:   1,
				rules.VarDay:       2,
				rules.VarTimePoint: 3,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday:   rules.DictWeekdays,
				rules.VarTimePoint: rules.DictTimeSpecific,
			},
			Format: "0 %timepoint %dayW %month %weekday",
		},
		{
			Name:    "weekday_nearest_day_in_time_period",
			Pattern: `(?i)(?:each|every|the)\s+(monday|tuesday|wednesday|thursday|friday|saturday|sunday)\s+nearest\s+(?:to\s+)?(?:the\s+)?(\d+)(?:st|nd|rd|th)?\s+(?:in|during)\s+(?:the\s+)?(morning|afternoon|evening|night)`,
			Variables: map[string]int{
				rules.VarWeekday:    1,
				rules.VarDay:        2,
				rules.VarTimePeriod: 3,
			},
			Dictionaries: map[string]string{
				rules.VarWeekday:    rules.DictWeekdays,
				rules.VarTimePeriod: rules.DictTimePeriods,
			},
			Format: "0 %timeperiod %dayW %month %weekday",
		},
	},
}

func init() {
	// Compile the patterns for all English rules
	for i := range RuleSet.Rules {
		_ = RuleSet.Rules[i].CompilePattern()
	}
}
