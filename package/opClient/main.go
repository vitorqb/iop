package opClient

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/vitorqb/iop/package/accountStorage"
	"github.com/vitorqb/iop/package/opClient/commandRunner"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tokenStorage"
)

const DEFAULT_CLIENT = "/usr/bin/op"

type IOpClient interface {
	EnsureLoggedIn()
	GetPassword(itemRef string) string
	ListItemTitles() []string
	ListAccounts() ([]string, error)
}

// A client for the `op` (1password cli) program
type OpClient struct {
	tokenStorage tokenStorage.ITokenStorage
	sys           system.ISystem
	path          string
	accountStorage  accountStorage.IAccountStorage
	commandRunner commandRunner.ICommandRunner
}

func (client OpClient) getToken() (string, error) {
	token, err := client.tokenStorage.Get()
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (client OpClient) runWithToken(args ...string) ([]byte, error) {
	token, err := client.getToken()
	if err != nil {
		return nil, err
	}
	fullArgs := append([]string{"--session", token}, args...)
	return client.commandRunner.Run(client.path, fullArgs...)
}

func (client OpClient) listItems() ([]itemListItem, error) {
	opRawResult, err := client.runWithToken("item", "list", "--format", "json")
	if err != nil {
		return []itemListItem{}, err
	}
	var items []itemListItem
	err = json.Unmarshal(opRawResult, &items)
	return items, err
}

func (client OpClient) getCurrentAccount() (string, error) {
	account, err := client.accountStorage.Get()
	if err != nil {
		return "", err
	}
	return account, nil
}

func (client OpClient) isLoggedIn() (bool, error) {
	token, err := client.getToken()
	if err != nil {
		client.sys.Crash("Something wen't wrong when recovering the token", err)
	}
	account, err := client.getCurrentAccount()
	if err != nil {
		client.sys.Crash("Something wen't wrong when recovering the account", err)
	}	
	_, err = client.commandRunner.Run(client.path, "whoami", "--session", token, "--account", account)

	// Whoami returns no error -> we are logged in
	if err == nil {
		return true, nil
	}

	// Whoami returns error 1 -> logged out
	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() == 1 {
			return false, nil
		}
	}

	// Unknown error
	return false, err
}

func (client OpClient) EnsureLoggedIn() {
	token, err := client.getToken()
	if err != nil {
		client.sys.Crash("Something wen't wrong when recovering the token", err)
	}
	account, err := client.getCurrentAccount()
	if err != nil {
		client.sys.Crash("Something wen't wrong when recovering the account", err)
	}
	isLoggedIn, err := client.isLoggedIn()
	if err != nil {
		client.sys.Crash("Failed to determine if was logged in", err)
	}
	if isLoggedIn {
		return
	}
	result, err := client.commandRunner.RunAsProxy(client.path, "signin", "--raw", "--session",  token, "--account", account)
	if err != nil {
		client.sys.Crash("Something wen't wrong during signin", err)
	}
	err = client.tokenStorage.Put(string(result))
	if err != nil {
		log.Printf("WARNING: could not save token: %s\n", err)
	}
}

func (client OpClient) GetPassword(itemRef string) string {
	result, err := client.runWithToken("item", "get", itemRef, "--field", "label=password", "--format", "json")
	if err != nil {
		client.sys.Crash("Something wen't wrong during item get", err)
	}
	var field passwordField
	err = json.Unmarshal(result, &field)
	if err != nil {
		client.sys.Crash("Something wen't wrong when parsing the json response for the password", err)
	}
	return field.Value
}

func (client OpClient) ListItemTitles() []string {
	var items, err = client.listItems()
	if err != nil {
		client.sys.Crash("Something went wrong recovering the list of items", err)
	}
	var result []string
	for _, item := range items {
		result = append(result, item.Title)
	}
	return result
}

func (client OpClient) ListAccounts() ([]string, error) {
	output, err := exec.Command(client.path, "account", "list", "--format=json").Output()
	if err != nil || string(output) == "" {
		return []string{}, err
	}
	var accountListItems []accountListItem
	err = json.Unmarshal(output, &accountListItems)
	var accounts []string
	for _, accountListItem := range accountListItems {
		accounts = append(accounts, accountListItem.Shorthand)
	}
	return accounts, err;
}

func New(
	sys system.ISystem,
	tokenStorage tokenStorage.ITokenStorage,
	accountStorage accountStorage.IAccountStorage,
	commandRunner commandRunner.ICommandRunner,
) *OpClient {
	client := OpClient{
		sys:          sys,
		path:         DEFAULT_CLIENT,
		tokenStorage: tokenStorage,
		accountStorage: accountStorage,
		commandRunner: commandRunner,
	}
	return &client
}
