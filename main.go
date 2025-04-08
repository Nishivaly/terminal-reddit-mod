package main

import (
	"fmt"
	"os"

	"github.com/Nishivaly/terminal-reddit-mod/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(tui.Model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
