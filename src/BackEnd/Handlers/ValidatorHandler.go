package Handlers

import (
	"regexp"
)

func DNAsequenceValidator(DNApengguna string) (string, bool) {
	errorMessage := ""
	errorCount := 0
	reg1, _ := regexp.Compile("[a-z]")
	if reg1.MatchString(DNApengguna) {
		errorMessage += "huruf kapital"
		errorCount++
	}

	reg2, _ := regexp.Compile("[1234567890]+")

	if reg2.MatchString(DNApengguna) {
		if errorCount > 0 {
			errorMessage += ", "
		}
		errorCount++
		errorMessage += "tidak mengandung angka"
		// c.Error(err)
	}

	reg3, _ := regexp.Compile("[^A-Z]")

	if reg3.MatchString(DNApengguna) {
		if errorCount > 0 {
			errorMessage += ", "
		}
		errorCount++
		errorMessage += "tidak mengandung ilegal karakter"
		// c.Error(err)
	}

	reg4, _ := regexp.Compile("[^ACGT]")
	if reg4.MatchString(DNApengguna) {
		if errorCount > 0 {
			errorMessage += ", "
		}
		errorCount++
		errorMessage += "hanya dalam rangkaian A C G T"

	}
	if errorCount == 0 {
		return errorMessage, false
	} else {
		return "Sequence DNA harus " + errorMessage, true
	}
}
