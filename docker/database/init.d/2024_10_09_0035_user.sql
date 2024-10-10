CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    first_name VARCHAR(255) NOT NULL,  
    last_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_user_id INT ,
    updated_at TIMESTAMP NOT NULL,
    updated_user_id INT 
    -- CONSTRAINT fk_user_user_createduserid
    -- FOREIGN KEY (created_user_id) 
    -- REFERENCES user (id)
    -- CONSTRAINT fk_user_user_updateduserid
    -- FOREIGN KEY (updated_user_id) 
    -- REFERENCES user (id)
)