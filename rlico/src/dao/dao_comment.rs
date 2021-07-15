use sqlx::Executor;

pub struct CommentDao<DB, E: Executor<Database=DB>> {
    pub executor: E,
}

impl<DB, E: Executor<Database=DB>> CommentDao<DB, E> {

}
