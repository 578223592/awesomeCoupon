package usercontext

import (
	"context"

	"github.com/578223592/awesomeCoupon/internal/domain/entity"
)

type key struct{}

var userKey = &key{}

func WithUser(ctx context.Context, user *entity.UserInfo) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func FromContext(ctx context.Context) *entity.UserInfo {
	if u, ok := ctx.Value(userKey).(*entity.UserInfo); ok {
		return u
	}
	return nil
}

func GetUserID(ctx context.Context) string {
	if u := FromContext(ctx); u != nil {
		return u.UserId
	}
	return ""
}

func GetUsername(ctx context.Context) string {
	if u := FromContext(ctx); u != nil {
		return u.Username
	}
	return ""
}

func GetShopNumber(ctx context.Context) int64 {
	if u := FromContext(ctx); u != nil {
		return u.ShopNumber
	}
	return 0
}
