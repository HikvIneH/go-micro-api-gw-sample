package routes

import (
	"context"
	"net/http"

	"github.com/hikvineh/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindAll(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.FindAll(context.Background(), &pb.FindAllRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
