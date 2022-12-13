package entities

type Book struct {
	BookID        int    `json:"bookId"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	Publication   string `json:"publication"`
	PublishedDate string `json:"publishedDate"`
}
