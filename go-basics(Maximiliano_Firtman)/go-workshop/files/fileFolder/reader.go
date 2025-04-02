package fileutils

import "os"

func ReadTextFile(filename string) (string, error){
	content, err := os.ReadFile(filename)

	if err != nil {
		// we could not read the file
		return "", err
	} else {
		// Operation successful
		return string(content), nil
	}
}