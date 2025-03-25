USE dbname;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, age, created_at) VALUES
('Taro', 25, '2025-03-25 00:00:00'),
('Hanako', 30, '2025-03-20 00:00:00'),
('Jiro', 22, '2025-03-22 00:00:00');