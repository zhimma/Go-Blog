package routers

import (
	"Blog/controllers/backend"
	"github.com/astaxie/beego"
)

func init() {
	backend := beego.NewNamespace("backend",
		beego.NSRouter("/auth/login", &backend.AuthController{}, "post:Login"),
		// beego.NSRouter("/auth/register", &backend.AuthController{}, "post:Register"),
		beego.NSRouter("/users", &backend.UserController{}, "get:Index"),
		beego.NSRouter("/users/:id", &backend.UserController{}, "get:Show"),
		beego.NSRouter("/users/:id", &backend.UserController{}, "delete:Destroy"),
		beego.NSRouter("/users", &backend.UserController{}, "post:Store"),
		beego.NSRouter("/users/:id", &backend.UserController{}, "put:Update"),
	)
	beego.AddNamespace(backend)
}
