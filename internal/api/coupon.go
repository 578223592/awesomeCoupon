package api

import (
	"net/http"

	"github.com/578223592/awesomeCoupon/internal/api/model/request"
	"github.com/578223592/awesomeCoupon/internal/api/model/response"
	"github.com/578223592/awesomeCoupon/internal/domain/errno"
	"github.com/gin-gonic/gin"
)

type Coupon struct {
}

func NewCoupon() *Coupon {
	return &Coupon{}
}

// CreateCouponTemplate godoc
// @Summary CreateCouponTemplate
// @Schemes
// @Description do CreateCouponTemplate
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /coupon/create [post]
func (c *Coupon) CreateCouponTemplate(gtx *gin.Context) {
	req := request.CreateCouponTemplateReq{}
	if err := gtx.ShouldBindJSON(&req); err != nil {
		commonError := errno.NewCommonError(errno.InvalidParam, "")
		Response(gtx, commonError, nil)
		return
	} //todo 写到这了 2025年09月07日18:36:41
	//check data ：可以用责任链模式依次进行：1.单纯的入参字段判断（比如订单id不能为空）；
	//	2. 数据检查：比如说传入的orderId对应的数据是否存在； 3.依次检查；
	//  依次的一个原则是：性能。比如说先进行字段判空再查询数据库判断数据是否存在。
	//  检查顺序基本上是没有争论的，争论在于：在ddd领域驱动中，这应该放在哪里去实现。
	// 先说结论：放在不同层去实现，比如说：用户接口层判断入参，应用层或中间件判断数据是否存在，领域层判断数据是否符合业务逻辑。
	// 这样做的虽然校验逻辑分散了，但是每一层干的事情都是符合自己职责的，不会混在一起。
	// 有些人会用责任链模式去实现，并在用户接口层去调用，这样的话也可以，不过就需要每一层都单独提供一个校验函数，供责任链去调用了。
	return
}

func Response(gtx *gin.Context, commonError *errno.CommonError, data any) {
	gtx.JSON(http.StatusOK, response.Common{
		Errno:  commonError.Code(),
		Data:   data,
		ErrMsg: commonError.Message(),
	})
}
