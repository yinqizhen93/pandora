package permission

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"pandora/service"
//)
//
//func GetPolicy(c *gin.Context) {
//	fmt.Println("查看policy")
//	list := service.Enforcer.GetPolicy()
//	for _, vlist := range list {
//		for _, v := range vlist {
//			fmt.Printf("value: %s, ", v)
//		}
//	}
//}
//
//func AddPolicy(c *gin.Context) {
//	fmt.Println("增加Policy")
//	if ok, _ := service.Enforcer.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
//		fmt.Println("Policy已经存在")
//	} else {
//		fmt.Println("增加成功")
//	}
//}
//
//func UpdatePolicy(c *gin.Context) {
//	oldPolicy := []string{"admin", "/api/v1/hello", "GET"}
//	newPolicy := []string{"admin", "/api/v1/hello2", "GET"}
//	if ok, _ := service.Enforcer.UpdatePolicy(oldPolicy, newPolicy); !ok {
//		fmt.Println("更新失败")
//	} else {
//		fmt.Println("更新成功")
//	}
//}
//
//func DeletePolicy(c *gin.Context) {
//	fmt.Println("删除Policy")
//	if ok, _ := service.Enforcer.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
//		fmt.Println("Policy不存在")
//	} else {
//		fmt.Println("删除成功")
//	}
//}
