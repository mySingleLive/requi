package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/tui/layout"
	"strings"
)

var (
	headerPromptStyle = lipgloss.NewStyle().MarginRight(1).Foreground(lipgloss.Color("#FFFFFF"))
	headerInputStyle  = lipgloss.NewStyle().Width(35)
	headerBlock       = layout.NewBlock().Width(80)
)

type HeaderRow struct {
	nameInput *textinput.Model
	valInput  *textinput.Model
}

type HeaderView struct {
	Title        string
	Rows         []HeaderRow
	FocusedInput *textinput.Model
}

func NewHeaderView() *HeaderView {
	return &HeaderView{
		Title: "headers",
		Rows:  []HeaderRow{},
	}
}

func (h *HeaderView) AddEmptyHeader() *HeaderRow {
	nameInput := textinput.New()
	nameInput.Prompt = ">"
	nameInput.PromptStyle = headerPromptStyle
	nameInput.Placeholder = "header name"
	nameInput.TextStyle = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("#FFFFFF"))
	nameInput.Width = 30
	valInput := textinput.New()
	valInput.Prompt = ">"
	valInput.PromptStyle = headerPromptStyle
	valInput.Placeholder = "header value"
	valInput.Width = 20
	row := HeaderRow{
		nameInput: &nameInput,
		valInput:  &valInput,
	}
	h.Rows = append(h.Rows, row)
	return &row
}

func (h *HeaderView) Init() tea.Cmd {
	return nil
}

func (h *HeaderView) Update(msg tea.Msg) (*HeaderView, tea.Cmd) {
	var cmd tea.Cmd
	if h.FocusedInput != nil {
		if len(h.Rows) > 0 {
			input := h.Rows[0].nameInput
			*h.Rows[0].nameInput, cmd = input.Update(msg)
		}
	}
	return h, cmd
}

func (h *HeaderView) Focus() tea.Cmd {
	if len(h.Rows) > 0 {
		h.FocusedInput = h.Rows[0].nameInput
		return h.FocusedInput.Focus()
	}
	return nil
}

func (h *HeaderView) View() string {
	if len(h.Rows) > 0 {
		str := strings.Builder{}
		for i := range h.Rows {
			row := h.Rows[i]
			str.WriteByte('\n')
			str.WriteString(
				layout.HLeft(
					fmt.Sprintf("  %d. ", i+1),
					headerInputStyle.Render(row.nameInput.View()),
					" ",
					headerInputStyle.Render(row.valInput.View()),
				),
			)
			str.WriteByte('\n')
		}
		str.WriteByte('\n')
		return headerBlock.Render(str.String())
	}
	return ""
}
