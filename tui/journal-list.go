package tui

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#03e759"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Mood string

const (
	Sad      Mood = "Sad"
	Happy    Mood = "Happy"
	Confused Mood = "Confused"
	Angry    Mood = "Angry"
	Scared   Mood = "Scared"
)

type MoodItem struct {
	entryDate time.Time
	mood      Mood
	thoughts  string // This will match text on the sqlite db
}

func (i MoodItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(MoodItem)
	if !ok {
		return
	}

	subString := ""
	if len(i.thoughts) > 25 {
		subString = fmt.Sprintf("%s...", strings.TrimRight(i.thoughts[0:25], " "))
	} else {
		subString = i.thoughts
	}

	year, month, day := i.entryDate.Date()
	entryDate := fmt.Sprintf("%v %d, %d", month, day, year)

	str := fmt.Sprintf("%d. %s - %s - %s", index+1, entryDate, i.mood, subString)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type JournalListModel struct {
	moodList list.Model
}

func InitializeJournalList() JournalListModel {
	items := []list.Item{
		MoodItem{time.Now(), Sad, "Some text"},
		MoodItem{time.Now(), Angry, "Some text 2"},
		MoodItem{time.Now(), Happy, "Some text 3"},
		MoodItem{time.Now(), Confused, "Some text 3"},
		MoodItem{time.Now(), Scared, "Some text that should be really long because we want to test the behavior of the thingy"},
	}

	m := JournalListModel{moodList: list.New(items, itemDelegate{}, 0, listHeight)}
	m.moodList.Title = "Some title for the list"

	return m
}

func (m JournalListModel) Init() tea.Cmd {
	return nil
}

func (m JournalListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.moodList.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			// m.quitting = true
			return m, tea.Quit

			// case "enter":
			// 	i, ok := m.moodList.SelectedItem().(item)
			// 	if ok {
			// 		m.choice = string(i)
			// 	}
			// 	return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.moodList, cmd = m.moodList.Update(msg)
	return m, cmd
}

func (m JournalListModel) View() string {
	return m.moodList.View()
}
