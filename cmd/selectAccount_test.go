package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/system"
)

// Mocks & Helpers
type MockOpClient struct {
	emails []string
}
func(MockOpClient) EnsureLoggedIn() {}
func(MockOpClient) GetPassword(itemRef string) string {
	return ""
}
func(MockOpClient) ListItemTitles() []string {
	return []string{}
}
func(m MockOpClient) ListEmails() ([]string, error) {
	return m.emails, nil;
}

type MockEmailStorage struct {
	current string
}
func(m *MockEmailStorage) Put(email string) error {
	m.current = email
	return nil
}
func(m *MockEmailStorage) Get() (string, error) {
	return m.current, nil
}


// Tests
func TestUserSelectsAnEmail(t *testing.T) {
	emails := []string{"a@b.c"}
	opClient := MockOpClient{ emails: emails }
	system := system.NewMock()
	emailStorage := MockEmailStorage{}
	run(opClient, &system, &emailStorage)
	storedEmail, err := emailStorage.Get()
	assert.Nil(t, err)
	assert.Equal(t, "a@b.c", storedEmail)
}
