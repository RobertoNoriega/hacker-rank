package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scoreAlice, scoreBob int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pointsAlice := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	pointsBob := strings.Split(scanner.Text(), " ")
	if len(pointsAlice) == len(pointsBob) {
		var i = 0
		for i < len(pointsAlice) {
			alice, err := strconv.Atoi(pointsAlice[i])
			if err != nil {
				fmt.Println(err)
				break
			}
			bob, err := strconv.Atoi(pointsBob[i])
			if err != nil {
				fmt.Println(err)
				break
			}
			if alice > bob {
				scoreAlice = scoreAlice + 1
			} else if alice < bob {
				scoreBob = scoreBob + 1
			}
			i = i + 1
		}
	}
	fmt.Println(scoreAlice, scoreBob)
}
