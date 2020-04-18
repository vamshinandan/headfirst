package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	number := rand.Intn(100) + 1
	fmt.Println("Guess the random number: ")
	var guessed bool = false

	for i := 0; i < 4; i++ {
		input := bufio.NewReader(os.Stdin)
		entry, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		entry = strings.TrimSpace(entry)
		guess, err := strconv.ParseInt(entry, 0, 32)
		if err != nil {
			log.Fatal(err)
		}
		if int(guess) == number {
			fmt.Println("You guessed it right at ", i, " attempt")
			guessed = true
			break
		}
	}
	if guessed == false {
		fmt.Println("Could not Guess Properly!!")
	}
}
