package DNA

type PenyakitInput struct {
	NamaPenyakit string
	SequenceDNA  string
}

type TestDNAInput struct {
	NamaPengguna string
	SequenceDNA  string
	NamaPenyakit string
}

type SearchingInput struct {
	Tanggal      string
	NamaPenyakit string
}
