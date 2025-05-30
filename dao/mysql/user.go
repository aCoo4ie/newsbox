package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
)

const SALT = "WAHAHA"

// check if password is equal to password exists
func CheckPasswordEquals(password string, oldPassword string) bool {
	return password == encryptPassword(oldPassword)
}

// CheckUserExists checks if a user with the given username exists
func CheckUserExists(username string) (bool, error) {
	sqlStr := `select count(user_id) from users where username = ?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// QueryUserByUsername fetches a user by their username
func QueryUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	sqlStr := `select user_id, username, password from users where username = ?`
	err := db.Get(user, sqlStr, username)
	return user, err
}

// InsertUser inserts a new user into the database
func InsertUser(uid int64, p models.ParamSignUp) error {
	// Create a new user
	user := &models.User{
		UserID:   uid,
		Username: p.Username,
		Password: p.Password,
	}

	// Hash the password before storing
	user.Password = encryptPassword(user.Password)

	// Insert the user into the database
	sqlStr := `insert into users(user_id, username, password) values(?, ?, ?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(p string) string {
	h := md5.New()
	h.Write([]byte(SALT))
	return hex.EncodeToString(h.Sum([]byte(p)))
}
