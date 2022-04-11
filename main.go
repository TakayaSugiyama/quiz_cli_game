package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)
	problemsCount := 0
	correctCount := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record[0])
		problemsCount++
		scanner.Scan()
		input := scanner.Text()
		ans := strings.ReplaceAll(record[1], " ", "")

		if input == ans {
			correctCount++
		}

	}
	fmt.Printf("%d問中、%d問正解しました.\n", problemsCount, correctCount)
}
