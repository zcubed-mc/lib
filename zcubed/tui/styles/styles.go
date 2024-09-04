package styles

import "github.com/charmbracelet/lipgloss"

var ContentStyle = lipgloss.NewStyle().
	Background(BackgroundColor).
	Foreground(ForegroundColor).
	MarginTop(1)

var WindowStyle = ContentStyle.MarginLeft(1)

var ShadowStyle = lipgloss.NewStyle().Background(ShadowColor).Foreground(ShadowColor)

var OverflowStyle = lipgloss.NewStyle().
	Background(ShadowColor).
	Foreground(TerminalColor).
	AlignHorizontal(lipgloss.Center)
