package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/zcubed-mc/lib/zcubed/tui/styles"
)

func (model Model) renderContent() string {
	height := model.Height - 3

	var content strings.Builder
	for i, line := range model.content[model.Scroll:] {
		if i >= height {
			break
		}

		if lipgloss.Width(line) > MIN_WIDTH {
			content.WriteString(line[:MIN_WIDTH])
		} else {
			content.WriteString(line)
		}

		if i < height-1 {
			content.WriteRune('\n')
		}
	}

	view := styles.ContentStyle.Width(MIN_WIDTH).
		Height(height).
		Render(content.String())

	if len(model.content)-model.Scroll > height {
		view = lipgloss.JoinVertical(0, view, model.renderFooter(len(model.content), height))
	}

	if model.Scroll > 0 {
		parts := strings.SplitN(view, "\n", 2)
		view = model.renderHeader() + parts[1]
	}

	return view
}

func (model *Model) renderFooter(length int, height int) string {
	overflowLabel := fmt.Sprintf("[ %d more lines ]", length-height-model.Scroll)
	overflowWhitespace := strings.Repeat(
		"v",
		MIN_WIDTH_HALF-lipgloss.Width(overflowLabel)+8,
	)
	return limitWidth(
		styles.OverflowStyle.Render(
			fmt.Sprintf("%s %s %s", overflowWhitespace, overflowLabel, overflowWhitespace),
		),
		MIN_WIDTH,
	)
}

func (model Model) renderHeader() string {
	overflowLabel := fmt.Sprintf("[ %d more lines ]", model.Scroll)
	overflowWhitespace := strings.Repeat(
		"^",
		MIN_WIDTH_HALF-(lipgloss.Width(overflowLabel)/2)-1,
	)

	return limitWidth(
		fmt.Sprintf("%s %s %s\n", overflowWhitespace, overflowLabel, overflowWhitespace),
		MIN_WIDTH,
	)
}

func limitWidth(content string, width int) string {
	var suffix string
	if strings.HasSuffix(content, "\n") {
		content = strings.TrimSuffix(content, "\n")
		suffix = "\n"
	}

	if lipgloss.Width(content) > width {
		return content[:width] + suffix
	}

	return content + suffix
}
