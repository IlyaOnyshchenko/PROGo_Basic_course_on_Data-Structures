package main

import (
	"fmt"
	"strings"
)

func main() {
	var line, line1 string
	fmt.Scan(&line, &line1)
	hashTable1 := make(map[string]int, 37)
	hashTable2 := make(map[string]int, 37)
	hashTable1 = alphabetFill(hashTable1)
	hashTable2 = alphabetFill(hashTable2)
	hashTable1 = hashCountLetters(hashTable1, line)
	hashTable2 = hashCountLetters(hashTable2, line1)
	fmt.Println(compareTwoHash(hashTable1, hashTable2))
}
func alphabetFill(hash map[string]int) map[string]int {
	for i := 'a'; i <= 'z'; i++ {
		hash[string(i)] = 0
	}
	for i := '0'; i <= '9'; i++ {
		hash[string(i)] = 0
	}
	return hash
}
func hashCountLetters(hash map[string]int, word string) map[string]int {
	for _, letter := range word {
		count := strings.Count(word, string(letter))
		hash[string(letter)] = count
	}
	return hash
}
func compareTwoHash(hash1 map[string]int, hash2 map[string]int) string {
	for i := 'a'; i <= 'z'; i++ {
		if hash1[string(i)] == hash2[string(i)] {
			continue
		} else {
			return "NO"
		}
	}
	for i := '0'; i <= '9'; i++ {
		if hash1[string(i)] == hash2[string(i)] {
			continue
		} else {
			return "NO"
		}
	}
	return "YES"
}
