package zcubed

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zcubed-mc/lib/zcubed/tui"
)

const DEFAULT_ZCUBED = `{"version": 1, "modules": []}`

func Init() {
	if len(os.Args) < 2 {
		RunTUI()
		return
	}

	if os.Args[1] != "init" {
		println("Usage: zcubed [init (optional)]")
		return
	}

	if configExists() {
		println("Config already exists")
		return
	}

	if err := os.WriteFile(path(), []byte(DEFAULT_ZCUBED), os.ModePerm); err != nil {
		panic(err)
	}

	println("Created a new project in", path())
}

func RunTUI() {
	if !configExists() {
		println("Not a Z³ project. Use `zcubed init` to create a new project")
		return
	}

	if _, err := tea.NewProgram(tui.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func configExists() bool {
	if _, err := os.Stat(path()); os.IsNotExist(err) {
		return false
	}

	return true
}

func path() string {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(workDir, "zcubed.json")
}
