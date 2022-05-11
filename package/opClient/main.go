package opClient

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tokenStorage"
)

const DEFAULT_CLIENT = "/usr/bin/op"

// Runs a command leaving the stdin and stderr of the current execution.
func runProxyCmd(cmd *exec.Cmd) ([]byte, error) {
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return stdout.Bytes(), err
}

// A client for the `op` (1password cli) program
type OpClient struct {
	tokenStorage tokenStorage.ITokenStorage
	sys          system.ISystem
	path  string
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
	cmd := exec.Command(client.path, fullArgs...)
	return cmd.Output()
}

func (client OpClient) listItems() ([]itemListItem, error) {
	opRawResult, err := client.runWithToken("item", "list", "--format", "json")
	if err != nil {
		return []itemListItem{}, err
	}
	var items []itemListItem
	json.Unmarshal(opRawResult, &items)
	return items, nil
}

func (client OpClient) EnsureLoggedIn() {
	token, err := client.getToken()
	if err != nil {
		client.sys.Crash("Something wen't wrong when recovering the token", err)
	}
	cmd := exec.Command(client.path, "signin", "--raw", "--session", token)
	result, err := runProxyCmd(cmd)
	if err != nil {
		client.sys.Crash("Something wen't wrong during signin", err)
	}
	err = client.tokenStorage.Put(string(result))
	if err != nil {
		log.Printf("WARNING: could not save token: %s\n", err)
	}
}

func (client OpClient) GetPassword(itemRef string) string {
	result, err := client.runWithToken("item", "get", itemRef, "--field", "label=password")
	if err != nil {
		client.sys.Crash("Something wen't wrong during item get", err)
	}
	return strings.Trim(string(result), "\n")
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

func New(sys system.ISystem, tokenStorage tokenStorage.ITokenStorage) *OpClient {
	client := OpClient{
		sys:          sys,
		path:         DEFAULT_CLIENT,
		tokenStorage: tokenStorage,
	}
	return &client
}
