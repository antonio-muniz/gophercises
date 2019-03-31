package stdio

import (
	"fmt"
	"os"
	"os/exec"
)

func AwaitConfirmation() {
	fmt.Println("\nPress enter to continue...")
	RequestWord()
}

func RequestWord() string {
	var word string
	fmt.Scanf("%s\n", &word)
	return word
}

func UpdateText(text string) {
	clearScreen()
	fmt.Print(text)
}

func clearScreen() {
	clearCommand := exec.Command("clear")
	clearCommand.Stdout = os.Stdout
	clearCommand.Run()
}
