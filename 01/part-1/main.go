package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sort.Ints(first_list)
	sort.Ints(second_list)

	result := 0

	for i := range first_list {
		distance := first_list[i] - second_list[i]

		if distance < 0 {
			result -= distance
		} else {
			result += distance
		}
	}

	fmt.Println(result)
}
