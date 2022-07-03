package clients

type IClient interface {
	EnsureLoggedIn()
	GetPassword(itemRef string) string
	ListItemTitles() []string
	ListAccounts() ([]string, error)
}
