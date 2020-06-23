package models

const (
	JudgeTrue  = true
	JudgeFalse = false
)

//判断题
type JudgmentQuestion struct {
	BaseQuestion
	Judge bool `gorm:"type:TINYINT(2);"`
}

func (q *JudgmentQuestion) TableName() string {
	return "question_judgment"
}

func (q *JudgmentQuestion) GetType() int {
	return Judgment
}
