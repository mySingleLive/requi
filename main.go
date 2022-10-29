package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mySingleLive/requi/tui"
)

func main() {
	p := tea.NewProgram(
		tui.Context.SimpleReqView,
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
