package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups! something go wrong \n", "***** Details ******\n", err)
		}
	}()

	if file, err := os.Open("./files/contacts.txt"); err != nil {
		panic("File not found!")
	} else {

		closeFile := func() {
			file.Close()
		}

		defer closeFile()

		content := make([]byte, 254)

		size, _ := file.Read(content)

		fileContent := string(content[0:size])
		fmt.Println(fileContent)
	}
}
