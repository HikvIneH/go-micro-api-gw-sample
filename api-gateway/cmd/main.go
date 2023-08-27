package main

import (
	"log"
	"net/http"

	"github.com/hikvineh/api-gateway/pkg/auth"
	"github.com/hikvineh/api-gateway/pkg/config"
	"github.com/hikvineh/api-gateway/pkg/order"
	"github.com/hikvineh/api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()

	r.Use(handle500Error)

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}

func handle500Error(c *gin.Context) {
	c.Next()

	// Check if the response has a 5** status code
	statusCode := c.Writer.Status()

	if statusCode >= http.StatusInternalServerError && statusCode < http.StatusInternalServerError+100 {
		// Abort the request and set a 403 status code
		c.AbortWithStatus(http.StatusForbidden)
	}

}
