package context_enum

import "strconv"

type IntEnum int

func (enum IntEnum) ToInt() int {
	return int(enum)
}
func (enum IntEnum) ToStr() string {
	return strconv.Itoa(int(enum))
}
