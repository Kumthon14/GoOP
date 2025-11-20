package Models

import (
	"fmt"

	"Go_OOP/Technical_Service/Entity/EntityStruct"

	_ "gorm.io/driver/sqlserver"
)

type User struct {
	*EntityStruct.User
}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" `
	Address  string `json:"address"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRes struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone" `
	Address  string `json:"address"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (b *User) TableName() string {
	return "User"
}

func (u *User) CreateUser(user *User) (err error) {
	if err = globalAdapterInstance.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GetAllUsers(user *[]User) error {
	if err := globalAdapterInstance.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserById(user *User, id string) error {
	if err := globalAdapterInstance.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(user *User, id string) error {
	fmt.Println(user)
	globalAdapterInstance.Save(user)
	return nil
}

func (u *User) DeleteUser(user *User, id string) error {
	globalAdapterInstance.Where("id = ?", id).Delete(user)
	return nil
}
