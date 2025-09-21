package services

import "github.com/578223592/awesomeCoupon/internal/infrastructure/repository/mysql"

type Coupon struct {
	couponRepository *mysql.CouponRepository
}

func NewCoupon(coupon *mysql.CouponRepository) *Coupon {
	return &Coupon{
		couponRepository: coupon,
	}
}
