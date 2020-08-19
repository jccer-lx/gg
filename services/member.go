package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/models"
)

//充值
func Recharge(memberId uint, money float64) (*models.Member, error) {
	memberModel := new(models.Member)
	memberModel.ID = memberId
	db := databases.NewDB().Begin()
	defer db.Commit()
	err := db.First(memberModel).Error
	if err != nil {
		db.Rollback()
		return nil, err
	}
	//添加金额
	err = db.Model(memberModel).Update("money", gorm.Expr("money + ?", money)).Error
	if err != nil {
		db.Rollback()
		return nil, err
	}
	//TODO 记录流水
	return memberModel, nil
}
