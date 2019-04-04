package backend

import "C"
import (
	"Blog/controllers"
	"Blog/helper"
	"Blog/models"
	"encoding/json"
	"fmt"
	"time"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Login() {
	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	status, message := user.CheckAuth()
	if status == false {
		c.Data["json"] = controllers.ErrResponse{-1, message}
		c.ServeJSON()
		return
	}

	t := helper.Token{
		Account: user.Account,
		Id:      user.Id,
		Expires: time.Now().Unix() + 360000,
	}
	token, err := t.GetToken()
	if token == "" || err != nil {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = controllers.ErrResponse{-1, err}
	} else {

		user.Token = token
		c.Data["json"] = controllers.Response{0, "success.", user}
	}

	c.ServeJSON()
	return
}
