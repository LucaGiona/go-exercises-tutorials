package main

import (
	"fmt"
	"os"

	"lucagiona.guru/go/files/data"
	fileutils "lucagiona.guru/go/files/fileFolder"
)

func main(){
	rootpath, _ := os.Getwd()
	filepath := rootpath + "/data/text.txt"

	c, err := fileutils.ReadTextFile(filepath)
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Originals  %v\n Double the Original: %v%v", c,c,c) 
		fileutils.WriteToFile(filepath + ".output.txt", newContent)
	} else {
		fmt.Printf("ERROR Panic!, %v", err)
	}

	fmt.Println("Starte test...")

	data.Test()
	data.LocationTest()
}