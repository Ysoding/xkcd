package comic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Downloader struct {
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) Download() error {
	start := 1
	if localStoreExists() {
		start = getLastComicNum()
	} else {
		err := createLocalStore()
		if err != nil {
			return err
		}
	}

	end := getCurrentComicNum()
	for i := start; i <= end; i++ {
		if i > 10 {
			return nil
		}
		bytes, err := fetch(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))
		if err != nil {
			fmt.Printf("Get %d error: %v\n", i, err)
			continue
		}

		err = save(bytes)
		if err != nil {
			fmt.Printf("Save %d error: %v\n", i, err)
			continue
		}
	}

	return nil
}

func getCurrentComicNum() int {
	bytes, err := fetch("https://xkcd.com/info.0.json")
	if err != nil {
		panic(fmt.Sprintf("get current comic num error:%v", err))
	}

	var c Comic
	_ = json.Unmarshal(bytes, &c)
	return c.Num
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
