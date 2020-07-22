package main

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
)

func main() {
	databases.InitMysqlDB()
	tableList := []interface{}{
		models.Admin{},
	}

	for i, t := range tableList {
		err := databases.NewDB().CreateTable(t).Error
		if err != nil {
			fmt.Println(i)
			panic(err)
		}
	}
}
