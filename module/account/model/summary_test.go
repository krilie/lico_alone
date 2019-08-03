package model

import (
	"fmt"
	"testing"
)

func TestSummaryStruct(T *testing.T) {
	a := AccountSummary{}
	a.Accounts = make([]AccountItem, 0)
	fmt.Println(a)
}
