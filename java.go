package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// IsJavaInstalled returns true/false if Java is installed on the current system
func IsJavaInstalled() (bool, error) {
	out, err := exec.Command("java", "--version").Output()
	if err != nil {
		return false, err
	}
	if !strings.Contains(string(out), "17") {
		return false, nil
	}
	return true, nil
}

// InstallJava downloads the Java 17 installer for the platform, installs it, then deletes the installer executable
func InstallJava() error {
	// determine what URL we should download from
	var url string = ""
	var filename string = ""
	if runtime.GOOS == "windows" {
		url = "https://download.oracle.com/java/17/archive/jdk-17.0.7_windows-x64_bin.exe"
		filename = "java17.exe"
	} else if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "arm64" {
			url = "https://download.oracle.com/java/17/archive/jdk-17.0.7_macos-aarch64_bin.dmg"
		} else {
			url = "https://download.oracle.com/java/17/archive/jdk-17.0.7_macos-x64_bin.dmg"
		}
		filename = "java17.dmg"
	} else if runtime.GOOS == "linux" {
		return fmt.Errorf("Install OpenJDK 17 from package manager")
	} else {
		return fmt.Errorf("Unrecognized paltform")
	}

	// download the file
	DownloadFile(filename, url)

	// begin the installation
	if runtime.GOOS == "windows" {
		cmd := exec.Command(".\\java17.exe")
		cmd.Start()
		err := cmd.Wait()
		if err != nil {
			return err
		}
	} else if runtime.GOOS == "darwin" {
		// mount the dmg
		cmd := exec.Command("hdiutil", "attach", "java17.dmg")
		cmd.Start()
		err := cmd.Wait()
		if err != nil {
			return err
		}

		// run the macOS installer script
		cmd = exec.Command("installer", "-package", "'/Volumes/JDK 17.0.7/JDK 17.0.7.pkg'", "-target", "/")
		cmd.Start()
		err = cmd.Wait()
		if err != nil {
			return err
		}

		// unmount the dmg
		cmd = exec.Command("hdiutil", "detach", "'/Volumes/JDK 17.0.7'")
		cmd.Start()
		err = cmd.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}
