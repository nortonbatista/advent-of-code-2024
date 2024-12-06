package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	first_list := make([]int, 0)
	second_list := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for count := 0; scanner.Scan(); count += 1 {
		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		if count%2 == 0 {
			first_list = append(first_list, number)
		} else {
			second_list = append(second_list, number)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0

	for _, first_list_value := range first_list {
		times := 0

		for _, second_list_value := range second_list {
			if first_list_value == second_list_value {
				times += 1
			}
		}

		result += first_list_value * times
	}

	fmt.Println(result)
}
