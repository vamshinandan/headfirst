package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func readDir(directory string) error {
	fmt.Println(directory)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filepath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			err := readDir(filepath)
			if err != nil {
				fmt.Println("Ended up with error in :", filepath)
			}
		} else {
			fmt.Println(file.Name())
		}
	}
	return nil
}

func main() {
	err := readDir("../../headfirst")
	if err != nil {
		log.Fatal(err)
	}
}
