package models

import "github.com/jinzhu/gorm"

//会员支付码
type PaymentCode struct {
	gorm.Model
	UserId     int    //会员id
	Code       string //支付码值
	QrCodePath string //支付二维码
}

func (m *PaymentCode) TableName() string {
	return "zx_payment_code"
}
