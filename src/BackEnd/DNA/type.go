package DNA

type PenyakitInput struct {
	NamaPenyakit string
	SequenceDNA  string
}

type TestDNAInput struct {
	NamaPengguna string
	SequenceDNA  string
	NamaPenyakit string
	StringMatching int
}

type SearchingInput struct {
	Tanggal      string
	NamaPenyakit string
}
