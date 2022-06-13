package auth

//func GetRole(c *gin.Context) {
//	c.String(200, "Get User")
//}
//
//type RoleRequest struct {
//	Name string `json:"name" validate:"required"`
//}
//
//func CreateRole(c *gin.Context) {
//	var r RoleRequest
//	if err := c.Bind(&r); err != nil {
//		panic(err)
//	}
//
//	if err := service.Valid.Struct(r); err != nil {
//		logs.Logger.Error(fmt.Sprintf("请求参数有错误：%s; \n %s", err, debug.Stack()))
//		c.JSON(200, gin.H{
//			"success": false,
//			"code":    101,
//			"msg":     "请求参数有错误",
//		})
//		return
//	}
//
//	role := models.Role{Name: r.Name}
//	rst := db.DB.Create(&role)
//	if rst.Error != nil {
//		logs.Logger.Error(fmt.Sprintf("插入数据错误:%s; %s", rst.Error, string(debug.Stack())))
//		c.JSON(200, gin.H{
//			"success": false,
//			"code":    104,
//			"msg":     "插入数据错误",
//		})
//		return
//	}
//	c.JSON(200, gin.H{
//		"success": true,
//		"msg":     "添加成功",
//	})
//}
//
//func UpdateRole(c *gin.Context) {
//
//	c.String(200, "Update User")
//}
//
//func DeleteRole(c *gin.Context) {
//	c.String(200, "Delete User")
//}
