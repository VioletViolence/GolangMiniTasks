package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main(){
	type Person struct {
		FirstName string
		SecondName string
	}

	people := make([]Person, 0, 3)

	fmt.Println("Please enter the name of the file")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	abs,err := filepath.Abs("./FileWork.go")
	if err != nil {
		fmt.Println(err)
	}
	dir := path.Dir(abs)
	file, err := os.Open(dir + "/" + scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray := strings.Fields(scanner.Text())
		temp := Person{stringArray[0],stringArray[1]}
		people = append(people, temp)
	}

	for _, user := range people{
		fmt.Println("First Name: " + user.FirstName + " Second Name: " + user.SecondName)
	}
}
