package dto

type ResponseParams struct {
	StatusCode int
	Message    string
	Halaman    *Halaman
	Data       any
}
