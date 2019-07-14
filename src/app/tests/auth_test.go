package test

import (
	"app/models"
	"app/models/auth"
	"app/models/user"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidateLoginModel(t *testing.T) {
	Convey("Subject: Test ValidateLoginModel func\n", t, func() {
		Convey("Validate required fields", func() {
			m := models.Login{Email: "", Password: ""}
			valid := auth.ValidateLoginModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Validate email", func() {
			validEmailLength := 15
			validaPass := "validapass"
			m := models.Login{Email: getRandomString(validEmailLength), Password: validaPass}
			valid := auth.ValidateLoginModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Validate max email length", func() {
			invalidEmailLength := 300
			longString := getRandomString(invalidEmailLength)
			validEmailEnding := "@example.com"
			validaPass := "validpass"
			email := fmt.Sprintf("%s%s", longString, validEmailEnding)
			m := models.Login{Email: email, Password: validaPass}
			valid := auth.ValidateLoginModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Validate max password length", func() {
			invalidPasswordLength := 17
			password := getRandomString(invalidPasswordLength)
			validEmail := "email@example.com"
			m := models.Login{Email: validEmail, Password: password}
			valid := auth.ValidateLoginModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Not existing user", func() {
			validpassword := "validpass"
			validEmail := "email@example.com"
			m := models.Login{Email: validEmail, Password: validpassword}
			valid := auth.ValidateLoginModel(&m)
			So(valid.HasErrors(), ShouldBeTrue)
		})
		Convey("Valid data and existing user", func() {
			or := beginTransaction()

			validpassword := "validpass"
			validEmail := "email@example.com"
			m := models.User{Email: validEmail, Password: validpassword}
			or, err := auth.CreateUser(&m, or)
			So(err, ShouldBeNil)

			lm := models.Login{Email: validEmail, Password: validpassword}
			valid := auth.ValidateLoginModel(&lm)
			So(valid.HasErrors(), ShouldBeTrue)

			rollbackTransaction(or)
		})
	})
}

func TestValidateUserModel(t *testing.T) {
	Convey("Subject: Test ValidateUserModel func\n", t, func() {
		Convey("Validate required fields", func() {
			m := models.User{Email: "", Password: ""}
			valid := auth.ValidateUserModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Validate email", func() {
			validEmailLength := 15
			validaPass := "validapass"
			m := models.User{Email: getRandomString(validEmailLength), Password: validaPass}
			valid := auth.ValidateUserModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Validate max email length", func() {
			invalidEmailLength := 300
			longString := getRandomString(invalidEmailLength)
			validEmailEnding := "@example.com"
			validaPass := "validpass"
			email := fmt.Sprintf("%s%s", longString, validEmailEnding)
			m := models.User{Email: email, Password: validaPass}
			valid := auth.ValidateUserModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Validate max password length", func() {
			invalidPasswordLength := 17
			password := getRandomString(invalidPasswordLength)
			validEmail := "email@example.com"
			m := models.User{Email: validEmail, Password: password}
			valid := auth.ValidateUserModel(&m)
			So(valid.HasErrors(), ShouldNotBeNil)
		})
		Convey("Valid data", func() {
			validpassword := "validpass"
			validEmail := "email@example.com"
			m := models.User{Email: validEmail, Password: validpassword}
			valid := auth.ValidateUserModel(&m)
			fmt.Printf("%v", valid.Errors)
			So(valid.HasErrors(), ShouldBeFalse)
		})
	})
}

func TestCreateUser(t *testing.T) {
	Convey("Subject: Test CreateUser func\n", t, func() {
		or := beginTransaction()

		validpassword := "validpass"
		validEmail := "email@example.com"

		Convey("Dont see user", func() {
			user, _ := user.FindByEmail(validEmail, or)
			So(user.Id, ShouldBeZeroValue)
		})
		Convey("See User", func() {
			m := models.User{Email: validEmail, Password: validpassword}
			or, _ := auth.CreateUser(&m, or)
			user, _ := user.FindByEmail(validEmail, or)
			So(user.Id, ShouldBeGreaterThan, 0)
		})

		rollbackTransaction(or)
	})
}

func TestEncodePassword(t *testing.T) {
	Convey("Subject: Test EncodePassword func\n", t, func() {
		Convey("Create User", func() {
			validpassword := "validpass"
			validEmail := "email@example.com"
			m := models.User{Email: validEmail, Password: validpassword}
			auth.EncodePassword(&m)
			So(m.Password, ShouldNotBeNil)
			So(m.Password, ShouldNotEqual, validpassword)
			So(m.Salt, ShouldNotBeNil)
		})
	})
}
