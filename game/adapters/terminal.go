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
	city         [][]string
	cursorX      int
	cursorY      int
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
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			a.inputEvents <- game.QuitEvent
			return a, nil
		case "h":
			a.inputEvents <- game.CursorLeft
			return a, nil
		case "j":
			a.inputEvents <- game.CursorDown
			return a, nil
		case "k":
			a.inputEvents <- game.CursorUp
			return a, nil
		case "l":
			a.inputEvents <- game.CursorRight
			return a, nil
		case "r":
			a.inputEvents <- game.BuildResidential
			return a, nil
		case "p":
			a.inputEvents <- game.BuildPowerPlant
			return a, nil
		case "d":
			a.inputEvents <- game.BuildRoad
			return a, nil
		}
	case quitRequest:
		return a, tea.Quit
	}
	return a, a.waitForDrawRequest
}

func (a *TerminalAdapter) View() string {
	log.Println("tui View()")
	out := []string{}
	for _, row := range a.compositeLayers() {
		out = append(out, strings.Join(row, ""))
	}
	return strings.Join(out, "\n")
}

func (a *TerminalAdapter) UpdateCity(city [][]string) {
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

func (a *TerminalAdapter) MoveCursor(x, y int) {
	if y > len(a.city)-1 {
		y = len(a.city) - 1
	}
	if x > len(a.city[0])-1 {
		x = len(a.city[0]) - 1
	}
	a.cursorX = x
	a.cursorY = y
}

func (a *TerminalAdapter) compositeLayers() [][]string {
	if len(a.city) == 0 {
		return a.city
	}
	output := make([][]string, len(a.city))
	copy(output, a.city)

	row := output[a.cursorY]
	row[a.cursorX] = "◸"
	output[a.cursorY] = row

	return output
}
