package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	safeReportsCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		report := strings.Fields(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		isIncreasing := false
		first, _ := strconv.Atoi(report[0])
		second, _ := strconv.Atoi(report[1])
		diff := first - second

		if diff == 0 || diff > 3 || diff < -3 {
			continue
		} else if diff > 0 {
			isIncreasing = true
		}

		for i := 2; i < len(report); i += 1 {

			first, _ := strconv.Atoi(report[0])
			second, _ := strconv.Atoi(report[1])
			diff := first - second

			if diff == 0 || diff > 3 || diff < -3 {
				continue
			} else if diff > 0 {
				isIncreasing = true
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// result := 0
	//
	// for i := range first_list {
	// 	distance := first_list[i] - second_list[i]
	//
	// 	if distance < 0 {
	// 		result -= distance
	// 	} else {
	// 		result += distance
	// 	}
	// }
	//
	// fmt.Println(result)
}
