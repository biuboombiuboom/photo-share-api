CREATE TABLE IF NOT EXISTS pps.message(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    content text,
    from_user_id bigint,
    form_username nvarchar(50),
    to_user_id bigint,
    to_username nvarchar(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    send_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    read_state boolean);