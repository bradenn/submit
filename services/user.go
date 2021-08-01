package services

import (
	_ "crypto/sha512"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	CloudKey
	Username string `gorm:"column:username;unique"`
	First    string `gorm:"column:first"`
	Last     string `gorm:"column:last"`
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
}

func NewUser(username string, password string) *User {
	user := &User{
		CloudKey: CloudKey{},
		Username: username,
		First:    "",
		Last:     "",
		Email:    "",
		Password: password,
	}
	return user
}

// AuthUser takes a username and password, returning a user object or an error.
// This method can be used for basic password authentication.
func AuthUser(username string, password string) (user User, err error) {
	// Find a database record of type User with the username provided:
	err = Database().Where(User{Username: username}).First(&user).Error
	// Authentication errors like an invalid password or username are reported here by this
	// return statement.
	if err != nil {
		return user, err
	}
	// Using the stored hash, retrieved previously, verifyHash will either return an error,
	// or do nothing if the password and hash match.
	err = verifyHash(password, user.Password)
	if err != nil {
		return user, err
	}
	// The user has been successfully authenticated
	return user, err
}

func hashPassword(password string) (hashed string, err error) {
	// Hash the Password with bCrypt. This operation is a bit intensive.
	// It will take between 78ms and 120ms depending on the machine.
	// 78.2ms seems to be the fastest possible time.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	// Return any errors from bCrypt.
	if err != nil {
		return hashed, err
	}
	// Convert the byte array 'hash' into a string and return
	return string(hash), nil
}

func verifyHash(suspect string, source string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(source), []byte(suspect))
	// Return any errors from bCrypt.
	if err != nil {
		return err
	}
	// Convert the byte array 'hash' into a string and return
	return err
}

// UpdatePassword hashes the provided Password and sets it as the referenced user's Password.
func (u *User) UpdatePassword(password string) (err error) {
	// Hash the provided Password
	hash, err := hashPassword(password)
	// Return any errors if needed
	if err != nil {
		return err
	}
	// Set the referenced user's Password to the previously hashed string.
	u.Password = hash
	return nil
}

// BeforeCreate intercepts the creation routine for the User object.
// This function takes the creation Password and hashes it before storing it
// in the database.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	plaintext := u.Password
	// Hash the provided Password
	hash, err := hashPassword(plaintext)
	if err != nil {
		return err
	}
	// Set the referenced user's Password to the previously hashed string.
	fmt.Println(hash)
	u.Password = hash
	return nil
}
