package main

import (
	"fmt"
	"net/http"

	"KobokDNA.com/BackEnd/Script"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})

	})
	foundKMP, firstPosKMP, closedMatchKMP := Script.KnuthMorrisPratt("abacaabaccabacabaabb", "abacabb")
	if foundKMP {
		fmt.Printf("Found match at first index: %d\n", firstPosKMP)
	} else {
		fmt.Printf("No match found. ")
		fmt.Printf("Closed match: %s. ", closedMatchKMP)
		distance := Script.LCSDistance(closedMatchKMP, "abacadb", len(closedMatchKMP), len("abacabb"))
		fmt.Printf("Similarity Value : %d\n", Script.Similarity(len("abacabb"), distance))
	}
	foundBM, firstPosBM, closedMatchBM := Script.BoyerMoore("abacaabaccabacabaabb", "abacabb")
	if foundBM {
		fmt.Printf("Found match at first index: %d\n", firstPosBM)
	} else {
		fmt.Printf("No match found. ")
		fmt.Printf("Closed match: %s. ", closedMatchBM)
		distance := Script.LCSDistance(closedMatchBM, "abacadb", len(closedMatchBM), len("abacabb"))
		fmt.Printf("Similarity Value : %d\n", Script.Similarity(len("abacabb"), distance))
	}
}
