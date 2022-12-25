package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func startQuiz() error {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return err
	}

	var answer string
	totalQuestions := 0
	totalScore := 0

	for _, s := range data {
		totalQuestions += 1
		fmt.Printf("%s?", s[0])
		fmt.Scanln(&answer)
		if answer == s[1] {
			totalScore += 1
		}
	}

	fmt.Printf("You guessed %d/%d", totalScore, totalQuestions)

	return nil
}

func main() {
	startQuiz()
}