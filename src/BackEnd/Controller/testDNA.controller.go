package Controller

import (
	"net/http"
	"strconv"

	"KobokDNA.com/Models"
	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func TestDNAController(ctx *gin.Context) {
	var TestData Models.TestData

	if err := ctx.ShouldBindJSON(&TestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Semua field harus diisi"})
		return
	}

	ID, err := Service.TestDNA(TestData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success adding a result",
		"id":      ID,
	})
}

func GetTestResult(ctx *gin.Context) {
	ID, errors := strconv.Atoi(ctx.Query("id"))

	if errors != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": errors.Error()})
		return
	}

	result, err := Service.TestResult(&ID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date":             result.Date,
		"user_name":        result.UserName,
		"disease_name":     result.DiseaseName,
		"similarity_score": result.SimilarityScore,
		"status":           result.Status,
	})
}
