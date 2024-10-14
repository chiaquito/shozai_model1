CREATE TABLE company (
    id INT AUTO_INCREMENT PRIMARY KEY,   -- 会社ID (自動増加)
    `name` VARCHAR(255) NOT NULL,          -- 会社名 (必須)
    established_date DATE,                      -- 設立日
    created_at TIMESTAMP NOT NULL,
    created_user_id INT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_user_id INT NOT NULL,

    CONSTRAINT fk_company_user_createduserid
    FOREIGN KEY (created_user_id) 
    REFERENCES user (id),
    CONSTRAINT fk_company_user_updateduserid
    FOREIGN KEY (updated_user_id) 
    REFERENCES user (id)
);