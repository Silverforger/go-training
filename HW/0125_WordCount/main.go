package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

// SafeCounter provides a structure that allows for mutual exclusion of data values in go routines.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Increment function will increase the integer value of the word.
func (c *SafeCounter) Increment(word string) {
	c.mu.Lock()
	c.v[word]++
	c.mu.Unlock()
}

// Value function will return the final integer value of the word.
func (c *SafeCounter) Value(word string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[word]
}

// countWords function will increment the integer value of all the words in the array passed through it.
func (c *SafeCounter) countWords(wordsArray []string) {
	for i := 0; i < len(wordsArray); i++ {
		go c.Increment(wordsArray[i])
	}
}

// dedupeAndSort will remove all redundant words and sort the array alphabetically.
func dedupeAndSort(currentArray []string, addArray []string, ch chan []string) {
	currentArray = append(currentArray, addArray...)
	keys := make(map[string]bool)
	cleanArray := []string{}

	for _, entry := range currentArray {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			cleanArray = append(cleanArray, entry)
		}
	}

	sort.Strings(cleanArray)
	ch <- cleanArray
}

// fileOpener function will open the file in the same directory as the executable file, and will return an array of words that have been purged of non-alphanumeric characters.
func fileOpener(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%v %v", err, fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordsArray := make([]string, 0)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		tempArr := strings.Fields(scanner.Text())
		for _, notCleanedWord := range tempArr {
			if strings.ToLower(reg.ReplaceAllString(notCleanedWord, "")) != "" {
				wordsArray = append(wordsArray, strings.ToLower(reg.ReplaceAllString(notCleanedWord, "")))
			}
		}
	}

	return wordsArray
}

// The main function will count the words of each file opened in the fileOpener function, and will deduplicate and sort those files.
// It will then go through another deduplication and sorting for the final array of unique words in the files used, before printing them in "word value" form.
func main() {
	filesInput := os.Args[2:]
	c := SafeCounter{v: make(map[string]int)}
	ch := make(chan []string)
	keys := make(map[string]bool)
	uniqueWords := make([]string, 0)
	finalWords := make([]string, 0)

	for _, fileName := range filesInput {
		go c.countWords(fileOpener(fileName))
		go dedupeAndSort(uniqueWords, fileOpener(fileName), ch)
		uniqueWords = append(uniqueWords, <-ch...)
	}

	for _, entry := range uniqueWords {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			finalWords = append(finalWords, entry)
		}
	}

	sort.Strings(finalWords)
	time.Sleep(time.Second)
	for _, word := range finalWords {
		fmt.Println(word+" ", c.Value(word))
	}
}
