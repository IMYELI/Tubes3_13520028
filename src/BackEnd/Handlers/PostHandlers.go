package Handlers

import (
	"fmt"
	"time"
	"regexp"
	"KobokDNA.com/BackEnd/DNA"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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
			errorMessage := "Error on field " + e.Field() + ", condition: " + e.ActualTag() 
			errorMessages = append(errorMessages, errorMessage)
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorMessages)
	} else {
		NamaPengguna := TestDNA.NamaPengguna
		DNApengguna := TestDNA.SequenceDNA
		NamaPenyakit := TestDNA.NamaPenyakit
		Method := TestDNA.StringMatching
		
		reg1, _ := regexp.Compile("[a-z]")
		if (reg1.MatchString(DNApengguna)) {
			errorMessage := "Sequence DNA harus huruf kapital"
			// c.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"errors": errorMessage,
			})
			return
		}

		reg2, _ := regexp.Compile("[^ACGT]")
		if (reg2.MatchString(DNApengguna)) {
			errorMessage := "Sequence DNA tidak valid"
			// c.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"errors": errorMessage,
			})
			return
		}

		// if (reg1.MatchString(DNApengguna)) {

		// }
		// if (gak sesuai) {
		// 	errorMessage = ""
		// 	c.JSON(http.StatusUnauthorized, errorMessage)
		// }

		now := time.Now()
		dateNow := fmt.Sprintf("%d %s %d", now.Day(), Bulan[int(now.Month())], now.Year()) 

		// strlain := "world!!"
		diagnosis, similarityScore := TestDNAHandler(DNApengguna, NamaPenyakit, Method)
		c.JSON(http.StatusOK, gin.H {
			"Tanggal" : dateNow,
			"Pengguna" : NamaPengguna,
			"Penyakit" : NamaPenyakit,
			"SkorKesamaan" : similarityScore, 
			"Diagnosis" : diagnosis,
			"errors": nil,
		})
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
		c.JSON(http.StatusOK, gin.H {
			"hello": "world!!",
		})
	}
}