package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/db"
	"pandora/logs"
	"pandora/models"
	"pandora/service"
	"runtime/debug"
)

func GetUser(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	fmt.Println(users)
	//fmt.Println(result)
	c.JSON(200, users)
}

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func CreateUser(c *gin.Context) {
	var u UserRequest
	if err := c.Bind(&u); err != nil {
		panic(err)
	}

	if err := service.Valid.Struct(u); err != nil {
		logs.Logger.Error(fmt.Sprintf("请求参数有错误：%s; \n %s", err, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    101,
			"msg":     "请求参数有错误",
		})
		return
	}

	user := models.User{UserName: u.Username, Password: u.Password, Email: u.Email}
	rst := db.DB.Create(&user)
	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("插入数据错误:%s; %s", rst.Error, string(debug.Stack())))
		c.JSON(200, gin.H{
			"success": false,
			"code":    104,
			"msg":     "插入数据错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "添加成功",
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	email := c.PostForm("email")
	rst := db.DB.First(&user, id)

	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "更新失败",
		})
		return
	}

	user.Email = email
	if rst := db.DB.Save(&user); rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("保存数据库失败:%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "更新成功",
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	rst := db.DB.First(&user, id)

	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "删除失败",
		})
		return
	}

	if rst := db.DB.Delete(&user); rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("删除失败:%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "删除成功",
	})
}
