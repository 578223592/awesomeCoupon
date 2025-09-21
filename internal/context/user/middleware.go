package usercontext

import (
	"github.com/578223592/awesomeCoupon/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func Middware() func(g *gin.Context) {
	return func(g *gin.Context) {
		ctx := g.Request.Context()
		// 用户属于非核心功能，这里先通过模拟的形式代替。后续如果需要后管展示，会重构该代码
		userInfoEntity := entity.NewUserInfo("1810518709471555585", "pdd45305558318", 1810714735922956666)
		ctx = WithUser(ctx, userInfoEntity)
		g.Request.WithContext(ctx)
	}
}
