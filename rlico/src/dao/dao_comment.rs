use sqlx::{Executor, MySql, Database};
use crate::model::TbComment;
use crate::util::errors::LicoError;
use async_std::sync::Arc;
use crate::dao::create_mysql_pool;
use chrono::{DateTime, Local};
use tide::http::cookies::EncodedCookie;

pub async fn add_comment<'e, 'c: 'e,  E,DB>(executor:E, comment: &TbComment) -> Result<(), LicoError>
    where
        E: 'e + Executor<'c, Database = DB>,
        DB: 'e,
{
    let sql = r######"insert into
                          tb_comment(id, created_at, updated_at, deleted_at, user_id, comment_id, target_id, content, like_count, dislike_count, is_check)
                          VALUES(?,?,?,?,?,?,?,?,?,?,?)"######;
    let _ = sqlx::query(sql)
        .bind("wy")
        .bind("123456")
        .execute(executor).await?;
    Ok(())
}

#[test]
pub fn test() {
    async_std::task::block_on(async {
        let mut conn = create_mysql_pool("mysql://test:123456@lizo.top/test").await?;
        let comment = add_comment(&mut conn.acquire().await?, &TbComment {
            create_at: Local::now(),
            update_at: Local::now(),
            delete_at: None,
            id: "".to_string(),
            user_id: "".to_string(),
            comment_id: "".to_string(),
            target_id: "".to_string(),
            content: "".to_string(),
            like_count: 0,
            dislike_count: 0,
            is_check: false,
        }).await;
        match comment {
            Ok(_) => {}
            Err(err) => { println!("{}", err) }
        }
    })
}
