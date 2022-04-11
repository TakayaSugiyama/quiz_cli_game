package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record[0])
		scanner.Scan()
	}
}