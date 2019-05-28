# account info
create table tb_account (
    id varchar(32) primary key comment '主键',
    user_id varchar(32) not null comment '所有者用户的id',
    name varchar(50) not null comment '帐户名',
    create_time datetime not null comment '存北京时间，取北京时间',
    update_time datetime not null comment '存北京时间，取北京时间',
    num varchar(50) not null comment '编号',
    debit decimal(10,2) not null default 0 comment '借 数量',
    credit decimal(10,2) not null default 0 comment '贷 数量',
    balance decimal(10,2) not null default 0 comment '金钱 余额 单位 元',
    description varchar(100) not null comment '账户描述',
    image varchar(200) default null comment '帐户图像一张'
)  engine = innodb DEFAULT CHARSET=utf8mb4 comment = '帐户表';
# 一笔交易
create table tb_bill(
    id varchar(32) primary key comment '主键',
    user_id varchar(32) not null comment '用户的id',
    create_time datetime not null comment '',
    note varchar(100) not null comment '',
    amount decimal(10,2) not null default 0 comment '',
    image varchar(500) default null comment '',
    is_valid boolean default true comment '是否有效'
)engine = innodb DEFAULT CHARSET=utf8mb4 comment = '一笔帐';

create table tb_bill_detail(
    id varchar(32) primary key comment '主键',
    user_id varchar(32) not null comment '用户的id',
    bill_id varchar(32) not null comment '交易id',
    account_id varchar(32) not null comment '帐户id',
    account_num varchar(50) not null comment '帐户号',
    account_name varchar(50) not null comment '帐户名称',
    note varchar(100) default null comment '',
    debit decimal(10,2) not null default 0 comment '借 数量',
    credit decimal(10,2) not null default 0 comment '贷 数量',
    balance decimal(10,2) not null default 0 comment '之后的余额'
)engine = innodb DEFAULT CHARSET=utf8mb4 comment = '一笔帐细节';

