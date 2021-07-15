use tide::new;

#[derive(Debug)]
pub struct LicoError {
    pub msg: String,
    pub kind: i32,
}

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
    // 常有错误
    pub fn not_found() -> &mut LicoError { LicoError::new().with_msg_kind(4004, "not found".to_string()) }
    pub fn internal_err() -> &mut LicoError { LicoError::new().with_msg_kind(5000, "internal err".to_string()) }
}
