package dates

import (
	"errors"
	"github.com/markusmobius/go-dateparser"
	"time"
)

func ParseDate(value string) (time.Time, error) {
	date, err := dateparser.Parse(&dateparser.Configuration{
		CurrentTime: time.Now(),
	}, value)
	return date.Time, err
}

func ParsePeriod(value string) (startTime time.Time, endTime time.Time, err error) {
	switch value {
	case "year":
		break
	case "month":
		break
	case "week":
		break
	case "day":
		break
	default:
		return startTime, endTime, errors.New("unknown time period (only year, month, week and day can be used)")
	}

	return startTime, endTime, nil
}
