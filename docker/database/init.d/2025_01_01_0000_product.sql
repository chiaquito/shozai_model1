CREATE TABLE product (
    id INT AUTO_INCREMENT PRIMARY KEY,     -- 製品ID (自動増加)
    `name` VARCHAR(255) NOT NULL,          -- 製品名 (必須)
    `version` INT NOT NULL,
    code     CHAR(5) NOT NULL,             -- 製品コード
    description  VARCHAR(255) NOT NULL,    -- 製品説明
    company_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_user_id INT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_user_id INT NOT NULL,

    CONSTRAINT fk_product_company_companyid
    FOREIGN KEY (company_id) 
    REFERENCES company (id),
    CONSTRAINT fk_product_user_createduserid
    FOREIGN KEY (created_user_id) 
    REFERENCES user (id),
    CONSTRAINT fk_product_user_updateduserid
    FOREIGN KEY (updated_user_id) 
    REFERENCES user (id),
    UNIQUE (code)
);