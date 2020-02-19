package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(strings.Split(scanner.Text(), ","))
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		m := make(map[string]int)
		   for _, word := range words { 
			  m[word] += 1 
		  } 
		  fmt.Print("Word Counts: ")
		  fmt.Println(m)
	   }
}

// func WordCount(scanner string) map[string]int {
// 	// words := scanner.Text()
// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 	}

// m := make(map[string]int)
// for _, word := range words {
// 	m[word] += 1
// }
// 	return count
// }
