package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootPages(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
