package test

import (
	"app/models"
	"app/models/auth"
	"app/models/user"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindByEmail(t *testing.T) {
	or := beginTransaction()

	Convey("Subject: FindByEmail func\n", t, func() {
		Convey("User is not found", func() {
			user, _ := user.FindByEmail("not@existing.email", or)
			So(user.Id, ShouldBeZeroValue)
		})
		Convey("Create new User", func() {
			m := models.User{Email: "user@example.com", Password: "password"}
			_, err := auth.CreateUser(&m, or)
			So(err, ShouldEqual, nil)
		})
		Convey("Find created new User", func() {
			user, _ := user.FindByEmail("user@example.com", or)
			So(user.Id, ShouldBeGreaterThan, 0)
		})
	})

	rollbackTransaction(or)
}
