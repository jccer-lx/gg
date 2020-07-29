package main

import (
	"encoding/json"
	"fmt"
	"github.com/lvxin0315/gg/databases"
	model "github.com/lvxin0315/gg/models"
)

type f struct {
	Field string `json:"field"`
	Title string `json:"title"`
	Sort  string `json:"sort"`
}

func main() {
	databases.InitMysqlDB()
	m := model.Goods{}
	var fList []*f
	for _, item := range databases.NewDB().NewScope(m).Fields() {
		fList = append(fList, &f{
			Field: item.DBName,
			Title: item.Name,
			Sort:  "true",
		})
	}
	b, _ := json.Marshal(fList)
	fmt.Println(string(b))
}
