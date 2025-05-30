package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"database/sql"
	"errors"
)

// Error definitions
var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found in databases")
)

// SignUp handles the user registration logic
func SignUp(p models.ParamSignUp) error {
	// Check if user already exists
	exists, err := mysql.CheckUserExists(p.Username)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserExists
	}

	// Generate unique user ID
	userId := snowflake.GenId()

	// Save to MySQL
	return mysql.InsertUser(userId, p)
}

// Login handles the user login logic
func Login(p models.ParamLogin) error {
	u, err := mysql.QueryUserByUsername(p.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}
	ok := mysql.CheckPasswordEquals(u.Password, p.Password)
	if !ok {
		return ErrUserNotFound
	}
	return nil
}
