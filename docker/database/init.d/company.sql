CREATE TABLE company (
    company_id INT AUTO_INCREMENT PRIMARY KEY,   -- 会社ID (自動増加)
    company_name VARCHAR(255) NOT NULL,          -- 会社名 (必須)
    established_date DATE                      -- 設立日
)