package input

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

func ReadString(promptText string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptText)
	usernameRaw, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Replace(usernameRaw, "\n", "", -1), nil
}

func ReadPassword(promptText string) (string, error) {
	fmt.Print(promptText)
	passwordBytes, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()

	if err != nil {
		return "", err
	}
	return string(passwordBytes), nil
}
