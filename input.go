package main

import (
	"bufio"
	"fmt"
)

func ReadInput(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}

		return "", false
	}

	return scanner.Text(), true
}
