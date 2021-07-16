use sqlx::Executor;
use crate::model::TbComment;
use crate::util::errors::LicoError;
use async_std::sync::Arc;

pub struct CommentDao<'a> {
    pub executor: Arc<dyn Executor<'a>>,
}

impl<'a> CommentDao<'a> {
    pub async fn add_comment(&mut self, comment: TbComment) -> Result<(), LicoError> {
        let sql = r######"insert into
                          tb_comment(id, created_at, updated_at, deleted_at, user_id, comment_id, target_id, content, like_count, dislike_count, is_check)
                          VALUES(?,?,?,?,?,?,?,?,?,?,?)"######;
        let _ = sqlx::query(sql).bind("wy").bind("123456").execute(&mut self.executor).await?;
        Ok(())
    }
}

