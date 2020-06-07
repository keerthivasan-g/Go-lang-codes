package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0, 3)

	for {
		var numString string

		fmt.Println("Add number to array : ")
		fmt.Scan(&numString)


		numString = strings.ToLower(numString)

		if ((numString == "x")||(numString == "X") {
      break
		}

		num, err := strconv.Atoi(numString)

		if err != nil {
			fmt.Println("Wrong input")
      break
		}

		s = append(s, num)

		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})


	}
  fmt.Println(s)

}
