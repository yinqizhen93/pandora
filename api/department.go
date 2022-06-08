package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "pandora/db"
	"pandora/logs"
	"pandora/models"
	"pandora/service"
	"runtime/debug"
)

func GetDepartment(c *gin.Context) {
	var departments []models.Department
	db.DB.Find(&departments)
	fmt.Println(departments)
	//fmt.Println(result)
	c.JSON(200, departments)
}

type DepartmentRequest struct {
	Name     string `json:"name" validate:"required"`
	Code     string `json:"code" validate:"required"`
	ParentId int32  `json:"parent_id" validate:"required"`
}

func CreateDepartment(c *gin.Context) {
	var d DepartmentRequest
	if err := c.Bind(&d); err != nil {
		panic(err)
	}

	if err := service.Valid.Struct(d); err != nil {
		logs.Logger.Error(fmt.Sprintf("请求参数有错误：%s; \n %s", err, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    101,
			"msg":     "请求参数有错误",
		})
		return
	}

	department := models.Department{Name: d.Name, Code: d.Code, ParentId: d.ParentId}
	rst := db.DB.Create(&department)
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

func UpdateDepartment(c *gin.Context) {
	var department models.Department
	id := c.Param("id")
	name := c.PostForm("name")
	rst := db.DB.First(&department, id)

	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "更新失败",
		})
		return
	}

	department.Name = name
	if rst := db.DB.Save(&department); rst.Error != nil {
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

func DeleteDepartment(c *gin.Context) {
	var department models.User
	id := c.Param("id")
	rst := db.DB.First(&department, id)

	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "删除失败",
		})
		return
	}

	if rst := db.DB.Delete(&department); rst.Error != nil {
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
