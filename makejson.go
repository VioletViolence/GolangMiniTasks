package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	name string
	address string
}

func main()  {
	var contacts = map[string]*Contact{}

	for i := 0; i<2;i++{
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		newName := scanner.Text()
		scanner.Scan()
		newAddress := scanner.Text()
		scanner.Scan()
		index := scanner.Text()

		contacts[index] = &Contact{name: newName, address: newAddress}
	}

	b , err := json.Marshal(contacts)
	if err != nil{
		fmt.Print(err)
	}
	fmt.Println(string(b))
}


