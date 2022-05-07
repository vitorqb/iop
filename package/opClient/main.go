package opClient

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/vitorqb/iop/package/system"
)

const DEFAULT_CLIENT = "/usr/bin/op"

type OpClient struct {
	sys system.ISystem
	token *string
	path  string
}

func runProxyCmd(cmd *exec.Cmd) ([]byte, error) {
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return stdout.Bytes(), err
}

func New() *OpClient {
	token := ""
	sys := system.New()
	client := OpClient{
		sys: &sys,
		token: &token,
		path:  DEFAULT_CLIENT,
	}
	return &client
}

func (client OpClient) runWithToken(args ...string) ([]byte, error) {
	fullArgs := append([]string{"--session", *client.token}, args...)
	cmd := exec.Command(client.path, fullArgs...)
	return cmd.Output()
}

func (client OpClient) EnsureLoggedIn() {
	cmd := exec.Command(client.path, "signin", "--raw", "--session", *client.token)
	result, err := runProxyCmd(cmd)
	if err != nil {
		client.sys.Crash("Something wen't wrong during signin", err)
	}
	*client.token = string(result)
}

func (client OpClient) GetPassword(itemRef string) string {
	result, err := client.runWithToken("item", "get", itemRef, "--field", "label=password")
	if err != nil {
		client.sys.Crash("Something wen't wrong during item get", err)
	}
	return strings.Trim(string(result), "\n")
}
