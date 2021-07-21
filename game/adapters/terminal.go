package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type cityModel []rune

type tickMsg time.Time

func main() {
	p := tea.NewProgram(cityModel([]rune("hello")), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m cityModel) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, tickEverySecond())
}

func (m cityModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		//city.step()
		return m, nil

	}

	return m, nil
}

func (m cityModel) View() string {
	return fmt.Sprintf("\n\n" + (string)(m) + "\nfoo\n")
}

func tickEverySecond() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
