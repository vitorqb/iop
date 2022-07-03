package opClient

type itemListItem struct {
	Title string `json:"title"`
}

type accountListItem struct {
	Email string `json:"email"`
	Shorthand string `json:"shorthand"`
}

type passwordField struct {
	Value string `json:"value"`
}
