package controllers

import (
	"app/auth"
	"app/models"
	"app/s3"
	"app/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"image"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) setAuthData() {
	token := auth.GetToken(c.Ctx)
	c.Data["IsLoggedIn"] = token.IsLoggedIn
	if token.User != "" {
		var user models.User
		err := json.Unmarshal([]byte(token.User), &user)
		if err != nil {
			beego.Error(err)
		}
		c.Data["User"] = user
		if user.Role == models.RoleAdmin {
			c.Data["IsAdmin"] = true
		} else {
			c.Data["IsAdmin"] = false
		}
	}
}

func (c *BaseController) setResponseData(title, templateName string) {
	c.setRenderData(title, templateName)
	c.setAuthData()
}

func (c *BaseController) setRenderData(title, templateName string) {
	c.Data["title"] = title
	c.Layout = "layout.html"
	c.TplName = templateName + ".html"
}

func (c *BaseController) getImageData(imageFormName string) (imagePath, originalFilename, uuid string) {
	file, header, err := c.GetFile(imageFormName)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	ext, err := utils.GetImageExtension(file)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	i, _, err := image.Decode(file)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	originalFilename = header.Filename
	imagePath, uuid, err = s3.SaveImage(i, ext)
	if err != nil {
		beego.Error(err)
		return "", "", ""
	}
	defer file.Close()
	return
}
