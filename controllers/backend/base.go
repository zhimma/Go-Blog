package backend

import (
	"Blog/controllers"
	"Blog/helper"
	"strings"
)

type BaseController struct {
	controllers.MainController
	controllerName, actionName string
	res                        controllers.Response
	helper.Token
	helper.Paginator
}


func (c *BaseController) Prepare() {
	controller, action := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controller[0 : len(controller)-10])
	c.actionName = strings.ToLower(action)
	// 验证jwt
	if strings.ToLower(c.controllerName) != "auth" {
		token := c.Ctx.Input.Header("Authorization")
		if status, err := c.ValidateToken(token); status == false {
			c.Ctx.ResponseWriter.WriteHeader(401)
			c.Data["json"] = controllers.ErrResponse{-1, err.Error()}
			c.ServeJSON()
			return
		}
	}
}
