package ValidatorInput

import (
	"errors"
	"fmt"
	"regexp"
)

func DNAValidator(DNApengguna string) error {
	// errorMessage := ""
	// errorCount := 0
	// reg1, _ := regexp.Compile("[a-z]")
	// if reg1.MatchString(DNApengguna) {
	// 	errorMessage += "not capital letter"
	// 	errorCount++
	// }

	// reg2, _ := regexp.Compile("[1234567890]+")

	// if reg2.MatchString(DNApengguna) {
	// 	if errorCount > 0 {
	// 		errorMessage += ", "
	// 	}
	// 	errorCount++
	// 	errorMessage += "contains number"
	// 	// c.Error(err)
	// }

	// reg3, _ := regexp.Compile("[^A-Z]")

	// if reg3.MatchString(DNApengguna) {
	// 	if errorCount > 0 {
	// 		errorMessage += ", "
	// 	}
	// 	errorCount++
	// 	errorMessage += "contains illegal character"
	// 	// c.Error(err)
	// }

	reg1, _ := regexp.Compile("[^ACGT]")
	if reg1.MatchString(DNApengguna) {
		return errors.New("DNA sequence contains illegal character")
		// if errorCount > 0 {
		// 	errorMessage += ", "
		// }
		// errorCount++
		// errorMessage += "not in (A, C, G, T)"

	}
	return nil
	// if errorCount == 0 {
	// 	return errorMessage, false
	// } else {
	// 	return "DNA not valid because " + errorMessage, true
	// }
}

func SearchValidator(query string) (int, error) {
	bulan := "(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)"

	regexp1, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	if regexp1.MatchString(query) {
		regexp2, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)$")
		regexp3, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d) \\w")
		if regexp2.MatchString(query) {
			return 1, nil
		} else if regexp3.MatchString(query) {
			return 2, nil
		}
		return 0, errors.New("Query not valid")
	}

	regexp4, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")

	if regexp4.MatchString(query) {
		regexp5, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)$")
		regexp6, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d) \\w")
		if regexp5.MatchString(query) {
			return 3, nil
		} else if regexp6.MatchString(query) {
			return 4, nil
		}
		return 0, errors.New("Query not valid")
	}

	regexp7, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d)")

	if regexp7.MatchString(query) {
		regexp8, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d)$")
		regexp9, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d) \\w")
		if regexp8.MatchString(query) {
			return 5, nil
		} else if regexp9.MatchString(query) {
			return 6, nil
		}
		return 0, errors.New("Query not valid")
	}

	regexp10, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d)", bulan))

	if regexp10.MatchString(query) {
		regexp11, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d)$", bulan))
		regexp12, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d) \\w", bulan))
		if regexp11.MatchString(query) {
			return 7, nil
		} else if regexp12.MatchString(query) {
			return 8, nil
		}
		return 0, errors.New("Query not valid")
	}

	regexp13, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d)", bulan))

	if regexp13.MatchString(query) {
		regexp14, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d)$", bulan))
		regexp15, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d) \\w", bulan))
		if regexp14.MatchString(query) {
			return 9, nil
		} else if regexp15.MatchString(query) {
			return 10, nil
		}
		return 0, errors.New("Query not valid")
	}

	regexp16, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d)", bulan))

	if regexp16.MatchString(query) {
		regexp17, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d)$", bulan))
		regexp18, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d) \\w", bulan))
		if regexp17.MatchString(query) {
			return 11, nil
		} else if regexp18.MatchString(query) {
			return 12, nil
		}
		return 0, errors.New("Query not valid")
	}
	// regexp2, _ := regexp.Compile("(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])/((19|20)\\d\\d)")

	// if regexp2.MatchString(date) {
	// 	return 2, nil

	// }

	// regexp5, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")

	// if regexp5.MatchString(date) {
	// 	return 5, nil
	// }

	// regexp6, _ := regexp.Compile("(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])-((19|20)\\d\\d)")

	// if regexp6.MatchString(date) {
	// 	return 6, nil
	// }

	// regexp9, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d)")

	// if regexp9.MatchString(date) {
	// 	return 9, nil
	// }

	// regexp10, _ := regexp.Compile("(0?[1-9]|1[012]) (0?[1-9]|[12][0-9]|3[01]) ((19|20)\\d\\d)")

	// if regexp10.MatchString(date) {
	// 	return 10, nil
	// }

	// regexp13, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d)", bulan))

	// if regexp13.MatchString(date) {
	// 	return 13, nil
	// }

	// regexp14, _ := regexp.Compile(fmt.Sprintf("%s/(0?[1-9]|[12][0-9]|3[01])/((19|20)\\d\\d)", bulan))

	// if regexp14.MatchString(date) {
	// 	return 14, nil
	// }

	// regexp17, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d)", bulan))

	// if regexp17.MatchString(date) {
	// 	return 17, nil
	// }

	// regexp18, _ := regexp.Compile(fmt.Sprintf("%s-(0?[1-9]|[12][0-9]|3[01])-((19|20)\\d\\d)", bulan))

	// if regexp18.MatchString(date) {
	// 	return 18, nil
	// }

	// regexp21, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d)", bulan))

	// if regexp21.MatchString(date) {
	// 	return 21, nil
	// }

	// regexp22, _ := regexp.Compile(fmt.Sprintf("%s (0?[1-9]|[12][0-9]|3[01]) ((19|20)\\d\\d)", bulan))

	// if regexp22.MatchString(date) {
	// 	return 22, nil
	// }

	return 13, nil
}
