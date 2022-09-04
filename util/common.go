package util

import (
	"encoding/binary"
	"strconv"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func CapitalizeString(s string) string {
	return cases.Title(language.English, cases.Compact).String(s)
}

func MaxLen(slice []string) int {
	max := 0
	for _, s := range slice {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

func StringSliceToIntSlice(slice []string) ([]int, error) {
	ints := make([]int, len(slice))

	for i, s := range slice {
		n, err := strconv.Atoi(s)
		if err != nil {
			return ints, err
		}
		ints[i] = n
	}

	return ints, nil
}

func StringSliceToDurationSlice(slice []string) ([]time.Duration, error) {
	durations := make([]time.Duration, len(slice))

	for i, s := range slice {
		d, err := time.ParseDuration(s)
		if err != nil {
			return durations, err
		}
		durations[i] = d
	}

	return durations, nil
}
