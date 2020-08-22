package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"sync"

	"microconsultor/libs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mu sync.Mutex

type User struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	Account  string `gorm:"type:varchar(20);not null;index:username" validate:"required"`
	Password string `gorm:"type:char(32);not null;" validate:"required"`
	Nickname string `gorm:"type:char(100);DEFAULT '';"`
	Descript string `gorm:"type:varchar(255);DEFAULT '';"`
	Email    string `gorm:"type:varchar(100);DEFAULT '';"`
	Headico  string `gorm:"type:varchar(200);DEFAULT '';"`
}

func (this *User) Login() (User, error) {
	db := libs.DB
	var user User
	if this.Account == "" || this.Password == "" {
		return User{}, errors.New("Cuenta o contraseña está vacía")
	}

	has := md5.Sum([]byte(this.Password))
	md5_password := fmt.Sprintf("%x", has)
	if db.Where(&User{Account: this.Account, Password: md5_password}).First(&user).RecordNotFound() {
		return User{}, errors.New("Cuenta o contraseña incorrecta")
	}
	return user, nil
}
