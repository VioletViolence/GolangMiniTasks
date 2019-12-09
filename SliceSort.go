package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var sli = make([]int, 0, 3)

	for{
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		e, err := strconv.Atoi(scanner.Text())
		if err != nil && scanner.Text() == "x" {
			break
		}else {
			sli = append(sli, e)
		}
		sort.Ints(sli)
		fmt.Print(sli)
	}
}
