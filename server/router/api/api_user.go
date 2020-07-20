package api

import (
	"EK-Server/model"

	"github.com/labstack/echo"
)

var modelUser model.User

// 用户API
func user(g *echo.Group) {
	modelUser.BaseControll.Model = &modelUser
	modelUser.Install(g, "/users")
}
