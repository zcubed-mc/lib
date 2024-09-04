package internal

func clamp(value, lower, upper int) int {
	return min(max(value, lower), upper)
}
