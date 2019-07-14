package auth

import (
	"app/models"
	"app/models/user"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/gbrlsnchs/jwt/v2"
	"log"
	"strconv"
	"time"
)

const (
	tokenName = "AccessToken"
)

type Token struct {
	*jwt.JWT
	IsLoggedIn  bool   `json:"isLoggedIn"`
	CustomField string `json:"customField,omitempty"`
}

func ValidateLoginModel(m *models.Login) *validation.Validation {
	valid := validation.Validation{}

	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")

	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")

	valid.Email(m.Email, "email")
	or := orm.NewOrm()
	user, _ := user.FindByEmail(m.Email, or)
	if user.Id == 0 {
		valid.SetError("email", "User not found")
	} else {
		passwordValid := password.Verify(m.Password, user.Salt, user.Password, nil)
		if passwordValid == false {
			valid.SetError("password", "Password is wrong")
		}
	}

	return &valid
}

func ValidateUserModel(m *models.User) *validation.Validation {
	valid := validation.Validation{}

	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")

	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")

	valid.Email(m.Email, "email")

	return &valid
}
func ValidateUserModelOnRegister(m *models.User, v *validation.Validation) *validation.Validation {
	or := orm.NewOrm()
	user, _ := user.FindByEmail(m.Email, or)
	if user.Id != 0 {
		v.SetError("email", "Email is already in use")
	}

	return v
}

func CreateUser(m *models.User, or orm.Ormer) (orm.Ormer, error) {
	m = EncodePassword(m)
	_, e := or.Insert(m)
	if e != nil {
		log.Fatal(e)
	}

	return or, e
}

func EncodePassword(m *models.User) *models.User {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.Salt = salt

	return m
}

func LoginHandler(u *models.User, Ctx *context.Context) {
	now := time.Now()
	hs256 := jwt.NewHS256("secret")
	jot := &Token{
		JWT: &jwt.JWT{
			Issuer:         "gbrlsnchs",
			Subject:        "someone",
			Audience:       "gophers",
			ExpirationTime: now.Add(24 * 30 * 12 * time.Hour).Unix(),
			NotBefore:      now.Add(30 * time.Minute).Unix(),
			IssuedAt:       now.Unix(),
			ID:             strconv.Itoa(int(u.Id)),
		},
		IsLoggedIn:  true,
		CustomField: "myCustomField",
	}
	jot.SetAlgorithm(hs256)
	jot.SetKeyID("kid")
	payload, err := jwt.Marshal(jot)
	if err != nil {
		log.Printf("token = %s", err.Error())
		return
	}
	token, err := hs256.Sign(payload)
	if err != nil {
		log.Printf("token = %s", err.Error())
		return
	}
	Ctx.SetCookie(tokenName, string(token))
}

func ValidateAuth(Ctx *context.Context) (IsLoggedIn bool, jot Token) {
	now := time.Now()
	hs256 := jwt.NewHS256("secret")
	token := Ctx.GetCookie(tokenName)
	payload, sig, err := jwt.Parse(token)
	if err != nil {
		log.Printf("token = %s", err.Error())
		return
	}
	if err = hs256.Verify(payload, sig); err != nil {
		log.Printf("token = %s", err.Error())
		return
	}
	if err = jwt.Unmarshal(payload, &jot); err != nil {
		log.Printf("token = %s", err.Error())
		return
	}
	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now)
	if err = jot.Validate(iatValidator, expValidator); err != nil {
		switch err {
		case jwt.ErrIatValidation:
			log.Printf("token = %s", "iat error")
		case jwt.ErrExpValidation:
			log.Printf("token = %s", "exp error")
		case jwt.ErrAudValidation:
			log.Printf("token = %s", "aud error")
		}
		return
	}
	IsLoggedIn = true
	return
}

func Logout(Ctx *context.Context) {
	Ctx.SetCookie(tokenName, string(""))
}
