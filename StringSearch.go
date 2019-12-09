package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var sli = make([]int, 3)
	var x string
	for i := 0; ; i++{
		fmt.Scan(&x)
		if x == "x" {break
		}else{
			if i < 3 {
				c, _ := strconv.Atoi(x)
				sli[i] = c
				if(i == 2){
					sort.Ints(sli)
				}
				fmt.Println(sli)
			}else {
				i, _ := strconv.Atoi(x)
				sli = append(sli, i)
				sort.Ints(sli)
				fmt.Println(sli)
			}
		}
	}
}