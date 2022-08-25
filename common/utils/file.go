package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	LOTUS_JSON_RPC_ID      = 7878
	LOTUS_JSON_RPC_VERSION = "2.0"
)

func DownloadFile(sourceUrl string, destFilepath string) error {
	// Create the file
	out, err := os.Create(destFilepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(sourceUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
