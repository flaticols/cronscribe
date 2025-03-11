package ru

import "github.com/flaticols/cronscribe/pkg/rules"

// Russian language constants
const (
	// Time of day periods
	TimeUtra    = "утра"   // morning
	TimeDnya    = "дня"    // afternoon
	TimeVechera = "вечера" // evening
	TimeNochi   = "ночи"   // night

	// Ordinals
	OrdinalPervyj      = "первый"    // first (masculine)
	OrdinalPervaya     = "первая"    // first (feminine)
	OrdinalPervoe      = "первое"    // first (neuter)
	OrdinalVtoroj      = "второй"    // second (masculine)
	OrdinalVtoraya     = "вторая"    // second (feminine)
	OrdinalVtoroe      = "второе"    // second (neuter)
	OrdinalTretij      = "третий"    // third (masculine)
	OrdinalTretya      = "третья"    // third (feminine)
	OrdinalTrete       = "третье"    // third (neuter)
	OrdinalChetvertyj  = "четвертый" // fourth (masculine)
	OrdinalChetvertaya = "четвертая" // fourth (feminine)
	OrdinalChetvertoe  = "четвертое" // fourth (neuter)
	OrdinalPyatyj      = "пятый"     // fifth (masculine)
	OrdinalPyataya     = "пятая"     // fifth (feminine)
	OrdinalPyatoe      = "пятое"     // fifth (neuter)
	OrdinalPoslednij   = "последний" // last (masculine)
	OrdinalPoslednyaya = "последняя" // last (feminine)
	OrdinalPoslednee   = "последнее" // last (neuter)

	// Weekdays with case variations
	WeekdaySreda    = "среда"   // Wednesday (nominative)
	WeekdaySredu    = "среду"   // Wednesday (accusative)
	WeekdayPyatnica = "пятница" // Friday (nominative)
	WeekdayPyatnicu = "пятницу" // Friday (accusative)
	WeekdaySubbota  = "суббота" // Saturday (nominative)
	WeekdaySubbotu  = "субботу" // Saturday (accusative)
)

// Russian-specific dictionary values
var RussianWeekdayValues = map[string]string{
	"воскресенье":   "0",
	"понедельник":   "1",
	"вторник":       "2",
	WeekdaySreda:    "3",
	"четверг":       "4",
	WeekdayPyatnica: "5",
	WeekdaySubbota:  "6",
}

var RussianOrdinalValues = map[string]string{
	OrdinalPervyj:     "1",
	OrdinalVtoroj:     "2",
	OrdinalTretij:     "3",
	OrdinalChetvertyj: "4",
	OrdinalPyatyj:     "5",
	OrdinalPoslednij:  "L",
}

var RussianTimeAmPmValues = map[string]string{
	TimeUtra:    rules.TimeAm,
	TimeDnya:    rules.TimePm,
	TimeVechera: rules.TimePm,
	TimeNochi:   rules.TimeAm,
}

var RussianMonthValues = map[string]string{
	"января":   "1",
	"февраля":  "2",
	"марта":    "3",
	"апреля":   "4",
	"мая":      "5",
	"июня":     "6",
	"июля":     "7",
	"августа":  "8",
	"сентября": "9",
	"октября":  "10",
	"ноября":   "11",
	"декабря":  "12",
}

// RuleSet contains Russian language rules and dictionaries
var RuleSet = &rules.RuleSet{
	Language: "ru",
	Dictionaries: map[string]rules.Dictionary{
		rules.DictWeekdays: RussianWeekdayValues,
		rules.DictOrdinals: RussianOrdinalValues,
		rules.DictTimeAmPm: RussianTimeAmPmValues,
		rules.DictMonths:   RussianMonthValues,
	},
	Rules: []rules.Rule{
		{
			Name:    "nth_weekday_of_month",
			Pattern: `(?i)кажд(?:ый|ая|ое)\s+(перв(?:ый|ая|ое)|втор(?:ой|ая|ое)|трет(?:ий|ья|ье)|четверт(?:ый|ая|ое)|пят(?:ый|ая|ое)|последн(?:ий|яя|ее))\s+(понедельник|вторник|сред[ау]|четверг|пятниц[ау]|суббот[ау]|воскресенье)(?:\s+(?:месяца|в месяце))?`,
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
					Condition: rules.VarOrdinal + " == '" + OrdinalPoslednij + "'",
					Format:    "0 0 * * %weekdayL",
				},
			},
			Transformations: map[string][]rules.Transformation{
				rules.VarOrdinal: {
					{
						Condition: rules.VarOrdinal + " == '" + OrdinalPervaya + "' || " + rules.VarOrdinal + " == '" + OrdinalPervoe + "'",
						Operation: "'" + OrdinalPervyj + "'",
					},
					{
						Condition: rules.VarOrdinal + " == '" + OrdinalVtoraya + "' || " + rules.VarOrdinal + " == '" + OrdinalVtoroe + "'",
						Operation: "'" + OrdinalVtoroj + "'",
					},
					{
						Condition: rules.VarOrdinal + " == '" + OrdinalTretya + "' || " + rules.VarOrdinal + " == '" + OrdinalTrete + "'",
						Operation: "'" + OrdinalTretij + "'",
					},
					{
						Condition: rules.VarOrdinal + " == '" + OrdinalChetvertaya + "' || " + rules.VarOrdinal + " == '" + OrdinalChetvertoe + "'",
						Operation: "'" + OrdinalChetvertyj + "'",
					},
					{
						Condition: rules.VarOrdinal + " == '" + OrdinalPyataya + "' || " + rules.VarOrdinal + " == '" + OrdinalPyatoe + "'",
						Operation: "'" + OrdinalPyatyj + "'",
					},
					{
						Condition: rules.VarOrdinal + " == '" + OrdinalPoslednyaya + "' || " + rules.VarOrdinal + " == '" + OrdinalPoslednee + "'",
						Operation: "'" + OrdinalPoslednij + "'",
					},
				},
				rules.VarWeekday: {
					{
						Condition: rules.VarWeekday + " == '" + WeekdaySredu + "'",
						Operation: "'" + WeekdaySreda + "'",
					},
					{
						Condition: rules.VarWeekday + " == '" + WeekdayPyatnicu + "'",
						Operation: "'" + WeekdayPyatnica + "'",
					},
					{
						Condition: rules.VarWeekday + " == '" + WeekdaySubbotu + "'",
						Operation: "'" + WeekdaySubbota + "'",
					},
				},
			},
		},
		{
			Name:    "weekly_day_at_time",
			Pattern: `(?i)кажд(?:ый|ую)\s+(понедельник|вторник|сред[ау]|четверг|пятниц[ау]|суббот[ау]|воскресенье)(?:\s+в\s+(\d+)(?::(\d+))?\s*(?:час(?:ов|а)?)?(?:\s+(утра|дня|вечера|ночи))?)?`,
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
				rules.VarWeekday: {
					{
						Condition: rules.VarWeekday + " == '" + WeekdaySredu + "'",
						Operation: "'" + WeekdaySreda + "'",
					},
					{
						Condition: rules.VarWeekday + " == '" + WeekdayPyatnicu + "'",
						Operation: "'" + WeekdayPyatnica + "'",
					},
					{
						Condition: rules.VarWeekday + " == '" + WeekdaySubbotu + "'",
						Operation: "'" + WeekdaySubbota + "'",
					},
				},
				rules.VarHour: {
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeDnya + "' || " + rules.VarAmPm + " == '" + TimeVechera + "') && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeUtra + "' || " + rules.VarAmPm + " == '" + TimeNochi + "') && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "daily_at_time",
			Pattern: `(?i)кажд(?:ый|ую)\s+день\s+в\s+(\d+)(?::(\d+))?\s*(?:час(?:ов|а)?)?(?:\s+(утра|дня|вечера|ночи))?`,
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
						Condition: "(" + rules.VarAmPm + " == '" + TimeDnya + "' || " + rules.VarAmPm + " == '" + TimeVechera + "') && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeUtra + "' || " + rules.VarAmPm + " == '" + TimeNochi + "') && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "hourly",
			Pattern: `(?i)кажд(?:ый|ую)\s+час`,
			Format:  "0 * * * *",
		},
		{
			Name:    "every_n_minutes",
			Pattern: `(?i)кажд(?:ые|ую)\s+(\d+)\s+минут(?:ы|у)?`,
			Variables: map[string]int{
				rules.VarMinutes: 1,
			},
			Format: "*/%minutes * * * *",
		},
		{
			Name:    "every_n_hours",
			Pattern: `(?i)кажд(?:ые|ую)\s+(\d+)\s+час(?:а|ов)?`,
			Variables: map[string]int{
				rules.VarHours: 1,
			},
			Format: "0 */%hours * * *",
		},
		{
			Name:    "specific_day_of_month",
			Pattern: `(?i)кажд(?:ое|ого)\s+(\d+)(?:-е|-го)?\s+(?:число|дня)?\s+(?:месяца|в месяце)?(?:\s+в\s+(\d+)(?::(\d+))?\s*(?:час(?:ов|а)?)?(?:\s+(утра|дня|вечера|ночи))?)?`,
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
						Condition: "(" + rules.VarAmPm + " == '" + TimeDnya + "' || " + rules.VarAmPm + " == '" + TimeVechera + "') && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeUtra + "' || " + rules.VarAmPm + " == '" + TimeNochi + "') && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "specific_month_day",
			Pattern: `(?i)кажд(?:ого|ое)\s+(\d+)(?:-е|-го)?\s+(января|февраля|марта|апреля|мая|июня|июля|августа|сентября|октября|ноября|декабря)(?:\s+в\s+(\d+)(?::(\d+))?\s*(?:час(?:ов|а)?)?(?:\s+(утра|дня|вечера|ночи))?)?`,
			Variables: map[string]int{
				rules.VarDay:    1,
				rules.VarMonth:  2,
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
						Condition: "(" + rules.VarAmPm + " == '" + TimeDnya + "' || " + rules.VarAmPm + " == '" + TimeVechera + "') && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeUtra + "' || " + rules.VarAmPm + " == '" + TimeNochi + "') && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "last_day_of_month",
			Pattern: `(?i)(?:каждый|в)\s+последни(?:й|е)\s+день\s+(?:месяца|в месяце)(?:\s+в\s+(\d+)(?::(\d+))?\s*(?:час(?:ов|а)?)?(?:\s+(утра|дня|вечера|ночи))?)?`,
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
						Condition: "(" + rules.VarAmPm + " == '" + TimeDnya + "' || " + rules.VarAmPm + " == '" + TimeVechera + "') && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeUtra + "' || " + rules.VarAmPm + " == '" + TimeNochi + "') && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
		{
			Name:    "weekday_nearest_day",
			Pattern: `(?i)(?:каждый|в)\s+(понедельник|вторник|сред[ау]|четверг|пятниц[ау]|суббот[ау]|воскресенье)\s+ближайший\s+к\s+(\d+)(?:-му|-ому)?(?:\s+числу)?(?:\s+в\s+(\d+)(?::(\d+))?\s*(?:час(?:ов|а)?)?(?:\s+(утра|дня|вечера|ночи))?)?`,
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
				rules.VarWeekday: {
					{
						Condition: rules.VarWeekday + " == '" + WeekdaySredu + "'",
						Operation: "'" + WeekdaySreda + "'",
					},
					{
						Condition: rules.VarWeekday + " == '" + WeekdayPyatnicu + "'",
						Operation: "'" + WeekdayPyatnica + "'",
					},
					{
						Condition: rules.VarWeekday + " == '" + WeekdaySubbotu + "'",
						Operation: "'" + WeekdaySubbota + "'",
					},
				},
				rules.VarHour: {
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeDnya + "' || " + rules.VarAmPm + " == '" + TimeVechera + "') && " + rules.VarHour + " < 12",
						Operation: rules.VarHour + " + 12",
					},
					{
						Condition: "(" + rules.VarAmPm + " == '" + TimeUtra + "' || " + rules.VarAmPm + " == '" + TimeNochi + "') && " + rules.VarHour + " == 12",
						Operation: "0",
					},
				},
			},
		},
	},
}

func init() {
	// Compile the patterns for all Russian rules
	for i := range RuleSet.Rules {
		_ = RuleSet.Rules[i].CompilePattern()
	}
}
