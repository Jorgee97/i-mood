package journal

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {

// 	// Is it a key press?
// 	case tea.KeyMsg:

// 		// Cool, what was the actual key pressed?
// 		switch msg.String() {

// 		// These keys should exit the program.
// 		case "ctrl+c", "q":
// 			return m, tea.Quit

// 		// The "up" and "k" keys move the cursor up
// 		case "up", "k":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}

// 		// The "down" and "j" keys move the cursor down
// 		case "down", "j":
// 			if m.cursor < len(m.choices)-1 {
// 				m.cursor++
// 			}

// 		// The "enter" key and the spacebar (a literal space) toggle
// 		// the selected state for the item that the cursor is pointing at.
// 		case "enter", " ":
// 			_, ok := m.selected[m.cursor]
// 			if ok {
// 				delete(m.selected, m.cursor)
// 			} else {
// 				m.selected[m.cursor] = struct{}{}
// 			}
// 		}

// 	}

// 	// Return the updated model to the Bubble Tea runtime for processing.
// 	// Note that we're not returning a command.
// 	return m, nil
// }

type JournalModel struct {
	textarea textarea.Model
	moods    []string
	cursor   int
	selected map[int]struct{}
}

func NewJournalModel() JournalModel {
	ta := textarea.New()
	ta.Placeholder = "Write your thoughts here..."

	return JournalModel{
		textarea: ta,
		moods:    []string{"😀 Happy", "☹️ Sad", "😡 Angry", "😱 Scared", "🤔 Confused"},
		selected: make(map[int]struct{}),
	}
}

func (jm JournalModel) Init() tea.Cmd {
	return textarea.Blink
}

func (jm JournalModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if jm.textarea.Focused() {
				jm.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return jm, tea.Quit
		case tea.KeyUp:
			if jm.cursor > 0 && !jm.textarea.Focused() {
				jm.cursor--
			}
		case tea.KeyDown:
			if jm.cursor < len(jm.moods)-1 && !jm.textarea.Focused() {
				jm.cursor++
			}
		case tea.KeyTab:
			if !jm.textarea.Focused() {
				cmd = jm.textarea.Focus()
				cmds = append(cmds, cmd)
			} else {
				jm.cursor = 0
				jm.textarea.Blur()
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case tea.KeyEnter, tea.KeySpace:
			if !jm.textarea.Focused() {
				_, ok := jm.selected[jm.cursor]
				if ok {
					delete(jm.selected, jm.cursor)
				} else {
					jm.selected[jm.cursor] = struct{}{}
				}
			}
		}
	default:
		// if len(jm.selected) > 0 {

		// 	if !jm.textarea.Focused() {
		// 		cmd = jm.textarea.Focus()
		// 		cmds = append(cmds, cmd)
		// 	}
		// }
	}

	jm.textarea, cmd = jm.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return jm, tea.Batch(cmds...)
}

func (jm JournalModel) View() string {
	s := "What is your mood today?\n\n"

	for i, mood := range jm.moods {

		cursor := " " // no cursor
		if jm.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not selected
		if _, ok := jm.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, mood)
	}

	s += fmt.Sprintf("\n\n%s\n\n", jm.textarea.View())

	// The footer
	s += "\nPress Ctrl-C to quit.\n"
	return s
}
