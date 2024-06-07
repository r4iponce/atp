package atp

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// checkCharRedundant  verify is a character is redundant in a string.
func checkCharRedundant(source string, char string) bool { // Verify if a char is redundant
	return strings.Count(source, char) > 1
}

var errRedundantChar = errors.New("redundant characters")

// parseDurationBeforeSeparator parse duration before defined separator.
// return 0 (and error) when error, only take positive number.
func parseDurationBeforeSeparator(source *string, sep string, duration time.Duration) (time.Duration, error) {
	if checkCharRedundant(*source, sep) {
		return 0, errRedundantChar
	}

	if !strings.Contains(*source, sep) {
		return 0, nil
	}

	var value int
	var err error

	value, err = strconv.Atoi(strings.Split(*source, sep)[0])
	if err != nil {
		return 0, errors.New("cannot parse: " + strings.Split(*source, sep)[0] + "as int")
	}
	if value < 0 { // Only positive value is correct
		return 0, errors.New("format only take positive value:" + strings.Split(*source, sep)[0] + "as int")
	}

	*source = strings.Join(strings.Split(*source, sep)[1:], "")

	return time.Duration(value) * duration, nil
}

// ParseDuration Parse "1d1h1m1s" duration format. Return 0 & error if error.
func ParseDuration(source string) (time.Duration, error) {
	unitMap := map[string]time.Duration{
		"d": 24 * time.Hour, //nolint:mnd
		"h": time.Hour,
		"m": time.Minute,
		"s": time.Second,
	}

	orderedKey := []string{"d", "h", "m", "s"}

	if source == "" {
		return 0, errors.New("empty string")
	}

	if source == "0" {
		return 0, nil
	}

	var expiration time.Duration
	var tempOutput time.Duration
	var err error

	source = strings.ToLower(source)

	for _, k := range orderedKey {
		tempOutput, err = parseDurationBeforeSeparator(&source, k, unitMap[k])
		if err != nil {
			return 0, err
		}
		expiration = tempOutput + expiration
	}

	return expiration, nil
}
