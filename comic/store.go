package comic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	f, err := os.Create(getLastComicNumFilepath())
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func getLastComicNum() int {
	bytes, err := os.ReadFile(getLastComicNumFilepath())
	if err != nil {
		panic(err)
	}

	strContent := strings.TrimSpace(string(bytes))

	number, _ := strconv.Atoi(strContent)
	return number
}

func save(bytes []byte) error {
	var c *Comic
	err := json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	path := filepath.Join(localStoreDirectory, fmt.Sprintf("%d.json", c.Num))

	return os.WriteFile(path, bytes, 0644)
}

func getLastComicNumFilepath() string {
	return filepath.Join(localStoreDirectory, "lastNum")
}
