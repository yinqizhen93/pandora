package main

//
//import (
//	"fmt"
//	"github.com/spf13/viper"
//	"log"
//	"pandora/db"
//	"strings"
//)
//
////func BatchInsert(table string, cols []string, data [][]string) error {
////
////	valueStrings := make([]string, 0, len(data))
////	valueArgs := make([]interface{}, 0, len(data)*len(data[0]))
////	header := strings.Join(cols, ",")
////	for _, row := range data {
////		// 此处占位符要与插入值的个数对应
////		var oneRowStr []string
////		for _, c := range row {
////			valueArgs = append(valueArgs, c)
////			oneRowStr = append(oneRowStr, "?")
////		}
////		valueStrings = append(valueStrings, "("+strings.Join(oneRowStr, ",")+")")
////	}
////	fmt.Println("valueStrings", valueStrings)
////	fmt.Println("valueArgs", valueArgs)
////	// 自行拼接要执行的具体语句
////	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", table, header, strings.Join(valueStrings, ","))
////	fmt.Println(stmt)
////	db.Client.ExecQuery(stmt, valueArgs...)
////	return nil
////}
//
//func InitConfig() {
//	viper.SetConfigFile("../config/config.dev.yaml")
//	err := viper.ReadInConfig()
//	if err != nil {
//		fmt.Println("获取配置文件失败")
//		panic(err)
//	}
//}
//
//func main() {
//	InitConfig()
//	db.InitDB()
//	file := "/Users/tison.yin/Desktop/tison/600001.xlsx"
//	c, err := ReadXlsx(file, "600000")
//	fmt.Println(c)
//	if err != nil {
//		log.Fatal(err)
//	}
//	cols := []string{"market", "code", "name", "date", "open", "close", "high", "low", "volume", "outstanding_share", "turnover"}
//	err = BatchInsert("stock", cols, c[1:])
//	if err != nil {
//		log.Println(err)
//	}
//}
