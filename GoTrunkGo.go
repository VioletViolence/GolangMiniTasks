package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, _ := strconv.ParseFloat(scanner.Text(), 64)
	var integer = int64(number)
	fmt.Println(integer)
}
