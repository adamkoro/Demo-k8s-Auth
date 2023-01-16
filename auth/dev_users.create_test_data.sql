-- Create a test user
INSERT INTO users (username, password)
VALUES ('test', '12345');
SELECT *
FROM users;
SELECT *
FROM users
WHERE username = 'test';