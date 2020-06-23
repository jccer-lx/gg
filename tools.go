package main

import (
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
)

//工具
func main() {
	databases.InitMysqlDB()
	//model -> mysql
	createTable()
}

func createTable() {
	tableModels := []interface{}{
		models.ChoiceQuestion{},
		models.MultipleChoiceQuestion{},
		models.QuestionCategory{},
	}
	databases.NewDB().DropTable(tableModels...)
	databases.NewDB().CreateTable(tableModels...)
}