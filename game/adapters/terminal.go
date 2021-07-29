package adapters

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mackstann/exopolis/game"
)

type TerminalAdapter struct {
	inputEvents  chan game.InputEvent
	drawRequests chan struct{}
	quitRequest  chan struct{}
	quitComplete chan struct{}
	city         []string
}

func NewTerminalAdapter() *TerminalAdapter {
	adapter := &TerminalAdapter{
		inputEvents:  make(chan game.InputEvent, 10),
		drawRequests: make(chan struct{}),
		quitRequest:  make(chan struct{}),
		quitComplete: make(chan struct{}),
	}

	go func() {
		p := tea.NewProgram(adapter) //, tea.WithAltScreen())
		if err := p.Start(); err != nil {
			log.Fatal(err) // TODO don't control exit here
		}
		adapter.quitComplete <- struct{}{}
	}()

	return adapter
}

func (a *TerminalAdapter) waitForDrawRequest() tea.Msg {
	<-a.drawRequests
	return struct{}{}
}

type quitRequest struct{}

func (a *TerminalAdapter) waitForQuitRequest() tea.Msg {
	<-a.quitRequest
	return quitRequest{}
}

func (a *TerminalAdapter) Shutdown() {
	a.quitRequest <- quitRequest{}
	<-a.quitComplete
}

func (a *TerminalAdapter) GetInputEventsNonBlocking() []game.InputEvent {
	// we'll usually get 0 events
	var events []game.InputEvent
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
	return tea.Batch(tea.EnterAltScreen, a.waitForDrawRequest, a.waitForQuitRequest)
}

// Satisfies bubbletea's interface
func (a *TerminalAdapter) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	// draw requests hit this function but there's no need to take any action; bubbletea will redraw simply because
	// Update was called.
	log.Println("tui Update()")
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			a.inputEvents <- game.QuitEvent
			return a, nil
		}
	case quitRequest:
		return a, tea.Quit
	}
	return a, a.waitForDrawRequest
}

func (a *TerminalAdapter) View() string {
	log.Println("tui View()")
	return strings.Join(a.city, "\n")
}

func (a *TerminalAdapter) UpdateCity(city []string) {
	log.Println("tui getting new rendered city")
	a.city = city
}

func (a *TerminalAdapter) Redraw() {
	// try to send a draw request, but avoid being blocked (if one is already pending)
	select {
	case a.drawRequests <- struct{}{}:
	default:
	}
}
