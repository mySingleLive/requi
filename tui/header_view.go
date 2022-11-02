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
	headerPromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#828282"))
)

type HeaderRow struct {
	nameInput textinput.Model
	valInput  textinput.Model
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
	nameInput.Prompt = "header name"
	nameInput.PromptStyle = headerPromptStyle
	nameInput.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	nameInput.Width = 20
	valInput := textinput.New()
	valInput.Prompt = "header value"
	valInput.PromptStyle = headerPromptStyle
	valInput.Width = 20
	row := HeaderRow{
		nameInput: nameInput,
		valInput:  valInput,
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
			h.Rows[0].nameInput, cmd = input.Update(msg)
		}
	}
	return h, cmd
}

func (h *HeaderView) Focus() tea.Cmd {
	if len(h.Rows) > 0 {
		input := h.Rows[0].nameInput
		h.FocusedInput = &input
		h.FocusedInput.Focus()
	}
	return nil
}

func (h *HeaderView) View() string {
	if len(h.Rows) > 0 {
		str := strings.Builder{}
		for i := range h.Rows {
			row := h.Rows[i]
			str.WriteString(
				layout.HLeft(
					fmt.Sprintf("%d. ", i+1),
					row.nameInput.View(),
					" ",
					row.valInput.View(),
				),
			)
			str.WriteString("\n")
		}
		return str.String()
	}
	return ""
}
