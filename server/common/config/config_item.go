package config

type Config struct {
	GinMode       string   `mapstructure:"gin_mode" json:"gin_mode" toml:"gin_mode"`
	HttpPort      int      `mapstructure:"http_port" json:"http_port" toml:"http_port"`
	WebPort       int      `mapstructure:"web_port" json:"web_port" toml:"web_port"`
	SslPri        string   `mapstructure:"ssl_pri" json:"ssl_pri" toml:"ssl_pri"`
	SslPub        string   `mapstructure:"ssl_pub" json:"ssl_pub" toml:"ssl_pub"`
	LogFile       string   `mapstructure:"log_file" json:"log_file" toml:"log_file"` // 配置文件 空为控制台
	LogLevel      uint32   `mapstructure:"log_level" json:"log_level" toml:"log_level"`
	EnableSwagger bool     `mapstructure:"enable_swagger" json:"enable_swagger" toml:"enable_swagger"`
	DB            DB       `mapstructure:"db" json:"db" toml:"db"`
	JWT           JWT      `mapstructure:"jwt" json:"jwt" toml:"jwt"`
	FileSave      FileSave `mapstructure:"file_save" json:"file_save" toml:"file_save"`
	Email         Email    `mapstructure:"email" json:"email" toml:"email"`
	AliSms        AliSms   `mapstructure:"ali_sms" json:"ali_sms" toml:"ali_sms"`
	ElfLog        ElfLog   `mapstructure:"elf_log" json:"elf_log" toml:"elf_log"`
}

type DB struct {
	Host            string `mapstructure:"host" json:"host" toml:"host"`
	Port            int    `mapstructure:"port" json:"port" toml:"port"`
	DbName          string `mapstructure:"db_name" json:"db_name" toml:"db_name"`
	User            string `mapstructure:"user" json:"user" toml:"user"`
	Password        string `mapstructure:"password" json:"password" toml:"password"`
	MaxOpenConn     int    `mapstructure:"max_open_conn" json:"max_open_conn" toml:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn" json:"max_idle_conn" toml:"max_idle_conn"`
	ConnMaxLeftTime int    `mapstructure:"conn_max_left_time" json:"conn_max_left_time" toml:"conn_max_left_time"`
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

type Email struct {
	Address  string `mapstructure:"address" json:"address" toml:"address"`
	Host     string `mapstructure:"host" json:"host" toml:"host"`
	Port     int    `mapstructure:"port" json:"port" toml:"port"`
	UserName string `mapstructure:"user_name" json:"user_name" toml:"user_name"`
	Password string `mapstructure:"password" json:"password" toml:"password"`
}

type AliSms struct {
	Key    string `mapstructure:"key" json:"key" toml:"key"`
	Secret string `mapstructure:"secret" json:"secret" toml:"secret"`
}

type ElfLog struct {
	Key    string `mapstructure:"key" json:"key" toml:"key"`
	Secret string `mapstructure:"secret" json:"secret" toml:"secret"`
	Url    string `mapstructure:"url" json:"url" toml:"url"`
}
