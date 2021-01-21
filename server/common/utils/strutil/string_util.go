package strutil

import (
	"database/sql"
	"github.com/deckarep/golang-set"
	"strconv"
	"strings"
)

func NewString(str string) *string {
	return &str
}

func SqlStringOrEmpty(str sql.NullString) string {
	if str.Valid {
		return str.String
	} else {
		return ""
	}
}

func EmptyOrDefault(ori, def string) string {
	if ori == "" {
		return def
	} else {
		return ori
	}
}

func MustToString(str interface{}) string {
	value, ok := str.(string)
	if !ok {
		panic("convert interface{} to string err")
	}
	return value
}

func JoinWith(set mapset.Set, sep string) string {
	a := set.ToSlice()
	switch len(a) {
	case 0:
		return ""
	case 1:
		return MustToString(a[0])
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(MustToString(a[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(MustToString(a[0]))
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(MustToString(s))
	}
	return b.String()
}

func GetInt64(ori string) (int64, error) {
	return strconv.ParseInt(ori, 10, 64)
}

func GetIntOrDef(ori string, def int) int {
	parseInt, err := strconv.ParseInt(ori, 10, 32)
	if err != nil {
		return def
	}
	return int(parseInt)
}

func GetFloat64(ori string) (float64, error) {
	return strconv.ParseFloat(ori, 64)
}
