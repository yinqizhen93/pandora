package auth

import (
	"context"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/api"
	"pandora/db"
	"pandora/ent/user"
	"pandora/logs"
	"pandora/service"
	"runtime/debug"
	"strconv"
)

func GetUser(c *gin.Context) {
	ctx := context.Background()
	cols := []string{"id", "username", "email"}
	users, err := db.Client.User.Query().Select(cols...).All(ctx)
	if err != nil {
		c.JSON(200, api.FailResponse(2002, "获取用户失败"))
	}
	fmt.Println(users)
	c.JSON(200, users)
}

func GetCurrentUser(c *gin.Context) {
	//var users ent.Users
	ctx := context.Background()
	curUserId, ok := c.Get("userId")
	if !ok {
		c.JSON(200, api.FailResponse(2005, "用户不存在"))
	}
	users, err := db.Client.User.Query().Where(user.IDEQ(curUserId.(int))).Select().All(ctx)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("获取用户失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(2006, "获取用户失败"))
	}
	//db.DB.Where("id = ?", curUserId).First(&users)
	fmt.Println(users)
	//fmt.Println(result)
	c.JSON(200, gin.H{
		"success": true,
		"data":    users[0],
	})
}

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func CreateUser(c *gin.Context) {
	var ur UserRequest
	if err := c.Bind(&ur); err != nil {
		logs.Logger.Error(fmt.Sprintf("请求参数解析失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1001, "请求参数解析失败"))
	}

	if err := service.Valid.Struct(ur); err != nil {
		logs.Logger.Error(fmt.Sprintf("请求参数有错误：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1001, "请求参数错误"))
		return
	}
	ctx := context.Background()
	u, err := db.Client.User.Create().
		SetUsername(ur.Username).
		SetPassword(ur.Password).
		SetEmail(ur.Email).
		Save(ctx)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("创建用户失败:%s; %s", err, string(debug.Stack())))
		errMsg := "添加失败"
		if sqlgraph.IsUniqueConstraintError(err) {
			errMsg = "存在重复"
		}
		c.JSON(200, api.FailResponse(1002, errMsg))
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "添加成功",
		"id":      u.ID,
	})
}

func UpdateUser(c *gin.Context) {
	strId := c.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(200, api.FailResponse(1002, "参数错误"))
		logs.Logger.Error(fmt.Sprintf("参数错误:%s; %s", err, string(debug.Stack())))
		return
	}
	ur, err := api.ParseJsonFormInputMap(c)
	fmt.Println("ur", ur)
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("请求参数解析失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1001, "请求参数解析失败"))
		return
	}
	upd := db.Client.User.UpdateOneID(intId)
	if username, ok := ur["username"]; ok {
		upd.SetUsername(username.(string))
	}
	if password, ok := ur["password"]; ok {
		upd.SetUsername(password.(string))
	}
	if email, ok := ur["email"]; ok {
		upd.SetUsername(email.(string))
	}
	if _, err := upd.Save(context.Background()); err != nil {
		logs.Logger.Error(fmt.Sprintf("更新保存失败：%s; \n %s", err, debug.Stack()))
		c.JSON(200, api.FailResponse(1005, "更新保存失败"))
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "更新成功",
	})
}

func DeleteUser(c *gin.Context) {
	strId := c.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(200, api.FailResponse(1002, "参数错误"))
		logs.Logger.Error(fmt.Sprintf("参数错误:%s; %s", err, string(debug.Stack())))
		return
	}
	if err := db.Client.User.DeleteOneID(intId).Exec(context.Background()); err != nil {
		c.JSON(200, api.FailResponse(1002, "删除失败"))
		logs.Logger.Error(fmt.Sprintf("删除失败:%s; %s", err, string(debug.Stack())))
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"msg":     "删除成功",
	})
}
