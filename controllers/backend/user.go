package backend

import (
	"Blog/controllers"
	"Blog/helper"
	"Blog/models"
	"encoding/json"
	"fmt"
)

//  UserController operations for User
type UserController struct {
	BaseController
}

// 用户列表
func (c *UserController) Index() {
	var user *models.User
	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 10)
	users, err := user.UserList(page, size)
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, err}
		c.ServeJSON()
		return
	}
	// 计算总数量
	count, err := user.UserCount()
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, err}
		c.ServeJSON()
		return
	}
	total := count / size
	if count%size != 0 {
		total = total + 1
	}
	results := controllers.Result{
		Item:     users,
		PageItem: controllers.Page{page, size, total},
	}
	c.Data["json"] = controllers.Response{0, "success", results}
	c.ServeJSON()
	return
}

// 获取用户详细信息
func (c *UserController) Show() {
	var user models.User
	id := c.Ctx.Input.Param(":id")
	err := user.GetUserDetailByFiled("id", id)
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, err}
		c.ServeJSON()
		return
	}
	c.Data["json"] = controllers.Response{0, "success", user}
	c.ServeJSON()
	return
}

func (c *UserController) Destroy() {
	var user models.User
	id := c.Ctx.Input.Param(":id")
	err := user.DeleteUserDetailByFiled("id", id)
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, err}
		c.ServeJSON()
		return
	}
	c.Data["json"] = controllers.Response{0, "success", user}
	c.ServeJSON()
	return
}

func (c *UserController) Store() {
	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	password, err := helper.Encrypt(user.Password)
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, "密码加密失败"}
		return
	}
	user.Password = password
	user.CreateUser()
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, err}
		c.ServeJSON()
		return
	}
	c.Data["json"] = controllers.Response{0, "添加用户成功", user}
	c.ServeJSON()
	return
}
func (c *UserController) Update() {
	var user *models.User
	id := c.Ctx.Input.Param(":id")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	err = user.UpdateUserInfo("id", id)
	if err != nil {
		c.Data["json"] = controllers.ErrResponse{-1, err}
		c.ServeJSON()
		return
	}
	user.GetUserDetailByFiled("id", id)
	c.Data["json"] = controllers.Response{0, "更新用户信息成功", user}
	c.ServeJSON()
	return
}
