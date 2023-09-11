package timeutil

import (
	"github.com/markusmobius/go-dateparser"
	"time"
)

//const maxNanoSec = 999999999

func ParseDate(value string) (time.Time, error) {
	date, err := dateparser.Parse(&dateparser.Configuration{
		CurrentTime: time.Now(),
	}, value)
	return date.Time, err
}

//func ParsePeriod(value string) (startTime time.Time, endTime time.Time, err error) {
//	currentTime := time.Now()
//
//	switch value {
//	case "year":
//		startTime = time.Date(currentTime.Year(), time.January, 1, 0, 0, 0, 0, currentTime.Location())
//		endTime = time.Date(currentTime.Year(), time.December, 31, 23, 59, 59, maxNanoSec, currentTime.Location())
//		break
//
//	case "month":
//		startTime = time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
//		endTime = startTime.AddDate(0, 1, 0).Add(-time.Second)
//		break
//
//	case "week":
//		daysUntilMonday := 7 - int(currentTime.Weekday()) + int(time.Monday)
//		startTime = currentTime.AddDate(0, 0, daysUntilMonday).Truncate(24 * time.Hour)
//		endTime = startTime.AddDate(0, 0, 6)
//		break
//
//	case "day":
//		startTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
//		endTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, maxNanoSec, currentTime.Location())
//		break
//
//	default:
//		err = errors.New("unknown time period (only year, month, week and day can be used)")
//		break
//	}
//
//	return startTime, endTime, err
//}
