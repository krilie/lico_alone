use sqlx::mysql::{MySqlPoolOptions, MySqlRow};
use sqlx::query::{Map, Query};
use sqlx::{Error, FromRow, Connection, Acquire, Pool, MySql};
use std::any::Any;
use async_std::future::Future;

pub async fn create_mysql_pool(conn_str: &str) -> Result<Pool<MySql>, Error> {
    MySqlPoolOptions::new()
        .max_connections(5)
        .connect(conn_str).await
}

#[test]
fn test() -> Result<(), Error> {
    async_std::task::block_on( async {
        let pool = create_mysql_pool("mysql://test:123456@lizo.top/test").await?;
        pool.close();

        #[derive(sqlx::FromRow, Debug)]
        struct forTest {
            pub id: i32,
            pub age: i32,
            pub r#type: Option<i32>,
            pub r#int: Option<i32>,
        }
        let mut x1 = pool.begin().await?;
        let mut connection = pool.acquire().await?;
        connection.ping().await?;

        // sqlx::query("DELETE FROM tablename WHERE 1=2").execute(&pool).await?;
        let x: Vec<forTest> = sqlx::query_as("select * from for_test").fetch_all(&mut connection).await?;
        for x2 in x {
            println!("{:?}", x2);
        }

        let mut stream = sqlx::query_as::<_, forTest>("select * from for_test where name=?")
            .bind("user_email")
            .fetch(&mut connection);

        Ok(())
    })
}
