package handler

import (
	"entgo.io/ent/dialect/sql/sqlgraph"
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/api"
	"runtime/debug"
)

type DepartmentRequest struct {
	Code     string `json:"code" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ParentId int    `json:"parent_id" binding:"required"`
}

func (h Handler) CreateDepartment(c *gin.Context) {
	var dr DepartmentRequest
	if err := c.ShouldBind(&dr); err != nil {
		h.logger.Error(fmt.Sprintf("参数错误\"：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1001, "参数错误", err))
		return
	}

	//if err := service.Valid.Struct(ur); err != nil {
	//	logger.Error(fmt.Sprintf("请求参数有错误：%s; \n %s", err, debug.Stack()))
	//	c.JSON(200, api.FailResponse(1001, "请求参数错误"))
	//	return
	//}
	ctx := c.Request.Context()
	u, err := h.db.Department.Create().
		SetCode("code").
		SetName("name").
		SetParentID(0).
		SetCreatedBy(api.CurrentUserId(c)).
		SetUpdatedBy(api.CurrentUserId(c)).
		Save(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("创建用户失败:%s; %s", err, string(debug.Stack())))
		errMsg := "添加失败"
		if sqlgraph.IsUniqueConstraintError(err) {
			errMsg = "存在重复"
		}
		c.JSON(200, api.FailResponse(1002, errMsg, err))
		return
	}
	_, err = h.db.Department.Create().
		SetCode(dr.Code).
		SetName(dr.Name).
		SetParentID(dr.ParentId).
		SetCreatedBy(api.CurrentUserId(c)).
		SetUpdatedBy(api.CurrentUserId(c)).
		SetParent1(u).
		Save(ctx)
	if err != nil {
		h.logger.Error(fmt.Sprintf("创建用户失败:%s; %s", err, string(debug.Stack())))
		errMsg := "添加失败"
		if sqlgraph.IsUniqueConstraintError(err) {
			errMsg = "存在重复"
		}
		c.JSON(200, api.FailResponse(1002, errMsg, err))
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "添加成功",
		"id":      u.ID,
	})
}
