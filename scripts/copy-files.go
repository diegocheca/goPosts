package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get the source and destination folders.
	sourceFolder := os.Args[1]
	destinationFolder := os.Args[2]

	//srcFolder := "copy/from/path"
	//destFolder := "copy/to/path"
	cpCmd := exec.Command("cp", "-r", sourceFolder, destinationFolder)
	err := cpCmd.Run()
	if err != nil {
		fmt.Print("eerrrrorrrr en cp")
		os.Exit(1)
	}
	fmt.Println(err)
	fmt.Println("Files copied successfully.")
}

// llamada:  ./copy-files source destination

/*

Lista de files a copiar:

Enums - toda la carpeta

tests/TestCase.php

model/Images.php  --> menos en blog x ahora

read.me

*/
