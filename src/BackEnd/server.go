package main

import (
	// "fmt"
	// "KobokDNA.com/BackEnd/Script"
	// "net/http"
	"KobokDNA.com/BackEnd/Handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// var Bulan map[int]string = make(map[int]string { 1 : "Januari",
// 	2 : "Februari",
// 	3 : "Maret",
// 	4 : "April",
// 	5 : "Mei",
// 	6 : "Juni",
// 	7 : "Juli",
// 	8 : "Agustus",
// 	9 : "September",
// 	10 : "Oktober",
// 	11 : "November",
// 	12 : "Desember",
// }


func main() {
	router := gin.Default()
	router.Use(cors.Default())


	router.POST("/Penyakit", Handlers.PostPenyakit)
	router.POST("/TestDNA", Handlers.PostTestDNA)
	router.POST("/Searching", Handlers.PostSearching)
	// router.GET("/Penyakit", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"resp": "dontol!!",
	// 	})

	// })
	router.Run(":8888")
	// foundKMP, firstPosKMP, closedMatchKMP := Script.KnuthMorrisPratt("abacaabaccabacabaabb", "abacabb")
	// if foundKMP {
	// 	fmt.Printf("Found match at first index: %d\n", firstPosKMP)
	// } else {
	// 	fmt.Printf("No match found. ")
	// 	fmt.Printf("Closed match: %s. ", closedMatchKMP)
	// 	distance := Script.LCSDistance(closedMatchKMP, "abacadb", len(closedMatchKMP), len("abacabb"))
	// 	fmt.Printf("Similarity Value : %d\n", Script.Similarity(len("abacabb"), distance))
	// }
	// foundBM, firstPosBM, closedMatchBM := Script.BoyerMoore("abacaabaccabacabaabb", "abacabb")
	// if foundBM {
	// 	fmt.Printf("Found match at first index: %d\n", firstPosBM)
	// } else {
	// 	fmt.Printf("No match found. ")
	// 	fmt.Printf("Closed match: %s. ", closedMatchBM)
	// 	distance := Script.LCSDistance(closedMatchBM, "abacadb", len(closedMatchBM), len("abacabb"))
	// 	fmt.Printf("Similarity Value : %d\n", Script.Similarity(len("abacabb"), distance))
	// }
}


