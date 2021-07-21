extern crate serde;
extern crate toml;

use serde_derive::Serialize;

#[derive(Debug, Clone, Serialize, Deserialize, )]
pub struct DbConfig {
    pub conn_str: String,
}

#[derive(Serialize)]
pub struct TestVal {
    pub name: String,
    pub http_addr: String,
    pub db: DbConfig,
}

#[test]
pub fn test2() {
    let test = TestVal { name: "123".to_string(), db: DbConfig { conn_str: "234".to_string() }, http_addr: "123".to_string() };
    let val = toml::to_string(&test).unwrap();
    println!("{}", val);
}
