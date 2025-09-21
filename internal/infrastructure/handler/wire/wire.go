//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/578223592/awesomeCoupon/internal/application/services"
	"github.com/578223592/awesomeCoupon/internal/infrastructure/repository/mysql"
	"github.com/google/wire"
)

var couponSet = wire.NewSet(
	services.NewCoupon,
	mysql.NewCouponRepository)

//func NewRoleService() role.Service {
//	wire.Build(roleSet)
//
//	return role.Service{}
//}

func NewCoupon() *services.Coupon {
	wire.Build(couponSet)

	return nil
}
