package util

import (
	"os"
	"strconv"
)

func GetenvDefault(key string, default_value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_value
	} else {
		return val
	}
}

func GetenvIntDefault(key string, default_value int) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_value, nil
	} else {
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		return v, nil
	}
}
