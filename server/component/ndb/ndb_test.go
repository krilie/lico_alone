package ndb

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/stretchr/testify/assert"
	"testing"
)

var container = appdig.NewAppDig()

func TestMain(m *testing.M) {

	container.MustProvide(ncfg.NewNConfigByCfgStrFromEnvJson("MYAPP_TEST_CONFIG"))
	container.MustProvide(nlog.NewLogger)
	container.MustProvide(NewNDb)

	m.Run()
}

func TestNewNDb(t *testing.T) {
	container.MustInvoke(func(db *NDb) {
		err := db.Ping()
		assert.Equal(t, nil, err, "should not err")
	})
}

func TestMigrate(t *testing.T) {
	container.MustInvoke(func(db *NDb, cfg *ncfg.NConfig) {
		err := db.Ping()
		assert.Equal(t, nil, err, "should not err")
		aff, err := db.Exec(context.Background(), "begin;CREATE TABLE `tb_article_master` (\n  `id` char(36) NOT NULL,\n  `created_at` datetime(3) NOT NULL,\n  `updated_at` datetime(3) NOT NULL,\n  `deleted_at` datetime(3) DEFAULT NULL,\n  `title` varchar(256) NOT NULL,\n  `description` varchar(512) NOT NULL,\n  `content` text NOT NULL,\n  `picture` varchar(512) NOT NULL,\n  `sort` int(11) NOT NULL,\n  `pv` int(11) NOT NULL,\n  PRIMARY KEY (`id`),\n  KEY `idx_tb_article_master_deleted_at` (`deleted_at`),\n  KEY `idx_tb_article_master_sort` (`sort`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;\n\nCREATE TABLE `tb_auth_permission` (\n  `created_at` datetime(3) DEFAULT NULL,\n  `updated_at` datetime(3) DEFAULT NULL,\n  `deleted_at` datetime(3) DEFAULT NULL,\n  `name` varchar(32) NOT NULL,\n  `description` varchar(100) NOT NULL,\n  `ref_method` varchar(255) NOT NULL,\n  `ref_path` varchar(255) NOT NULL,\n  `sort` int(11) NOT NULL,\n  PRIMARY KEY (`name`),\n  KEY `idx_tb_auth_permission_deleted_at` (`deleted_at`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;commit;")
		t.Log(aff)
		t.Log(err)
		//innerDb, _ := db.db.DB()
		//dbmigrate.Migrate(innerDb, "test", "file://c://sqls", 20210206140300)
	})
}
