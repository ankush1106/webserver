package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Email    string
	Name     string
	Phone    string
	Password string
}

type DBConfig struct {
	Type string
	URL  string
}

type DBClient struct {
	*gorm.DB
}

func ConnectDB(config *DBConfig) DBClient {
	db, err := gorm.Open(config.Type, config.URL)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return DBClient{db}

}

func (c *DBClient) FindUser(email string) *User {
	var user User
	c.First(&user, "email = ?", email)
	if user.Email == "" {
		return nil
	}
	return &user
}

func (c *DBClient) AuthenticateUser(email string, password string) *User {
	user := c.FindUser(email)
	if user.Email == "" {
		return nil
	}
	if user.Password == password {
		return user
	}
	return nil

}

func (c *DBClient) CreateUser(user *User) error {
	if exist := c.FindUser(user.Email); exist != nil {
		return errors.New("user alread exist")
	}
	c.Create(user)
	return nil

}
