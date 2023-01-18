package service

import (
	"database/sql"
	"demo-k8s-auth/model"
)

// Get all users
func GetAllUsers(db *sql.DB) ([]model.User, error) {
	var users []model.User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.LastName, &user.FirstName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get user by id
func GetUserById(db *sql.DB, id int) (model.User, error) {
	var user model.User
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get user by username
func GetUserByUsername(db *sql.DB, username string) (model.User, error) {
	var user model.User
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
	err := row.Scan(&user.ID, &user.Username, &user.LastName, &user.FirstName, &user.Email, &user.Phone, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get user by email
func GetUserByEmail(db *sql.DB, email string) (model.User, error) {
	var user model.User
	row := db.QueryRow("SELECT * FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Create a new user
func CreateUser(db *sql.DB, user *model.User) (model.User, error) {
	err := db.QueryRow("INSERT INTO users (id, username, password, email, phone, lastName, firstName, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		user.ID, user.Username, user.Password, user.Email, user.Phone, user.LastName, user.FirstName, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return *user, err
	}
	return *user, nil
}

// Update user by id
func UpdateUserById(db *sql.DB, id int, user model.User) (model.User, error) {
	_, err := db.Exec("UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4",
		user.Username, user.Password, user.Email, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Delete user by id
func DeleteUserById(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
