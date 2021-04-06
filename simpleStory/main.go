// TODO: add function to add a new page to the story
// TODO: add function to delete a page from the story

package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyPage struct {
	text     string
	nextPage *storyPage
	// prevPage *storyPage
} // This is a linked list

func (page *storyPage) readStory() {
	scanner := bufio.NewScanner(os.Stdin)
	/*
		if page == nil {
			return
		}
		fmt.Println(page.text)
		scanner.Scan()
		page.nextPage.readStory()
		//*/

	for page != nil {
		fmt.Println(page.text)
		scanner.Scan()
		page = page.nextPage
	}
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}

	page.nextPage = &storyPage{text, nil}
}

func main() {
	// Tedious, but basic
	page := storyPage{"this is the start of the story", nil}
	page.addToEnd("this is the middle of the story")
	page.addToEnd("this is the end of the story")
	// page.addAfter("this is a test")

	page.readStory()
}
