package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("Set AOC day folder:")

	var folderName string

	fmt.Scanln(&folderName)

	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Printf("Directory '%s' created successfully.\n", folderName)

	files := []string{"main.go", "sample.txt", "input.txt"}
	for _, file := range files {
		path := folderName + string(os.PathSeparator) + file
		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file, err)
			continue
		}

		if file == "main.go" {
			_, err := f.WriteString("package main\n\nfunc main() {\n\t\n}\n")
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", file, err)
			}
		}
		f.Close()
		fmt.Printf("File '%s' created successfully.\n", path)
	}
}
