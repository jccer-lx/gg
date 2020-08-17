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
		//models.QuestionCategory{},
		//models.ChoiceQuestion{},
		//models.MultipleChoiceQuestion{},
		//models.FillInQuestion{},
		//models.JudgmentQuestion{},
		//models.QuestionBank{},
		//models.User{},
		//models.QuestionBankCorrection{},
		models.AnswerRecord{},
	}
	//databases.NewDB().DropTable(tableModels...)
	databases.NewDB().CreateTable(tableModels...)
}