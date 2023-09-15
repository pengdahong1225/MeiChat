use meiChat;
#用户信息
create table if not exists user_info(
    uid bigint not null,
    account varchar(64) default '',
    pwd varchar(64) not null,
    gender tinyint default 2,
    pic varchar(256) default '',
    
    PRIMARY KEY(uid)
)engine = InnoDB charset = utf8mb4;
