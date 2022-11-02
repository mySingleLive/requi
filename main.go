package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mySingleLive/requi/http/request"
	"github.com/mySingleLive/requi/tui"
)

func main() {
	p := tea.NewProgram(
		tui.Context.SimpleReqView,
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
	if tui.Context.Req.Resp != nil {
		if tui.Context.Req.State == request.Success {
			fmt.Println(tui.Context.Req.Resp.Result())
		} else if tui.Context.Req.State == request.Error {
			fmt.Println(tui.Context.Req.Resp.Error.Error())
		}
	}
}
