package main

import (
	"fmt"
	"regexp"
)

func main() {

	text := "Hello World! Helli Patryk!"

	// Check if we got match
	found, _ := regexp.MatchString("([a-z]+)!", text)
	fmt.Printf("found=%v", found)

	// Store regular expression / pattern in variable
	re := regexp.MustCompile("(Hell).")

	first := re.FindString(text)
	fmt.Printf("Leftmost found match: %v\n", first)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	fmt.Printf("Matches: %v\n", matches)

	re, _ = regexp.Compile(`l\w+e`)
	replaced := re.ReplaceAllString("Can life be intåerpreted as like ?", "my bro")
	fmt.Println(replaced)
}
