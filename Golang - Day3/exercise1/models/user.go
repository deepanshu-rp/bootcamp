package models

import (
	f "fmt"
	"sample-api/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// Insert New user
func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// Update user with id
func UpdateUser(user *User, id string) (err error) {
	f.Println(user)
	config.DB.Save(user)
	return nil
}

// Delete user with id
func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
