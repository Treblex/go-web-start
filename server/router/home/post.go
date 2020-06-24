package home

import (
	"EK-Server/model"
	"EK-Server/util"
	"strings"

	"github.com/labstack/echo"
)

var modelPost model.Post

// Init 初始化
func post(g *echo.Group) {
	baseURL := "/post"
	post := g.Group(baseURL)

	post.GET("", func(c echo.Context) error {
		posts, err := modelPost.List(c)
		if err != nil {
			return util.JSONErr(c, err, "err")
		}
		return util.JSONSuccess(c, posts, "获取成功")
	})

	post.POST("/add", func(c echo.Context) error {
		empty := func(err interface{}, msg string) error {
			return util.JSONErr(c, err, msg)
		}
		name := c.FormValue("name")
		if strings.Trim(name, " ") == "" {
			return empty(nil, "作者不可以空")
		}
		return util.JSONSuccess(c, name, "添加成功")
	})

}
