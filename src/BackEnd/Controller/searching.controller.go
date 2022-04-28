package Controller

import (
	"net/http"

	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func GetAllPrediction(ctx *gin.Context) {
	query := ctx.Query("query")

	// fmt.Println(date_in_database)
	// fmt.Println("Type : ", types)

	results, err := Service.AllPrediction(&query)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"predictions": results,
	})
}

// func searchingDNARoutes(rg *gin.RouterGroup) {
// 	disease_route := rg.Group("/searching")
// 	disease_route.GET("/predictionResult", GetTestResult)
// }
