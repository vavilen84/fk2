package test

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSeePage(t *testing.T) {
	pages := []string{
		"/",
		"/post/create",
		"/auth/register",
	}
	for _, v := range pages {
		r, _ := http.NewRequest("GET", v, nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		Convey("Subject: See Page\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	}
}
