package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language" // This is our vulnerable dependency
)

func main() {
	fmt.Println("Vulnerable Go Program Started")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <language_tag>")
		fmt.Println("Example: go run . en-US")
		fmt.Println("Try a specially crafted long tag to trigger vulnerability (e.g., repeating 'a-' many times)")
		return
	}

	inputTag := os.Args[1]
	fmt.Printf("Attempting to parse language tag: %s\n", inputTag)

	// In older versions of golang.org/x/text (e.g., v0.3.5),
	// a specially crafted long input to language.Parse can cause a panic due to
	// an out-of-bounds read (GO-2021-0113).
	// We are intentionally using an older, vulnerable version for demonstration.
	tag, err := language.Parse(inputTag)
	if err != nil {
		fmt.Printf("Error parsing language tag: %v\n", err)
		// This is where a panic might occur with a vulnerable version and specific input
	} else {
		fmt.Printf("Successfully parsed language tag: %s\n", tag.String())
	}

	fmt.Println("Vulnerable Go Program Finished")
}
