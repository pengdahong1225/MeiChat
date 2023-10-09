use meiChat;
#用户信息
create table if not exists user_info(
    uid BIGINT NOT NULL,
    pwd VARCHAR(64) NOT null,
    phone BIGINT not NULL,         # 手机号唯一注册，不能为空
    email VARCHAR(64) DEFAULT '',
    username VARCHAR(64) DEFAULT '新用户',
    gender tinyint DEFAULT 0,   # 0:woman 1:man
    role tinyint DEFAULT 0,     # 0:user 1:admin
    head_pic VARCHAR(256) DEFAULT '',
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY(uid)
)engine = InnoDB charset = utf8mb4;
