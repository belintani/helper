package helper

import (
	"regexp"
	"strconv"
	"strings"
)

func StringToInt16(val string) (int16, error) {

	val = strings.TrimSpace(val)
	if "" == val {
		return 0, nil
	}

	reg, _ := regexp.Compile("[^0-9]+")
	i, err := strconv.ParseInt(reg.ReplaceAllString(val, ""), 10, 16)

	if err != nil {
		return 0, err
	}

	return int16(i), nil
}

func StringToInt32(val string) (int32, error) {
	val = strings.TrimSpace(val)
	if "" == val {
		return 0, nil
	}

	reg, _ := regexp.Compile("[^0-9]+")
	i, err := strconv.ParseInt(reg.ReplaceAllString(val, ""), 10, 32)

	if err != nil {
		return 0, err
	}

	return int32(i), nil
}

func StringToInt64(val string) (int64, error) {
	val = strings.TrimSpace(val)
	if "" == val {
		return 0, nil
	}

	reg, _ := regexp.Compile("[^0-9]+")
	i, err := strconv.ParseInt(reg.ReplaceAllString(val, ""), 10, 64)

	if err != nil {
		return 0, err
	}

	return int64(i), nil
}
