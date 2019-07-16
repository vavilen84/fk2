package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

const (
	RoleAdmin = 1
	RoleUser  = 2
)

type User struct {
	Id        int    `orm:"auto"`
	Email     string `orm:"unique"`
	Password  string
	Salt      string
	Role      int
	FirstName string
	LastName  string
}

func FindUserByEmail(o orm.Ormer, email string) (user User, err error) {
	err = o.QueryTable("user").Filter("email", email).One(&user)
	if err != nil {
		beego.Error(err)
	}
	return
}

func FindUserById(o orm.Ormer, id int) (user User, err error) {
	err = o.QueryTable("user").Filter("id", id).One(&user)
	if err != nil {
		beego.Error(err)
	}
	return
}

func ValidateUserModelOnUpdate(o orm.Ormer, m User) *validation.Validation {
	valid := validation.Validation{}

	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")

	oldUser, err := FindUserById(o, m.Id)
	if err != nil {
		beego.Error(err)
	}
	if oldUser.Email == "" {
		err := valid.SetError("email", "User not exists")
		if err != nil {
			beego.Error(err)
		}
	} else {
		if oldUser.Email != m.Email {
			existingUser, _ := FindUserByEmail(o, m.Email)
			if existingUser.Id != 0 {
				err := valid.SetError("email", "Email is already in use")
				if err != nil {
					beego.Error(err)
				}
			}
		}
	}

	if m.Password != "" {
		valid.MaxSize(m.Password, 16, "password")
	}

	valid.Email(m.Email, "email")

	valid.Required(m.FirstName, "first_name")
	valid.MaxSize(m.FirstName, 255, "first_name")

	valid.Required(m.LastName, "last_name")
	valid.MaxSize(m.LastName, 255, "last_name")

	return &valid
}

func ValidateUserModelOnRegister(o orm.Ormer, m User) *validation.Validation {
	valid := validation.Validation{}

	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")

	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")

	valid.Email(m.Email, "email")

	valid.Required(m.FirstName, "first_name")
	valid.MaxSize(m.FirstName, 255, "first_name")

	valid.Required(m.LastName, "last_name")
	valid.MaxSize(m.LastName, 255, "last_name")

	u, err := FindUserByEmail(o, m.Email)
	if err != nil {
		beego.Error(err)
	}
	if u.Id != 0 {
		err := valid.SetError("email", "Email is already in use")
		if err != nil {
			beego.Error(err)
		}
	}
	return &valid
}

func InsertUser(o orm.Ormer, m User) (err error) {
	m.EncodePassword()
	_, err = o.Insert(&m)
	if err != nil {
		beego.Error(err)
	}
	return
}

func UpdateUser(o orm.Ormer, m User) (err error) {
	oldUser, _ := FindUserById(o, m.Id)
	if m.Password != "" {
		m.EncodePassword()
	} else {
		m.Password = oldUser.Password
		m.Salt = oldUser.Salt
	}
	_, err = o.Update(&m)
	if err != nil {
		beego.Error(err)
	}
	return
}

func (m *User) EncodePassword() {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.Salt = salt
}

func init() {
	orm.RegisterModel(new(User))
}
