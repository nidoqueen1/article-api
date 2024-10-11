package helper

import (
	"strconv"
	"time"
)

// Converts keys of a Map to a list
func MapToList(m map[string]struct{}) []string {
	list := []string{}
	for key := range m {
		list = append(list, key)
	}
	return list
}

// Converts string to Uint
func StringToUint(s string) (uint, error) {
	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(u64), nil
}

// Pars string as time.Time
func StringToDate(s string) (time.Time, error) {
	var formats = []string{
		"2006-01-02", // YYYY-MM-DD
		"20060102",   // YYYYMMDD
		"02/01/2006", // DD/MM/YYYY
		"2006/01/02", // YYYY/MM/DD
		"01-02-2006", // MM-DD-YYYY
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, s)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, err
}
