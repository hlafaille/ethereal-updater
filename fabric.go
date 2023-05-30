package main

import (
	"os"
	"os/exec"
	"strings"
)

// IsFabricInstalled returns if fabric modloader is installed or not
func IsFabricInstalled() (bool, error) {
	// get the minecraft directory
	minecraftDirectory, err := GetMinecraftDirectory()
	if err != nil {
		return false, err
	}

	// get a list of files in the minecraft versions directory
	mcVersions, err := os.ReadDir(minecraftDirectory + "/versions")
	var fabricProfileFound = false
	for _, version := range mcVersions {
		versionName := version.Name()
		if strings.Contains(versionName, "fabric") && strings.Contains(versionName, "1.19.2") {
			fabricProfileFound = true
		}
	}
	return fabricProfileFound, nil
}

// InstallFabric downloads the Fabric installer for 1.19.2 & runs it
func InstallFabric() error {
	err := DownloadFile("fabricinstaller.jar", "https://maven.fabricmc.net/net/fabricmc/fabric-installer/0.11.2/fabric-installer-0.11.2.jar")
	if err != nil {
		return err
	}

	println("Please select Fabric version for 1.19.2")
	cmd := exec.Command("java", "-jar", "fabricinstaller.jar")
	cmd.Start()
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
