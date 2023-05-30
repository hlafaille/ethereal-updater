package main

import (
	"log"
)

func main() {
	println(" ---===[ ETHEREAL UPDATER]===--- ")
	// check if Java 17 is installed
	javaInstalled, _ := IsJavaInstalled()

	// if Java 17 is not installed
	if !javaInstalled {
		println("Java 17 not installed")
		err := InstallJava()
		if err != nil {
			log.Fatal(err)
		}
	}

	// check if fabric modloader is installed
	fabricInstalled, err := IsFabricInstalled()
	if err != nil {
		log.Fatal(err)
	}
	if !fabricInstalled {
		println("Fabric not installed")
		InstallFabric()
	}

	// check if the modpack is up to date
	modpackUpdateAvailable, err := IsModpackUpdateAvailable()
	if err != nil {
		log.Fatal(err)
	}
	if modpackUpdateAvailable {
		println("Modpack update available")
		InstallModpack()
		println("You may now launch Minecraft!")
	} else {
		println("Modpack is already up to date")
	}
}
