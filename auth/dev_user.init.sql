CREATE TABLE users (
    --id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    --email VARCHAR(255) NOT NULL,
    --firstName VARCHAR(255),
    --lastName VARCHAR(255),
    password VARCHAR(255) NOT NULL --created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    --updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Path: dev_user.session.sql
/*CREATE TABLE sessions (
 id SERIAL PRIMARY KEY,
 user_id INTEGER NOT NULL,
 token VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (user_id) REFERENCES users(id)
 );*/