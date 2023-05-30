package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Version represents a version response from the API
type Version struct {
	Version string `json:"version"`
}

// IsModpackUpdateAvailable checks if the modpack has an update available
func IsModpackUpdateAvailable() (bool, error) {
	// get the users home directory
	userHomeDir, homeDirErr := os.UserHomeDir()
	if homeDirErr != nil {
		return false, homeDirErr
	}

	// get the .ethereal/version file
	etherealVersionText, etherealVersionErr := os.ReadFile(userHomeDir + "/.ethereal/version")
	if etherealVersionErr != nil {
		os.Mkdir(userHomeDir+"/.ethereal", os.ModePerm)
		os.Create(userHomeDir + "/.ethereal/version")
	}

	// read the version text
	version := string(etherealVersionText)

	// get the latest version from the API
	resp, err := http.Get("https://hlafaille.xyz/api/data.json")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// marshal apiVersion
	apiVersion := Version{}
	json.Unmarshal(body, &apiVersion)

	// if the API version does not equal the local version, return true
	if apiVersion.Version != version {
		return true, nil
	}
	return false, nil
}

// InstallModpack downloads and installs the Ethereal modpack
func InstallModpack() error {
	// download the modpack
	err := DownloadFile("ethereal.tar.gz", "https://hlafaille.xyz/dl/ethereal.tar.gz")

	// extract it
	cmd := exec.Command("tar", "-xf", "ethereal.tar.gz")
	cmd.Start()
	err = cmd.Wait()
	if err != nil {
		return err
	}

	// get the minecraft directory
	minecraftDirectory, err := GetMinecraftDirectory()
	if err != nil {
		return err
	}

	// delete the minecraft mods folder
	err = os.RemoveAll(minecraftDirectory + "/mods")
	if err != nil {
		return err
	}

	// read the mods_new folder
	mods, err := os.ReadDir("mods_new")
	if err != nil {
		return err
	}

	// link each file to the minecraft mods directory
	for _, mod := range mods {
		fmt.Printf("Symlinking '%s' to '%s'\n", "mods_new/"+mod.Name(), minecraftDirectory+"/mods/"+mod.Name())
		os.Link("mods_new/"+mod.Name(), minecraftDirectory+"/mods/"+mod.Name())
	}
	return nil
}
