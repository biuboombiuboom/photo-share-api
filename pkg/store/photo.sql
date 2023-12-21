CREATE TABLE IF NOT EXISTS  pps.photo(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,user_id BIGINT,
    title VARCHAR(100) ,
    path VARCHAR(500) NOT NULL,
    description VARCHAR(50),
    is_public BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )



alter table pps.photo add column deleted boolean;
alter table pps.photo add column star bigint;
alter table pps.photo add column collect bigint;
alter table pps.photo add column comment bigint;

CREATE TABLE IF NOT EXISTS pps.comment(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id bigint,
    username nvarchar(50),
    photo_id bigint,
    reply_to bigint,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    content text    
)

CREATE TABLE IF NOT EXISTS pps.photo_collect(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id bigint,
    username nvarchar(50),
    photo_id bigint,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS pps.photo_star(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id bigint,
    username nvarchar(50),
    photo_id bigint,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

alter table pps.photo_star add unique index photo_userid (user_id,photo_id);
alter table pps.photo_collect add unique index photo_userid (user_id,photo_id);
alter table pps.comment modify column content text CHARACTER SET utf8mb4 not NULL；
alter table pps.user add column nickname nvarchar(50) DEFAULT '';

alter table pps.user add unique index idx_email(email);
alter table pps.user add unique index idx_user(username);
alter table pps.user add column  description nvarchar(2000) default '';
