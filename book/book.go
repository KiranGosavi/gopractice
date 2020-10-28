
package book

//Book struct to represent book data
type Book struct{
	ID int `json:"id"`
	BookName string `json:"name"`
	Writers []string `json:"writers"`
	//ReleaseDate time.Time `json:"release_at"`
	CopiesAvailable int `json:"copies_available"`
}
