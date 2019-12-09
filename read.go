package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := ""
	addr := ""

	fmt.Println("Enter a name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter an address: ")

	input := bufio.NewReader(os.Stdin)
	addr, err := input.ReadString('\n')

	addr = strings.Trim(addr, "\n")

	inputmap := map[string]string{
		"name": name,
		"address": addr}

	barr, err := json.Marshal(inputmap)

	if err == nil {
		fmt.Println("Error")
	}

	json := string(barr)
	fmt.Println(json)
}