use tide::new;
use std::error::Error;
use std::fmt::{Display, Formatter};

#[derive(Debug, Clone)]
pub struct LicoError {
    pub msg: String,
    pub kind: i32,
}

impl Display for LicoError {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "kind:{} msg:{}", self.kind, self.msg)
    }
}

impl Error for LicoError { /* 没有子错误*/ }

impl LicoError {
    pub fn new() -> LicoError { LicoError { msg: "success".to_string(), kind: 2000 } }
    pub fn with_msg(&mut self, msg: String) -> &mut LicoError {
        self.msg = msg;
        self
    }
    pub fn with_kind(&mut self, kind: i32) -> &mut LicoError {
        self.kind = kind;
        self
    }
    pub fn with_msg_kind(&mut self, kind: i32, msg: String) -> &mut LicoError {
        self.with_kind(kind);
        self.with_msg(msg);
        self
    }
    pub fn build(&mut self) -> LicoError {
        LicoError { msg: self.msg.clone(), kind: self.kind.clone() }
    }
    // 常有错误
    pub fn not_found() -> LicoError { LicoError { kind: 4004, msg: "not found".to_string() } }
    pub fn internal_err() -> LicoError { LicoError { kind: 5000, msg: "internal err".to_string() } }
}
