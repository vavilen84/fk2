package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "Get:Index")
	beego.Router("/404", &controllers.MainController{}, "Get:PageNotFound")

	beego.Router("/post/create", &controllers.PostController{}, "Get:Create")
	beego.Router("/post/save", &controllers.PostController{}, "Post:Save")
	beego.Router("/post/update", &controllers.PostController{}, "Post:Update")
	beego.Router("/post/delete", &controllers.PostController{}, "Post:Delete")
	beego.Router("/post/edit", &controllers.PostController{}, "Get:Edit")
	beego.Router("/post/view", &controllers.PostController{}, "Get:View")
	beego.Router("/post/edit-list", &controllers.PostController{}, "Get:EditList")

	beego.Router("/auth/login", &controllers.AuthController{}, "*:Login")
	beego.Router("/auth/logout", &controllers.AuthController{}, "Get:Logout")
	beego.Router("/auth/register", &controllers.AuthController{}, "*:Register")

	beego.Router("/profile/update", &controllers.ProfileController{}, "Get:Update")
	beego.Router("/profile/save", &controllers.ProfileController{}, "Post:Save")

	beego.Router("/gallery/student", &controllers.GalleryController{}, "Get:Student")
	beego.Router("/gallery/graduate", &controllers.GalleryController{}, "Get:Graduate")
	beego.Router("/gallery/user", &controllers.GalleryController{}, "Get:User")
	beego.Router("/gallery/update-portfolio", &controllers.GalleryController{}, "Get:UpdatePortfolio")
	beego.Router("/gallery/add-image-to-portfolio", &controllers.GalleryController{}, "Post:AddImageToPortfolio")

}
