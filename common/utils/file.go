package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(destFilepath string, sourceUrl string) error {
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
