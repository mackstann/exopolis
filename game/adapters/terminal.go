package adapters

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mackstann/exopolis/game/domain"
)

type TerminalAdapter struct {
	events chan domain.InputEvent
}

func NewTerminalAdapter() *TerminalAdapter {
	adapter := &TerminalAdapter{
		events: make(chan domain.InputEvent),
	}

	p := tea.NewProgram(adapter, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err) // TODO don't control exit here
	}

	return adapter
}

func (a *TerminalAdapter) Events() chan domain.InputEvent {
	return a.events
}

// Satisfies bubbletea's interface
func (a *TerminalAdapter) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Satisfies bubbletea's interface
func (a *TerminalAdapter) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			a.events <- domain.QuitEvent
			return a, nil
		}
	}
	return a, nil
}

// TODO: Wire up the real exit impl to trigger a tea.Quit

// TODO: Wire up output to this
func (a *TerminalAdapter) View() string {
	return fmt.Sprintf("\n\nfoo\n")
}
