package main

import (
	"flag"
	"fmt"
	"net/http"

	docs "github.com/578223592/awesomeCoupon/data/interface_docs"
	"github.com/578223592/awesomeCoupon/internal/api"
	usercontext "github.com/578223592/awesomeCoupon/internal/context/user"
	"github.com/578223592/awesomeCoupon/internal/infrastructure/handler/config"
	"github.com/578223592/awesomeCoupon/internal/infrastructure/redis"
	"github.com/578223592/awesomeCoupon/internal/infrastructure/repository/mysql"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

//swag init --output ./data/interface_docs -g ./main.go

func init() {
	var confPath string
	flag.StringVar(&confPath, "c", "./conf/config.dev.toml", "-c set config file path") // default config file is conf/app.toml
	flag.Parse()
	fmt.Printf("confPath is %s\n ", confPath)
	config.Init(confPath)
	mysql.Init()
	redis.Init()
}
func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode) //default is debug mode

	{

		couponGroup := r.Group("/coupon")
		couponGroup.Use(usercontext.Middware())
		coupon := api.NewCoupon()
		couponGroup.POST("/create", coupon.CreateCouponTemplate)
	}
	r.GET("/ping", Ping)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	fmt.Println("swagger docs: http://localhost:8080/swagger/index.html")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
