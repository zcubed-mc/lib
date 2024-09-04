package screen

import "strings"

type HomeScreen struct {
}

func NewHomeScreen() HomeScreen {
	return HomeScreen{}
}

func (screen HomeScreen) Title() string {
	return "Home"
}

func (screen HomeScreen) View(width int) []string {
	var result []string

	for range 100 {
		result = append(result, strings.Repeat("a", 100))
		result = append(result, strings.Repeat("b", 100))
	}
	result = append(result, strings.Repeat("c", 10))

	return result
}
