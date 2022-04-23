package DNA

type PenyakitInput struct {
	NamaPenyakit string
	SequenceDNA  string
}

type TestDNAInput struct {
	NamaPengguna string `binding:"required"`
	SequenceDNA  string
	NamaPenyakit string
	Method       int
}

type SearchingInput struct {
	Tanggal      string
	NamaPenyakit string
}
