use chrono::{DateTime, Local};

#[derive(sqlx::FromRow, Debug)]
pub struct TbComment {
    pub create_at: DateTime<Local>,
    pub update_at: DateTime<Local>,
    pub delete_at: Option<DateTime<Local>>,

    pub id: String,
    pub user_id: String,
    pub comment_id: String,
    pub target_id: String,

    pub content: String,
    pub like_count: i32,
    pub dislike_count: i32,
    pub is_check: bool,
}
