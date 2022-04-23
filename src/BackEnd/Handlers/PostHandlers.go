package Handlers

import (
	"fmt"
	"net/http"
	"time"

	"KobokDNA.com/BackEnd/DNA"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostPenyakit(c *gin.Context) {
	var Penyakit DNA.PenyakitInput
	// var json = []byte(`{"nama_penyakit":"Penyakit1","keterangan":"Keterangan1","gejala":"Gejala1","solusi":"Solusi1"}`)
	err := c.ShouldBindJSON(&Penyakit)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := "Error on field " + e.Field() + ", condition: " + e.ActualTag()
			errorMessages = append(errorMessages, errorMessage)
		}
	} else {
		// request, _ := http.NewRequest("POST", "http://localhost:8888/Penyakit", bytes.NewBuffer(json))
		// request.Header.Set("Content-Type", "application/json;charset=UTF-8")
		// client := &http.Client{}
		// response, error := client.Do(request)
		// if error != nil {
		// 	fmt.Println(error)
		// }
		// defer response.Body.Close()
		// fmt.Println(response.Status)
		// fmt.Println(response.Header)

	}
}

func PostTestDNA(c *gin.Context) {
	var TestDNA DNA.TestDNAInput

	err := c.ShouldBindJSON(&TestDNA)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		fmt.Println(errorMessages)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
	} else {
		NamaPengguna := TestDNA.NamaPengguna
		DNApengguna := TestDNA.SequenceDNA
		NamaPenyakit := TestDNA.NamaPenyakit
		Method := TestDNA.Method

		errorMessage, err := DNAsequenceValidator(DNApengguna)
		if err {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessage,
			})
		} else {
			now := time.Now()
			dateNow := fmt.Sprintf("%d %s %d", now.Day(), Bulan[int(now.Month())], now.Year())

			// strlain := "world!!"
			diagnosis, similarityScore := TestDNAHandler(DNApengguna, NamaPenyakit, Method)
			c.JSON(http.StatusOK, gin.H{
				"Tanggal":      dateNow,
				"Pengguna":     NamaPengguna,
				"Penyakit":     NamaPenyakit,
				"SkorKesamaan": similarityScore,
				"Diagnosis":    diagnosis,
				"errors":       nil,
			})
		}
		// if (reg1.MatchString(DNApengguna)) {

		// }
		// if (gak sesuai) {
		// 	errorMessage = ""
		// 	c.JSON(http.StatusUnauthorized, errorMessage)
		// }

	}
}

func PostSearching(c *gin.Context) {
	var Searching DNA.SearchingInput

	err := c.ShouldBindJSON(&Searching)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := "Error on field " + e.Field() + ", condition: " + e.ActualTag()
			errorMessages = append(errorMessages, errorMessage)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world!!",
		})
	}
}
