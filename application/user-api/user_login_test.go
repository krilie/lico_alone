package user_api

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	all_service "github.com/krilie/lico_alone/application/all-service"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestAppUser_UserRegister(t *testing.T) {
	user := NewAppUser(all_service.NewAllService(config.Cfg))
	err := user.UserRegister(context.Background(), "sss", "sss", "")
	t.Log(err)
}

func BenchmarkAppUser_UserRegister(b *testing.B) {
	cdb.Db.LogMode(false)
	user := NewAppUser(all_service.NewAllService(config.Cfg))
	for i := 0; i < b.N; i++ {
		err := user.UserRegister(context.Background(), utils.GetUUID()[:20], "sss", "")
		if err != nil {
			b.Log(err)
		}
	}
}
