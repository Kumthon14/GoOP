package Models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"Go_OOP/Technical_Service/Entity/EntityStruct"
)

var globalAdapterInstance *gorm.DB
var adapter *Adapter

const connectionString = "sqlserver://sa:1234@127.0.0.1:57329?database=first_go&instance=SQLEXPRESS"

type Adapter struct{}

func (s *Adapter) newAdapter() {}

func (s *Adapter) GetAdapterIntance() *Adapter {
	if adapter == nil {
		db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

		if err != nil {
			panic("failed to connect database: " + err.Error())
		}

		globalAdapterInstance = db

		adapter = new(Adapter)
		adapter.newAdapter()
	}
	return adapter
}

func (s *Adapter) GetGormInstance() *gorm.DB {
	return globalAdapterInstance
}

func (s *Adapter) Login(user *User) error {
	return globalAdapterInstance.Where(&User{
		User: &EntityStruct.User{
			Username: user.Username, Password: user.Password,
		},
	}).First(user).Error
}
