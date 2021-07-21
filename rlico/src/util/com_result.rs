extern crate serde;
extern crate serde_json;
use crate::util::errors::LicoError;
use tide::new;
use crate::json;
use std::any::Any;

#[derive(Debug, Clone,Serialize, Deserialize,)]
pub struct ComResult<Data> {
    pub code: i32,
    pub message: String,
    pub data: Data,
}

impl<Data> ComResult<Data> {
    pub fn new(code: i32, message: String, data: Data) -> Self {
        ComResult { code, message, data }
    }

    pub fn new_from_err(err: &LicoError) -> ComResult<Option<()>> {
        ComResult::new(err.kind, err.msg.to_string(), Option::None)
    }
}


#[test]
pub fn show(){
    println!("test");
    let result = ComResult{
        code: 230,
        message: "ad".to_string(),
        data: json!({"a":"b"})
    };
    let result1 = serde_json::to_string(&result).unwrap();
    println!("{}",result1);
    let obj:ComResult<serde_json::Value> = serde_json::from_str(result1.as_str()).unwrap();
    println!("{:?}",obj)
}
