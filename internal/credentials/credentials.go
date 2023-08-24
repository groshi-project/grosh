package credentials

import (
	"encoding/json"
	"errors"
	"github.com/Wessie/appdirs"
	"os"
	"path"
)

const credentialsStorageDirPermission = 0700  // owner: read, write, execute; others: no permissions
const credentialsStorageFilePermission = 0600 // owner: read, write; others: no permissions
const credentialsStorageFileName = "credentials.json"
const appName = "grosh"

func GetCredentialsStorageFilePath() string {
	app := &appdirs.App{Name: appName}
	return path.Join(app.UserData(), credentialsStorageFileName)
}

type Credentials struct {
	// URL of the groshi server
	URL string

	// JWT for authorizing at the server
	JWT string
}

func NewCredentials(url string, jwt string) *Credentials {
	return &Credentials{URL: url, JWT: jwt}
}

func NewCredentialsFromFile(filepath string) (*Credentials, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	credentials := Credentials{}
	if err := json.Unmarshal(data, &credentials); err != nil {
		return nil, err
	}

	return &credentials, nil
}

func NewCredentialsFromCredentialsFile() (*Credentials, error) {
	credentialsFilePath := GetCredentialsStorageFilePath()
	return NewCredentialsFromFile(credentialsFilePath)
}

func (c *Credentials) WriteToFile(filepath string) error {
	content, err := json.Marshal(c)
	if err != nil {
		return err
	}

	dirPath := path.Dir(filepath)
	if err := os.Mkdir(dirPath, credentialsStorageDirPermission); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
	}

	if err := os.WriteFile(filepath, content, credentialsStorageFilePermission); err != nil {
		return err
	}

	return nil
}
