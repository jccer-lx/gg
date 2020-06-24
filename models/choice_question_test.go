package models

import (
	"github.com/lvxin0315/gg/databases"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestChoiceQuestionInsert(t *testing.T) {
	databases.InitMysqlDB()
	Convey("insert 一条选择题", t, func() {
		err := databases.NewDB().Create(&ChoiceQuestion{
			BaseQuestion: BaseQuestion{
				Stem:       "测试选择题题干",
				Score:      2,
				Answer:     "o1",
				Analysis:   "测试选择题解析",
				CategoryId: 0,
				Difficulty: Level4,
			},
			Options: []*ChoiceOption{
				{OptionType: StringOption, Item: "o1"},
				{OptionType: StringOption, Item: "o2"},
				{OptionType: StringOption, Item: "o3"},
				{OptionType: StringOption, Item: "o4"},
			},
			AnswerIndex: 0,
		}).Error
		So(err, ShouldBeNil)
	})
}

func TestChoiceQuestionFirst(t *testing.T) {
	//先插入1条
	TestChoiceQuestionInsert(t)
	Convey("查询1条", t, func() {
		var q ChoiceQuestion
		err := databases.NewDB().Order("id DESC").First(&q).Error
		So(err, ShouldBeNil)
		So(q.Options[0].Item, ShouldEqual, "o1")
		So(q.Options[1].Item, ShouldEqual, "o2")
		So(q.Options[2].Item, ShouldEqual, "o3")
		So(q.Options[3].Item, ShouldEqual, "o4")
	})
}

func TestChoiceQuestionFind(t *testing.T) {
	//先插入10条
	limit := 10
	for limit > 0 {
		TestChoiceQuestionInsert(t)
		limit--
	}
	Convey("查询10条", t, func() {
		var qList []*ChoiceQuestion
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
