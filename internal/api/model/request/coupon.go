package request

// CreateCouponTemplateReq 优惠券保存请求结构体
type CreateCouponTemplateReq struct {
	// 优惠券名称
	Name string `json:"name" example:"用户下单满10减3特大优惠" binding:"required"`

	// 优惠券来源 0：店铺券 1：平台券
	Source int `json:"source" example:"0" binding:"required"`

	// 优惠对象 0：商品专属 1：全店通用
	Target int `json:"target" example:"1" binding:"required"`

	// 优惠商品编码
	Goods string `json:"goods,omitempty"`

	// 优惠类型 0：立减券 1：满减券 2：折扣券
	Type int `json:"type" example:"0" binding:"required"`

	// 有效期开始时间
	ValidStartTime string `json:"validStartTime" example:"2024-07-08 12:00:00" binding:"required"`

	// 有效期结束时间
	ValidEndTime string `json:"validEndTime" example:"2025-07-08 12:00:00" binding:"required"`

	// 库存
	Stock int `json:"stock" example:"200" binding:"required"`

	// 领取规则
	ReceiveRule string `json:"receiveRule" example:"{\"limitPerPerson\":1,\"usageInstructions\":\"3\"}" binding:"required"`

	// 消耗规则
	ConsumeRule string `json:"consumeRule" example:"{\"termsOfUse\":10,\"maximumDiscountAmount\":3,\"explanationOfUnmetConditions\":\"3\",\"validityPeriod\":\"48\"}" binding:"required"`
}
