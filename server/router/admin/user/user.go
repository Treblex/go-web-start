package user

import (
	"EK-Server/model"
	"EK-Server/util"
	"fmt"

	"github.com/labstack/echo"
)

var modelUser model.User

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/user"
	user := g.Group(baseURL)
	user.GET("/list", modelUser.BaseControll.List)
	user.POST("/addUser", modelUser.Add)
	user.POST("/updateUser", modelUser.Update)
	user.POST("/delUser", delUser)
	user.GET("/repeatOfName", repeatOfName)
	user.GET("/repeatOfEmail", repeatOfEmail)
	user.POST("/frozen", frozen)
}

func frozen(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(&u); err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}

	db := model.DB
	row := db.Model(&model.User{BaseControll: model.BaseControll{ID: u.ID}}).Update("status", u.Status)
	if row.Error != nil {
		return util.JSONErr(c, nil, "操作失败")
	}

	if u.Status == 0 {
		return util.JSONSuccess(c, nil, "冻结用户")
	}

	return util.JSONSuccess(c, nil, "解冻用户")
}

func repeatOfEmail(c echo.Context) error {
	user := new(model.User)
	email := c.QueryParam("email")
	user.Email = email
	err := user.HasUser()
	if err != nil {
		return util.JSONSuccess(c, nil, "没有重复")

	}
	return util.JSON(c, nil, "邮箱已被使用,尝试找回密码或者使用其他邮箱", -1)

}

func repeatOfName(c echo.Context) error {
	user := new(model.User)
	name := c.QueryParam("name")
	user.Name = name
	err := user.HasUser()
	if err != nil {
		return util.JSONSuccess(c, nil, "没有重复")
	}
	return util.JSON(c, nil, "用户名已存在", -1002)
}

func delUser(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	result, err := user.DelUser()

	if err != nil {
		return util.JSONErr(c, result, fmt.Sprintf("%s", err))
	}

	return util.JSONSuccess(c, result, "删除成功")
}
