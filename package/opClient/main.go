package opClient

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

type OpClient struct {
	token *string
	path string
}

func crash(errMsg string, err error) {
	log.Fatal(errMsg)
	log.Fatal(err)
	os.Exit(99)
}

func runProxyCmd(cmd *exec.Cmd) ([]byte, error) {
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return stdout.Bytes(), err
}

func New() (*OpClient) {
	token := ""
	client := OpClient{
		token: &token,
		path: "/usr/bin/op",
	}
	return &client
}

func(client OpClient) runWithToken(args ...string) ([]byte, error) {
	fullArgs := append([]string{"--session", *client.token}, args...)
	cmd := exec.Command(client.path, fullArgs...)
	return cmd.Output()
}

func(client OpClient) EnsureLoggedIn() {
	cmd := exec.Command(client.path, "signin", "--raw", "--session", *client.token)
	result, err := runProxyCmd(cmd)
	if err != nil {
		crash("Something wen't wrong during signin", err)
	}
	*client.token = string(result)
}

func(client OpClient) GetPassword(itemRef string) string {
	result, err := client.runWithToken("item", "get", itemRef, "--field", "label=password")
	if err != nil {
		crash("Something wen't wrong during item get", err)
	}
	return strings.Trim(string(result), "\n")
}
