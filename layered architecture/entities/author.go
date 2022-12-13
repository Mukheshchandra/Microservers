package entities

type Author struct {
	AuthorID  int    `json:"authorId"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Dob       string `json:"dob"`
	Penname   string `json:"penName"`
}
