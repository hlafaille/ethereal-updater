package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

// GetMinecraftDirectory returns the path to the .minecraft/minecraft directory for the current system
func GetMinecraftDirectory() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		return userHomeDir + "/AppData/Roaming/.minecraft", nil
	} else if runtime.GOOS == "darwin" {
		return userHomeDir + "/Library/Application Support/minecraft", nil
	} else if runtime.GOOS == "linux" {
		return userHomeDir + "/.minecraft", nil
	}
	return "", fmt.Errorf("Unsupported platform")
}

// DownloadFile downloads a file to filepath from url
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
