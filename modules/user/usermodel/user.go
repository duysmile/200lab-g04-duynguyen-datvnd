package usermodel

import (
	"200lab/food-delivery/common"
	"errors"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

type UserCreate struct {
	common.SQLModel `json:",inline`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"-" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserCreate) TableName() string {
	return "users"
}

func (User) TableName() string {
	return "users"
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExited",
	)

	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password is incorrect"),
		"email or password is incorrect",
		"ErrEmailOrPasswordInvalid",
	)
)
