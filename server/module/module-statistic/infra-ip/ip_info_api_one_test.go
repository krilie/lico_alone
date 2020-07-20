// +build !auto_test

package infra_ip

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

func TestIpInfoApiOne_GetIpInfo(t *testing.T) {
	var apiOne = NewIpInfoApiOne()
	info, err := apiOne.GetIpInfo(context.Background(), "223.104.147.153")
	fmt.Println(str_util.ToJsonPretty(info), err)
}
