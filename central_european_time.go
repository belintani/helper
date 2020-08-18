package helper

import (
	"strings"
	"time"
)

func ParseUTCEuropeanTime(s string) (*time.Time, error) {
	var offset int
	if strings.Contains(s, "CET") {
		offset = -1
	} else if strings.Contains(s, "CEST") {
		offset = -2
	} else {
		offset = 0
	}

	s = ReplaceAllSpaces(s, " ")
	s = strings.ReplaceAll(s, "CET", "")
	s = strings.ReplaceAll(s, "CEST", "")
	s = strings.TrimSpace(s)

	layout := "Jan 02 2006, 15:04:05"
	timeValue, err := time.Parse(layout, s)

	if err != nil {
		return nil, err
	}

	v := timeValue.Add(time.Hour * time.Duration(offset))
	return &v, nil

}
