CREATE TABLE IF NOT EXISTS  pps.photo(id BIGINT PRIMARY KEY AUTO_INCREMENT,user_id BIGINT,title VARCHAR(100) ,path VARCHAR(500) NOT NULL,description VARCHAR(50),is_public BOOLEAN,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)



alter table pps.photo add column deleted boolean;