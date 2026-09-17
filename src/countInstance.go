package main

import (
	"strings"
)

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
	// ?
}


func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	frequencyMap := make(map[rune]map[string]int)
	// var frequencyMap := map[rune]map[string]int{}

	for _, name := range names {
		firstLetter := rune(name[0])
		if _, ok := frequencyMap[firstLetter]; !ok {
			frequencyMap[firstLetter] = make(map[string]int)
			
		}
		
		frequencyMap[firstLetter][name]++	
	}

	return frequencyMap
}

func countDistinctWords(messages []string) int {
	// 
	uniqueWords := make(map[string]struct{}) //could have also user bool but struct{} saves memory. rn presense of key shows that word is found
	for _, message := range messages {
		message = strings.ToLower(message)
		messageDiv := strings.Fields(message) //fields  is used to split, could have user Split but fields also handles multiple spaces and also trims 

		for _, word := range messageDiv {
			uniqueWords[word] = struct{}{} //we user an empty struct to save memory and show that we have found a unique word
		}
	}
	return len(uniqueWords)
}
