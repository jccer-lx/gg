package models

import (
	"github.com/lvxin0315/gg/databases"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMultipleChoiceQuestionInsert(t *testing.T) {
	databases.InitMysqlDB()
	Convey("insert 一条多选择题", t, func() {
		err := databases.NewDB().Create(&MultipleChoiceQuestion{
			BaseQuestion: BaseQuestion{
				Stem:       "测试多选择题题干",
				Score:      4,
				Answer:     "o2,o3",
				Analysis:   "测试多选择题解析",
				CategoryId: 1,
				Difficulty: Level6,
			},
			Options: []*ChoiceOption{
				{OptionType: StringOption, Item: "o1"},
				{OptionType: StringOption, Item: "o2"},
				{OptionType: StringOption, Item: "o3"},
				{OptionType: StringOption, Item: "o4"},
			},
			AnswerIndex: "[1,2]",
		}).Error
		So(err, ShouldBeNil)
	})
}

func TestMultipleChoiceQuestionFirst(t *testing.T) {
	//先插入1条
	TestMultipleChoiceQuestionInsert(t)
	Convey("查询1条", t, func() {
		var q MultipleChoiceQuestion
		err := databases.NewDB().Order("id DESC").First(&q).Error
		So(err, ShouldBeNil)
		So(q.Options[0].Item, ShouldEqual, "o1")
		So(q.Options[1].Item, ShouldEqual, "o2")
		So(q.Options[2].Item, ShouldEqual, "o3")
		So(q.Options[3].Item, ShouldEqual, "o4")
	})
}

func TestMultipleChoiceQuestionFind(t *testing.T) {
	//先插入10条
	limit := 10
	for limit > 0 {
		TestMultipleChoiceQuestionInsert(t)
		limit--
	}
	Convey("查询10条", t, func() {
		var qList []*MultipleChoiceQuestion
		err := databases.NewDB().Order("id DESC").Limit(10).Find(&qList).Error
		So(err, ShouldBeNil)
		for _, q := range qList {
			So(q.Options[0].Item, ShouldEqual, "o1")
			So(q.Options[1].Item, ShouldEqual, "o2")
			So(q.Options[2].Item, ShouldEqual, "o3")
			So(q.Options[3].Item, ShouldEqual, "o4")
		}
	})
}
