package main

import (
	"flag"
	"fmt"
	"os"
	)

func main(){

	csvFilename := flag.String("csv", "problems.csv", "filepath to problems csv, expects format 'problem,answer'")

	flag.Parse()

	fmt.Println(*csvFilename)

	_, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("failed to read file %s", csvFilename))
	}
}

type problem struct{
	problem string
	answer string
}

// func getProblems(filename) [][] string{
// 	res = make()
// }

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}