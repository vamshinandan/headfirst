package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fileOpen(fileName string) (*os.File, error) {
	fmt.Println("Opened the file... ", fileName)
	return os.Open(fileName)
}

func fileClose(file *os.File) {
	fmt.Println("Closed the File...")
	file.Close()
}

func getfloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := fileOpen(fileName)
	defer fileClose(file)
	if err != nil {
		log.Fatal("Couldn't Open file... ", fileName)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal("Invalid number in the file!!!")
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func main() {
	numbers, err := getfloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, val := range numbers {
		sum = sum + val
	}
	fmt.Printf("Sum of the data is %.2f\n", sum)
}
