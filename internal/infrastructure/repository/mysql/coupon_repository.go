package mysql

import (
	"gorm.io/gorm"
)

type CouponRepository struct {
	client *gorm.DB
}

func NewCouponRepository() *CouponRepository {
	getClient := GetClient()
	return &CouponRepository{
		getClient,
	}
}
