package Handlers

func TestDNA(NamaPengguna, DNApengguna, NamaPenyakit, StringMatching) {
	str := []string{"hello", "world"}
	// if (gak sesuai) {
	// 	errorMessage = ""
	// 	c.JSON(http.StatusUnauthorized, errorMessage)
	// }
	if (StringMatching == 1) {
		found, firstPos, closedMatch := Script.KnuthMorrisPratt(DNApengguna, DNApenyakit)
	} else {
		found, firstPos, closedMatch := Script.BoyerMoore(DNApengguna, DNApenyakit)
	}
	
	if (found) {
		strSimilarityScore := "100 %"
	} else {
		distance := StringMatching.LCSDistance(closedMatchKMP, "abacadb", len(closedMatchKMP), len("abacabb"))
		similarityScore := StringMatching.Similarity(len("abacabb"), distance)
		strSimilarityScore := string(similarityScore) + " %"
		if (similarityScore > 80) {
			found = true
		}
	}
}