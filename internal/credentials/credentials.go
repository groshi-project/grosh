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
	// URL of the groshi server.
	URL string `json:"url"`

	// Token for authorizing at the server.
	Token string `json:"token"`
}

func New(url string, jwt string) *Credentials {
	return &Credentials{URL: url, Token: jwt}
}

func ReadFromPath(filepath string) (*Credentials, error) {
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

func ReadFromDefaultPath() (*Credentials, error) {
	credentialsFilePath := GetCredentialsStorageFilePath()
	return ReadFromPath(credentialsFilePath)
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
