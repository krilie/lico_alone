# account info

create table tb_account (
    id varchar(32) primary key comment '主键',
    user_id varchar(32) not null comment '所有者用户的id',
    name varchar(50) not null comment '帐户名',
    create_time datetime not null comment '存北京时间，取北京时间',
    update_time datetime not null comment '存北京时间，取北京时间',
    num varchar(50) not null comment '编号',
    balance decimal(10,2) not null default 0 comment '金钱 余额 单位 元',
    description varchar(100) not null comment '账户描述'
)  engine = innodb DEFAULT CHARSET=utf8mb4 comment = '帐户表';

create table tb_bill(

)engine = innodb DEFAULT CHARSET=utf8mb4 comment = '一笔帐';

create table tb_bill_detail(

)engine = innodb DEFAULT CHARSET=utf8mb4 comment = '一笔帐细节';

