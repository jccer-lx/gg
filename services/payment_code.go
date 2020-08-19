package services

import (
	"fmt"
	"github.com/lvxin0315/gg/databases"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/skip2/go-qrcode"
	"time"
)

//二维码过期时间
const payCodeExpireTime = 10 * 60

//获取用户支付码信息，没有创建，过期->更新
func PayCode(userId int) (*models.PaymentCode, error) {
	paymentCodeModel := new(models.PaymentCode)
	paymentCodeModel.UserId = userId
	databases.NewDB().Model(paymentCodeModel).Order("id DESC").First(paymentCodeModel)
	if paymentCodeModel.ID == 0 {
		//创建
		return CreatePaymentCode(userId)
	}
	//验证是否过期
	if time.Now().Unix()-paymentCodeModel.CreatedAt.Unix() > payCodeExpireTime {
		//创建
		return CreatePaymentCode(userId)
	}
	return paymentCodeModel, nil
}

//创建支付码
func CreatePaymentCode(userId int) (*models.PaymentCode, error) {
	paymentCodeModel := new(models.PaymentCode)
	code := helper.RandString(32)
	fileSavePath := fmt.Sprintf("%s/%s.png", "assets/qrcode", code)
	err := qrcode.WriteFile(code, qrcode.Medium, 256, fileSavePath)
	if err != nil {
		return nil, err
	}
	paymentCodeModel.UserId = userId
	paymentCodeModel.Code = code
	paymentCodeModel.QrCodePath = fileSavePath
	err = databases.NewDB().Save(paymentCodeModel).Error
	if err != nil {
		return nil, err
	}
	return paymentCodeModel, err
}
