CREATE TABLE db_test_cdc.tbl_test (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO db_test_cdc.tbl_test (name, email) VALUES ('John Doe', 'john.doe@example.com');