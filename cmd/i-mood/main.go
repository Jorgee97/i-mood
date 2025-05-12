package main

import (
	"fmt"
	"os"

	journal "github.com/Jorgee97/i-mood/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(journal.NewJournalModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There is been an error %v", err)
		os.Exit(1)
	}
}
