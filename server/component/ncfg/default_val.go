package ncfg

var defaultCfg = `
[ali_sms]
  key = ""
  secret = ""

[db]
  conn_max_left_time = 14400
  conn_str = "root:123456@tcp(localhost:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
  max_idle_conn = 10
  max_open_conn = 5

[email]
  address = ""
  host = ""
  password = ""
  port = 465
  user_name = ""

[file_save]
  channel = "local"
  oss_bucket = "static"
  oss_end_point = "http://localhost/static"
  oss_key = ""
  oss_secret = ""

[http]
  enable_swagger = false
  gin_mode = "debug"
  port = 80
  ssl_pri = ""
  ssl_pub = ""
  url = "http://localhost"

[jwt]
  hs_256_key = "wDcD3LZl*3L$gmsDd#qSXZ2eMPcM#ps^sWWrt5*zsOoZ5hKAzrsm4&$^Tpg2PIDGoh76hEWVWkCv%cSi%aZXnyXJYC#WxWhuMBp"
  normal_exp_duration = 604800

[log]
  log_file = "log.txt"
  log_level = 5

  [log.elf_log]
    key = ""
    secret = ""
    url = ""
`

/**

http.enable_swagger 		= false
http.gin_mode				= "debug"
http.port 					= 80
http.ssl_pri 				= ""
http.ssl_pub 				= ""
http.url 					= "http://localhost"
log.log_file 				= "log.txt"
log.log_level				= 5
log.elf_log.key				= ""
log.elf_log.secret			= ""
log.elf_log.url				= ""
db.conn_str				 	= "root:123456@tcp(localhost:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
db.max_open_conn			= 5
db.max_idle_conn 		    = 10
db.conn_max_left_time		= 14400
file_save.oss_key 			= ""
file_save.oss_secret		= ""
file_save.oss_end_point		= "http://localhost/static"
file_save.oss_bucket		= "static"
file_save.channel 			= "local"
jwt.normal_exp_duration		= 604800
jwt.hs_256_key 				= "wDcD3LZl*3L$gmsDd#qSXZ2eMPcM#ps^sWWrt5*zsOoZ5hKAzrsm4&$^Tpg2PIDGoh76hEWVWkCv%cSi%aZXnyXJYC#WxWhuMBp"
email.address 				= ""
email.host 					= ""
email.port 					= 465
email.user_name 			= ""
email.password 				= ""
ali_sms.key 				= ""
ali_sms.secret 				= ""

*/
