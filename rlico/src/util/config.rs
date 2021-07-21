extern crate serde;
extern crate toml;

#[derive(Debug, Clone, Serialize, Deserialize, )]
pub struct Config {
    pub db_config: DbConfig,
    #[serde(default)]
    pub http_addr: Vec<String>,

}

#[derive(Debug, Clone, Serialize, Deserialize, )]
pub struct DbConfig {
    pub conn_str: String,
}

#[test]
pub fn test() {
    let cfg = Config { db_config: DbConfig { conn_str: "mysql://test:123456@lizo.top/test".to_string() }, http_addr: vec!["0.0.0.0:80".to_string()] };
    println!("{}", serde_json::to_string(&cfg).unwrap());
    match toml::to_string(&cfg) {
        Ok(valStr) => { println!("{}", valStr) }
        Err(err) => {
            println!("err {:?}", err)
        }
    }
}
