package comic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Comic struct {
	Num        int
	Link       string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Year       string
	Month      string
	Day        string
}

func (c *Comic) Save() error {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	path := filepath.Join(localStoreDirectory, fmt.Sprintf("%d.json", c.Num))

	return os.WriteFile(path, bytes, 0644)
}
