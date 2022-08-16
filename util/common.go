package util

import (
	"encoding/binary"

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
