package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
	"github.com/mySingleLive/requi/http/response"
	"github.com/mySingleLive/requi/tui/layout"
	"strings"
)

var (
	respViewBoxStyle       = lipgloss.NewStyle()
	respTitleStyle         = lipgloss.NewStyle().MarginBottom(1).Foreground(lipgloss.Color("44"))
	respProtocolStyle      = lipgloss.NewStyle().MarginRight(1).Foreground(lipgloss.Color("66"))
	respHeaderNameStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("36"))
	respHeaderValueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("38"))
	respSuccessStatusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("120"))
	respTimeStyle          = lipgloss.NewStyle().MarginLeft(1).Foreground(lipgloss.Color("#A2A2A2"))
	respErrorStatusStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("41"))
)

type RespView struct {
}

func NewResView() *RespView {
	return &RespView{}
}

func (r *RespView) Update(msg tea.Msg) (*RespView, tea.Cmd) {
	if Context.Req.State == request.Success {
		return r, tea.Quit
	}
	return r, nil
}

func (r *RespView) headersView(resp *response.Resp) string {
	headers := resp.Headers()
	var strs strings.Builder
	for i := range headers {
		header := headers[i]
		strs.WriteString(fmt.Sprintf(
			"%s: %s\n",
			respHeaderNameStyle.Render(header.Name),
			respHeaderValueStyle.Render(header.Value),
		))
	}
	return strs.String()
}

func (r *RespView) View() string {
	req := Context.Req
	resp := req.Resp
	if req.State == request.Success && resp != nil {
		return respViewBoxStyle.Render(
			layout.VTop(
				layout.HLeft(
					respProtocolStyle.Render(resp.Protocol()),
					respSuccessStatusStyle.Render(resp.Status()),
					respTimeStyle.Render(fmt.Sprintf("%d ms", resp.Time().Milliseconds())),
					"\n",
				),
				r.headersView(resp),
			),
		)
	} else if req.State == request.Error && resp != nil {

	}
	return ""
}
