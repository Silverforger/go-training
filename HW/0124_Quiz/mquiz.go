package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	database := flag.String("csv", "questions.csv", "database")
	questionsCount := flag.Int("n", 10, "number of questions")
	flag.Parse()
	fp, err := os.Open(*database)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if filepath.Ext(strings.TrimSpace(*database)) != ".csv" {
		log.Fatalf("Database should be in .csv format.")
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	rows, _ := reader.ReadAll()
	if len(rows) < *questionsCount {
		log.Fatalf("Database should contain at least 10 questions.")
	}

	var input string
	var index, score int

	for i := 0; i < *questionsCount; i++ {
		rand.Seed(time.Now().UnixNano())
		index = 0 + rand.Intn(*questionsCount-0+1)
		fmt.Printf("Q: %s = ", rows[index][0])
		fmt.Scan(&input)
		if input == rows[index][1] {
			score++
		}
	}
	fmt.Println("You answered ", score, " out of ", *questionsCount, " questions correctly.")
}
