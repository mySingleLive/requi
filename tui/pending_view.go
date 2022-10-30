package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
)

var (
	pendingBoxStyle  = lipgloss.NewStyle().PaddingLeft(1)
	pendingTextStyle = lipgloss.NewStyle().MarginLeft(1).Padding(0).Bold(true)
)

type PendingView struct {
	spinner spinner.Model
	text    string
}

func NewPendingView() *PendingView {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return &PendingView{
		spinner: s,
		text:    "sending request ...",
	}
}

func (p *PendingView) Start() tea.Cmd {
	return p.spinner.Tick
}

func (p *PendingView) Update(msg tea.Msg) (*PendingView, tea.Cmd) {
	var cmd tea.Cmd
	if Context.req.State == request.Sending {
		p.spinner, cmd = p.spinner.Update(msg)
	}
	return p, cmd
}

func (p *PendingView) View() string {
	if Context.req.State == request.Sending {
		return pendingBoxStyle.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				p.spinner.View(),
				pendingTextStyle.Render(p.text),
			))
	}
	return ""
}
