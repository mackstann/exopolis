package adapters

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mackstann/exopolis/game/domain"
)

type TerminalAdapter struct {
	inputEvents  chan domain.InputEvent
	drawRequests chan struct{}
	city         []string
}

func NewTerminalAdapter() *TerminalAdapter {
	adapter := &TerminalAdapter{
		inputEvents:  make(chan domain.InputEvent, 10),
		drawRequests: make(chan struct{}),
	}

	go func() {
		p := tea.NewProgram(adapter) //, tea.WithAltScreen())
		if err := p.Start(); err != nil {
			log.Fatal(err) // TODO don't control exit here
		}
	}()

	return adapter
}

func (a *TerminalAdapter) waitForDrawRequest() tea.Msg {
	<-a.drawRequests
	return struct{}{}
}

func (a *TerminalAdapter) GetInputEventsNonBlocking() []domain.InputEvent {
	// we'll usually get 0 events
	var events []domain.InputEvent
	for {
		select {
		case ev := <-a.inputEvents:
			events = append(events, ev)
		default:
			return events
		}
	}
}

// Satisfies bubbletea's interface
func (a *TerminalAdapter) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, a.waitForDrawRequest)
}

// Satisfies bubbletea's interface
func (a *TerminalAdapter) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("tui Update()")
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			a.inputEvents <- domain.QuitEvent
			return a, nil
		}
	}
	a.inputEvents <- domain.TODONoopEvent
	return a, a.waitForDrawRequest
}

// TODO: Wire up the real exit impl to trigger a tea.Quit

// TODO: Wire up output to this
func (a *TerminalAdapter) View() string {
	log.Println("tui View()")
	return strings.Join(a.city, "\n")
}

func (a *TerminalAdapter) TODORenderJustCity(city []string) {
	log.Println("tui getting new rendered city")
	a.city = city

	// try to send a draw request, but avoid being blocked
	select {
	case a.drawRequests <- struct{}{}:
	default:
	}
}
