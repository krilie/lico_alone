use sqlx::Executor;
use crate::model::TbComment;
use crate::util::errors::LicoError;

// pub struct CommentDao<'a> {
//     pub executor: &'a Executor,
// }
//
// impl<DB, E: Executor<Database=DB>> CommentDao<DB, E> {
//     pub async fn add_comment(&self, comment: &mut TbComment) -> Option<LicoError> {
//         self.executor.execute()
//     }
// }
