package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/db"
	"pandora/logs"
	"pandora/models"
	"pandora/service"
	"runtime/debug"
)

func GetDoctor(c *gin.Context) {
	var doctors []models.Doctor
	db.DB.Find(&doctors)
	fmt.Println(doctors)
	//fmt.Println(result)
	c.JSON(200, doctors)
}

type DoctorRequest struct {
	Name   string `json:"name" validate:"required"`
	Code   string `json:"code" validate:"required"`
	DeptId int32  `json:"dept_id" validate:"required"`
}

func CreateDoctor(c *gin.Context) {
	var d DoctorRequest
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

	doctor := models.Doctor{Name: d.Name, Code: d.Code, DeptId: d.DeptId}
	fmt.Println("doctor", doctor)
	rst := db.DB.Create(&doctor)
	//rst := db.DB.Omit("Department").Create(&doctor) // 忽略关联的Department表
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

func UpdateDoctor(c *gin.Context) {
	var d DoctorRequest
	var doctor models.Doctor
	id := c.Param("id")
	if err := c.Bind(&d); err != nil {
		panic(err)
	}
	rst := db.DB.First(&doctor, id)

	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "更新失败",
		})
		return
	}

	doctor.Name = d.Name
	doctor.Code = d.Code
	doctor.DeptId = d.DeptId
	if rst := db.DB.Save(&doctor); rst.Error != nil {
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

func DeleteDoctor(c *gin.Context) {
	var doctor models.User
	id := c.Param("id")
	rst := db.DB.First(&doctor, id)

	if rst.Error != nil {
		logs.Logger.Error(fmt.Sprintf("%s; %s", rst.Error, debug.Stack()))
		c.JSON(200, gin.H{
			"success": false,
			"code":    111,
			"msg":     "删除失败",
		})
		return
	}

	if rst := db.DB.Delete(&doctor); rst.Error != nil {
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
