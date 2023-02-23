package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	)

func main(){

	csvFilename := flag.String("csv", "problems.csv", "filepath to problems csv, expects format 'problem,answer'")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("failed to read file %s", *csvFilename))
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exit("failed to get problems")
	}

	parsed := parseLines(lines)

	correctAns := 0

	for i, prob := range parsed {
		fmt.Printf("%d. %s = ", i, prob.problem)
		var userAns string
		fmt.Scan(&userAns)
		if (userAns == prob.answer) {
			fmt.Println("Correct!")
			correctAns += 1
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("Final Score: %d / %d" ,correctAns, len(parsed))

}

func parseLines(lines[][]string) []problem {
	parsed := make([]problem, len(lines))
	for i, probText := range lines {
		parsed[i] = problem{
			problem : probText[0],
			answer : probText[1],
		}
	}
	return parsed
}

type problem struct{
	problem string
	answer string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}