package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b := scanner.Text()
	fmt.Println(suma(a, strings.Split(b, " ")))
}

func suma(a int, arr []string) int {
	result := 0
	fmt.Println(arr)
	for _, value := range arr {
		number, _ := strconv.Atoi(value)
		result += number
	}

	return result
}
