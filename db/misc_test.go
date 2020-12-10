package db

import (
	"regexp"
	"testing"
)

func TestParseDate(t *testing.T) {
	date, _ := ParseDate("19930316")
	_, err := regexp.Match("([0-9]+)-([0-9]+)-([0-9])", []byte(date))
	if err != nil {
		t.Error(err)
	}

	t.Logf("This is the date: %v", date)
}
