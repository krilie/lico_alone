package ginutil

import "strconv"

func ParseBoolOrNil(val string) *bool {
	parseBool, err := strconv.ParseBool(val)
	if err != nil {
		return nil
	}
	return &parseBool
}

func ParseBoolOrDefault(val string, def bool) bool {
	parseBool, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return parseBool
}

func ParseIntOrNil(val string) *int64 {
	parseInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil
	}
	return &parseInt
}

func ParseIntOrDefault(val string, def int64) int64 {
	parseInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return def
	}
	return parseInt
}
