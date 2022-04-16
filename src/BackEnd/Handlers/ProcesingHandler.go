package Handlers

import  (
	"KobokDNA.com/BackEnd/StringMatching"
)

func TestDNAHandler(DNApengguna string, NamaPenyakit string, Method int) (bool, int) {
			// fmt.Println("Berhasil", TestDNA.NamaPengguna, TestDNA.SequenceDNA, TestDNA.NamaPenyakit)
	// cari di database
	found := false
	var closedMatch string;
	similarityScore := 0;

	DNApenyakit := "ATGATG"
	if (Method == 1) {
		found, closedMatch = StringMatching.KnuthMorrisPratt(DNApengguna, DNApenyakit)
	} else {
		found, closedMatch = StringMatching.BoyerMoore(DNApengguna, DNApenyakit)
	}
	
	if (found) {
		similarityScore = 100
	} else {
		distance := StringMatching.LCSDistance(closedMatch, DNApenyakit, len(closedMatch), len(DNApenyakit))
		similarityScore = StringMatching.Similarity(len(DNApenyakit), distance)
		if (similarityScore > 80) {
			found = true
		}
	}
	return found, similarityScore
}