package tui

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zcubed-mc/lib/zcubed/internal"
	"github.com/zcubed-mc/lib/zcubed/tui/screen"
)

const MIN_WIDTH = 64
const MIN_WIDTH_HALF = 32
const MIN_HEIGHT = 10

const HOME_SCREEN = 0

var Screens = []screen.Screen{
	screen.NewHomeScreen(),
}

type Model struct {
	Width  int
	Height int
	Screen int
	Scroll int

	content []string
}

func NewModel() Model {
	return Model{
		Width:  0,
		Height: 0,
		Screen: HOME_SCREEN,
	}
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		model.Width, model.Height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return model, tea.Quit

		case "k", "up":
			model.Scroll -= 1

		case "K", "shift+up":
			model.Scroll = 0

		case "j", "down":
			model.Scroll += 1

		case "J", "shift+down":
			model.Scroll = math.MaxInt
		}
	}

	model.content = Screens[model.Screen].View(MIN_WIDTH)
	model = model.limitScroll()

	return model, nil
}

func (model Model) View() string {
	if model.Width < MIN_WIDTH || model.Height < MIN_HEIGHT {
		return "Screen is too small!"
	}

	window := model.renderWindow()

	content := lipgloss.Place(
		MIN_WIDTH,
		model.Height-2,
		lipgloss.Center,
		lipgloss.Center,
		model.renderContent(),
	)

	return internal.PlaceOverlay(model.Width/2-MIN_WIDTH_HALF, 0, content, window)
}

func (model Model) limitScroll() Model {
	if model.Scroll < 0 {
		model.Scroll = 0
	}

	limit := len(model.content) - model.Height + 3
	if model.Scroll > limit {
		model.Scroll = limit
	}

	return model
}
