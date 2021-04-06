package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	if node.yesPath != nil && node.noPath != nil {

		scanner := bufio.NewScanner(os.Stdin)

		for {
			scanner.Scan()
			answer := scanner.Text()

			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("Invalid input, please answer yes or no")
			}
		}
	}
}

func (node *storyNode) printTree(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}

	fmt.Println(node.text)

	if node.yesPath != nil {
		node.yesPath.printTree(depth + 1)
	}
	if node.noPath != nil {
		node.noPath.printTree(depth + 1)
	}
}

func main() {
	root := storyNode{"You are at the entrace to a dark woods. Do you enter?", nil, nil}

	winning := storyNode{"you have won!", nil, nil}
	losing := storyNode{"you died", nil, nil}

	root.yesPath = &losing
	root.noPath = &winning

	root.printTree(0)

	root.play()
}
