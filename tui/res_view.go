package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
	"github.com/mySingleLive/requi/tui/layout"
)

var resViewBoxStyle = lipgloss.NewStyle()
var resTitleStyle = lipgloss.NewStyle().MarginBottom(1).Foreground(lipgloss.Color("44"))
var resStatusTileStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("44"))

type ResView struct {
}

func NewResView() *ResView {
	return &ResView{}
}

func (r *ResView) Update(msg tea.Msg) (*ResView, tea.Cmd) {
	if Context.req.State == request.Success {
		return r, tea.Quit
	}
	return r, nil
}

func (r *ResView) View() string {
	req := Context.req
	resp := req.Resp
	if req.State == request.Success && resp != nil {
		return resViewBoxStyle.Render(
			layout.VTop(
				resTitleStyle.Render("response"),
				layout.HLeft(
					resStatusTileStyle.Render("status: "),
					resp.Status(),
					"\n",
				),
				resp.Text(),
			),
		)
	}
	return ""
}
