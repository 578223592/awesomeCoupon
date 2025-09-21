package entity

type UserInfo struct {

	/**
	 * 用户 ID
	 */
	UserId string
	/**
	 * 用户名
	 */
	Username string
	/**
	 * 店铺编号
	 */
	ShopNumber int64
}

func NewUserInfo(userId string, username string, shopNumber int64) *UserInfo {
	return &UserInfo{UserId: userId, Username: username, ShopNumber: shopNumber}
}
