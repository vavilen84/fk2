package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "Get:Index")
	beego.Router("/post/create", &controllers.PostController{}, "Get:Create")
	beego.Router("/post/save", &controllers.PostController{}, "Post:Save")
	beego.Router("/post/update", &controllers.PostController{}, "Post:Update")
	beego.Router("/post/delete", &controllers.PostController{}, "Post:Delete")
	beego.Router("/post/edit/:id", &controllers.PostController{}, "Get:Edit")
	beego.Router("/auth/login", &controllers.AuthController{}, "*:Login")
	beego.Router("/auth/logout", &controllers.AuthController{}, "Get:Logout")
	beego.Router("/auth/register", &controllers.AuthController{}, "*:Register")

	beego.Router("/profile/update", &controllers.ProfileController{}, "Get:Update")
	beego.Router("/profile/save", &controllers.ProfileController{}, "Post:Save")
}
