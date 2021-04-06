package main

// TODO
// Add NPCs - talk, fight, etc
// NPC move on their own
// Items/Inventory to interact with the dungeon
// Accept natural language for input
//

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type choice struct {
	cmd      string
	desc     string
	nextNode *storyNode
}

type storyNode struct {
	text    string
	choices []*choices
}

var scanner *bufio.Scanner

func clear() {
	cmd := exec.Command("clear") // macOS and Linux functional
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (node *storyNode) addChoice(command string, description string, next *storyNode) {
	choice := choices{cmd: command, desc: description, nextNode: next}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)

	if node.choices != nil {
		for _, choice := range node.chocies {
			fmt.Println(choice.command, choice.description)
		}
	}

}

func (node *storyNode) execCmd(cmd string) *storyNode {
	for _, choice := range node.chocies {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			clear()
			return choice.nextNode
		}
		choice = choice.nextChoice
	}

	fmt.Println("Sorry, not sure what you wanted. Try again \n\n\n")
	return node
}

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		if strings.ToLower(scanner.Text()) == "q" {
			os.Exit(0)
		}
		node.execCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	root := storyNode{text: `
	You are in the middle of a field.
	You are carrying only a torch.
	To the North you see a house.
	To the South is a village.
	To the East is a dark cave.
	To the West is a wooded area.
	`}

	village := storyNode{text: `
	There are people bustling about a busy town square
	`}

	blacksmith := storyNode{text: `
	The blacksmith offers to sell you a sword.
	`}

	blacksmithBroke := storyNode{text: `
	Unfortunately you do not have any gold to buy the sword.
	`}
	/*
		blacksmithPurchase := storyNode{text: `
		You purchase the sword. It's a bit heavy in your arm but plenty sharp.
		`}
	*/
	innkeeper := storyNode{text: `
	Would you like to rent a room for the night?
	`}

	innkeeperBroke := storyNode{text: `
	I'm sorry, it appears you don't have any money to rent a room.
	Please come back when you have some money.
	`}
	/*
		innkeeperPurchase := storyNode{text: `
		You purchase a room for the night.
		You awake the next morning feeling refreshed.
		`}
	*/
	darkCave := storyNode{text: `
	You approach the entrance to the cave.
	Inside you can hear a low rumbling noise but 
	you are unable to see.
	You hear a crunching noise underfoot as you walk 
	further into the cave.
	`}

	darkCave2 := storyNode{text: `
	You find a wall with your hand and continue deeper into 
	the cave.
	The crunching under your feet is more frequent and the 
	rumbling is growing louder.
	`}

	darkCaveLit := storyNode{text: `
	You light your torch and look around the cave. 
	You can see that there are bones, possibly human, all around.
	The cave continues on further than the light will reach.
	`}

	darkCave2Lit := storyNode{text: `
	There are more bones everywhere and the 
	rumbling is growing louder.
	`}

	woods := storyNode{text: `
	You walk into the woods, you can hear the birds chirping 
	and the rustling of the wind in the leaves. 
	`}

	woodsTrail := storyNode{text: `
	You continue down the path in the peaceful spring air.
	You see a glint of something shiny ahead near the base 
	of a tree.
	`}

	treasureChest := storyNode{text: `
	You get closer and see what appears to be a small box tucked 
	into the base of the tree and hidden by bushes. You uncover 
	the box and find it locked.
	`}

	house := storyNode{text: `
	You approach the house.
	As you get closer you can smell something delicious cooking.
	`}

	houseDoor := storyNode{text: `
	You walk up the path to door and knock.
	You can hear some noise inside before the door opens.
	You are greeted by an old man holding a gun.
	He shoots you before you can even react.
	You have died.
	`}

	houseWindow := storyNode{text: `
	You creep up to the window and take a peak inside.
	You see someone inside working in the kitchen chopping meat.
	`}

	houseBack := storyNode{text: `
	You creep around to the back of the house.
	There is a tree stump and an axe used to split wood.
	Near the stump is a chair with a small table.
	`}

	houseTable := storyNode{text: `
	You approach the chair and table and notice a large key on the table.
	`}

	root.addChoice("N", "Go North", &house)
	root.addChoice("W", "Go West", &woods)
	root.addChoice("E", "Go East", &darkCave)
	root.addChoice("S", "Go South", &village)
	root.addChoice("Q", "Quit", nil)

	darkCave.addChoice("t", "Light the way with your torch", &darkCaveLit)
	darkCave.addChoice("c", "Continue in the dark", &darkCave2)
	darkCave.addChoice("f", "Back to the field", &root)

	darkCaveLit.addChoice("c", "Continue into the cave", &darkCave2Lit)
	darkCaveLit.addChoice("f", "Back to the field", &root)

	darkCave2.addChoice("b", "Turn Back", &darkCave)
	darkCave2Lit.addChoice("b", "Turn Back", &darkCaveLit)

	woods.addChoice("c", "Continue down the trail", &woodsTrail)
	woods.addChoice("f", "Back to the field", &root)

	woodsTrail.addChoice("s", "Inspect the shiny thing", &treasureChest)
	woodsTrail.addChoice("f", "Back to field", &root)

	treasureChest.addChoice("f", "Back to the field", &root)

	house.addChoice("k", "Knock on the front door.", &houseDoor)
	house.addChoice("s", "Sneak around to the back.", &houseBack)
	house.addChoice("p", "Peak in the window", &houseWindow)
	house.addChoice("f", "Back to the field", &root)

	houseBack.addChoice("t", "Approach the table and chair", &houseTable)
	houseBack.addChoice("f", "Back to the field", &root)

	houseWindow.addChoice("k", "Knock on the front door", &houseDoor)
	houseWindow.addChoice("s", "Sneak around to the back of the house", &houseBack)
	houseWindow.addChoice("f", "Back to the field", &root)

	// houseTable.addChoice("t", "Take the key", &houseTableKey)
	houseTable.addChoice("f", "Back to the field", &root)
	/*
		houseTableKey.addChoice("c", "Continue", &houseBack)
		houseTableKey.addChoice("d", "Return to the front door", &houseDoor)
		houseTableKey.addChoice("f", "Back to the field", &root)
	*/
	village.addChoice("b", "Visit the blacksmith", &blacksmith)
	village.addChoice("r", "Visit the innkeeper and rent a room", &innkeeper)
	village.addChoice("f", "Back to the field", &root)

	blacksmith.addChoice("p", "Purchase the sword", &blacksmithBroke)
	blacksmith.addChoice("v", "Back to the village", &village)

	blacksmithBroke.addChoice("v", "Back to the village", &village)

	innkeeper.addChoice("r", "Rent a room", &innkeeperBroke)
	innkeeper.addChoice("v", "Back to the village", &village)

	innkeeperBroke.addChoice("v", "Back to the village", &village)

	root.play()

}
