package comic

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var (
	localStoreDirectory string
)

func init() {
	homeDir, _ := os.UserHomeDir()
	localStoreDirectory = filepath.Join(homeDir, ".xkcd")
}

func localStoreExists() bool {
	_, err := os.Stat(localStoreDirectory)
	return err == nil
}

func createLocalStore() error {
	err := os.Mkdir(localStoreDirectory, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(localStoreDirectory + "/curr")
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func getLastComicNum() int {
	return 0
}

func save(bytes []byte) error {
	var c *Comic
	err := json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	c.Save()
	return nil
}
