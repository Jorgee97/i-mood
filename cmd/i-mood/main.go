package main

import (
	"fmt"
	"os"

	tui "github.com/Jorgee97/i-mood/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitializeJournalList())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There is been an error %v", err)
		os.Exit(1)
	}
}
