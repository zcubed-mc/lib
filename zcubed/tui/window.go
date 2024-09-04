package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/zcubed-mc/lib/zcubed/tui/styles"
)

var shadowSide = styles.ShadowStyle.Render("  ")
var shadowOffset = strings.Repeat(" ", 3)

func (model Model) renderWindow() string {
	window := styles.WindowStyle.Width(model.Width - 3).Height(model.Height - 3).Render("")

	lines := strings.Split(window, "\n")

	for i, line := range lines {
		if i < 2 {
			continue
		}

		lines[i] = line + shadowSide
	}

	return lipgloss.JoinVertical(
		0,
		strings.Join(lines, "\n"),
		shadowOffset+styles.ShadowStyle.Width(model.Width-3).Render(" "),
	)
}
