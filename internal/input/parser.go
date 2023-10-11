package input

import (
	"errors"
	"github.com/markusmobius/go-dateparser"
	"strconv"
	"strings"
	"time"
)

func ParseTime(x string) (time.Time, error) {
	currentTime := time.Now()
	result, err := dateparser.Parse(&dateparser.Configuration{
		CurrentTime: currentTime,
	}, x)
	return result.Time, err
}

func ParseOptionalTime(x string) (*time.Time, error) {
	if x == "" {
		return nil, nil
	}
	result, err := ParseTime(x)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ParseOptionalString(x string) *string {
	if x == "" {
		return nil
	}
	return &x
}

func ParseAmount(x string) (float64, error) {
	lastChar := string(x[len(x)-1])
	if lastChar == "-" {
		x = "-" + x[0:len(x)-1]
	}
	return strconv.ParseFloat(x, 64)
}

func ParseCurrency(x string) (string, error) {
	if len(x) != 3 {
		return "", errors.New("invalid currency format")
	}
	return strings.ToUpper(x), nil
}

func ParseOptionalCurrency(x string) (*string, error) {
	if x == "" {
		return nil, nil
	}
	currency, err := ParseCurrency(x)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}
