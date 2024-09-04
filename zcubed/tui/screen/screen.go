package screen

type Screen interface {
	Title() string
	View(width int) []string
}
