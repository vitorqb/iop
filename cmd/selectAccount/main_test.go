package selectAccount

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/pmwrap/package/system"
)

// Mocks & Helpers
type MockOpClient struct {
	accounts []string
}
func(MockOpClient) EnsureLoggedIn() {}
func(MockOpClient) GetPassword(itemRef string) string {
	return ""
}
func(MockOpClient) ListItemTitles() []string {
	return []string{}
}
func(m MockOpClient) ListAccounts() ([]string, error) {
	return m.accounts, nil;
}

type MockAccountStorage struct {
	current string
}
func(m *MockAccountStorage) Put(account string) error {
	m.current = account
	return nil
}
func(m *MockAccountStorage) Get() (string, error) {
	return m.current, nil
}


// Tests
func TestUserSelectsAnAccount(t *testing.T) {
	accounts := []string{"a@b.c"}
	opClient := MockOpClient{ accounts: accounts }
	system := system.NewMock()
	accountStorage := MockAccountStorage{}
	run(opClient, &system, &accountStorage)
	storedAccount, err := accountStorage.Get()
	assert.Nil(t, err)
	assert.Equal(t, "a@b.c", storedAccount)
}
