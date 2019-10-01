package config

type Config struct {
	GinMode  GinMode  `mapstructure:"gin_mode" json:"gin_mode" toml:"gin_mode"`
	DB       DB       `mapstructure:"db" json:"db" toml:"db"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" toml:"jwt"`
	FileSave FileSave `mapstructure:"file_save" json:"file_save" toml:"file_save"`
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
type FileSave struct {
	OssKey           string `mapstructure:"oss_key" json:"oss_key" toml:"oss_key"`
	OssSecret        string `mapstructure:"oss_secret" json:"oss_secret" toml:"oss_secret"`
	OssEndPoint      string `mapstructure:"oss_end_point" json:"oss_end_point" toml:"oss_end_point"`
	OssBucket        string `mapstructure:"oss_bucket" json:"oss_bucket" toml:"oss_bucket"`
	LocalFileSaveUrl string `mapstructure:"local_file_save_url" json:"local_file_save_url" toml:"local_file_save_url"` // 本地url
	LocalFileSaveDir string `mapstructure:"local_file_save_dir" json:"local_file_save_dir" toml:"local_file_save_dir"` // 本地路径
	SaveType         string `mapstructure:"save_type" json:"save_type" toml:"save_type"`                               // 保存类型
}
type JWT struct {
	NormalExpDuration string `mapstructure:"normal_exp_duration" json:"normal_exp_duration" toml:"normal_exp_duration"`
	HS256key          string `mapstructure:"hs_256_key" json:"hs_256_key" toml:"hs256_key"`
	PrivateKeyPath    string `mapstructure:"private_key_path" json:"private_key_path" toml:"private_key_path"`
}
