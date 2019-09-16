package config

type Config struct {
	GinMode      GinMode  `mapstructure:"gin_mode" json:"gin_mode" toml:"gin_mode"`
	DB           DB       `mapstructure:"db" json:"db" toml:"db"`
	QinuiOss     QinuiOss `mapstructure:"qinui_oss" json:"qinui_oss" toml:"qinui_oss"`
	JWT          JWT      `mapstructure:"jwt" json:"jwt" toml:"jwt"`
	OssS3        OssS3    `mapstructure:"oss_s3" json:"oss_s3" toml:"oss_s3"`
	FileSavePath string   `mapstructure:"file_save_path" json:"file_save_path" toml:"file_save_path"`
}

type GinMode struct {
	Mode string `json:"mode" toml:"mode"` // "debug" "release"
}

type DB struct {
	Host     string `mapstructure:"host" json:"host" toml:"host"`
	Port     string `mapstructure:"port" json:"port" toml:"port"`
	DbName   string `mapstructure:"db_name" json:"db_name" toml:"db_name"`
	User     string `mapstructure:"user" json:"user" toml:"user"`
	Password string `mapstructure:"password" json:"password" toml:"password"`
}

type QinuiOss struct {
	AccessKey     string `mapstructure:"access_key" json:"access_key" toml:"access_key"`
	SecretKey     string `mapstructure:"secret_key" json:"secret_key" toml:"secret_key"`
	BucketName    string `mapstructure:"common_bucket_name" json:"common_bucket_name" toml:"common_bucket_name"`
	BucketBaseUrl string `mapstructure:"common_bucket_base_url" json:"common_bucket_base_url" toml:"common_bucket_base_url"`
}

type OssS3 struct {
	OssKey      string `mapstructure:"oss_key" json:"oss_key" toml:"oss_key"`
	OssSecret   string `mapstructure:"oss_secret" json:"oss_secret" toml:"oss_secret"`
	OssEndPoint string `mapstructure:"oss_end_point" json:"oss_end_point" toml:"oss_end_point"`
	OssBucket   string `mapstructure:"oss_bucket" json:"oss_bucket" toml:"oss_bucket"`
}

type JWT struct {
	NormalExpDuration string `mapstructure:"normal_exp_duration" json:"normal_exp_duration" toml:"normal_exp_duration"`
	HS256key          string `mapstructure:"hs_256_key" json:"hs_256_key" toml:"hs256_key"`
	PrivateKeyPath    string `mapstructure:"private_key_path" json:"private_key_path" toml:"private_key_path"`
}
