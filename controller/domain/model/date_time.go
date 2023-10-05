package model

import "time"

const (
	DateTimeFormat  = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
	SlashDateFormat = "2006/01/02"
)

type DateTime time.Time

func DateTimeNow() DateTime {
	return DateTime(time.Now())
}

func DateTimeFromString(s string, format string) (DateTime, error) {
	t, err := time.ParseInLocation(format, s, time.Local)
	if err != nil {
		return DateTime{}, err
	}

	return DateTime(t), nil
}

func (d DateTime) ToTime() time.Time {
	return time.Time(d)
}

func (d DateTime) Format(format string) string {
	return d.ToTime().Format(format)
}

func (d DateTime) StartDay() DateTime {
	return DateTime(time.Date(d.ToTime().Year(), d.ToTime().Month(), d.ToTime().Day(), 0, 0, 0, 0, time.Local))
}

func (d DateTime) EndDay() DateTime {
	return DateTime(d.NextDay().ToTime().Add(-1))
}

func (d DateTime) NextDay() DateTime {
	return DateTime(d.StartDay().ToTime().Add(24 * time.Hour))
}

func (d DateTime) IsBefore(compDate DateTime) bool {
	return d.ToTime().Before(compDate.ToTime())
}
