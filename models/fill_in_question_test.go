package models

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFillInQuestionInsert(t *testing.T) {
	databases.InitMysqlDB()
	Convey("insert 一条填空题", t, func() {
		err := databases.NewDB().Create(&FillInQuestion{
			BaseQuestion: BaseQuestion{
				Stem:       fmt.Sprintf("测试%s填空题%s题干", FillInReplace, FillInReplace),
				Score:      3,
				Answer:     "o1,o2",
				Analysis:   "测试填空题解析",
				CategoryId: 0,
				Difficulty: Level8,
			},
			Answers: []string{"o1", "o2"},
		}).Error
		So(err, ShouldBeNil)
	})
}

func TestFillInQuestionFirst(t *testing.T) {
	//先插入1条
	TestFillInQuestionInsert(t)
	Convey("查询1条", t, func() {
		var q FillInQuestion
		err := databases.NewDB().Order("id DESC").First(&q).Error
		So(err, ShouldBeNil)
		So(q.Answers[0], ShouldEqual, "o1")
		So(q.Answers[1], ShouldEqual, "o2")
	})
}
