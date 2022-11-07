package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/tui/layout"
	"strings"
)

type HeaderRow struct {
	nameInput *textinput.Model
	valInput  *textinput.Model
}

type HeaderView struct {
	Title           string
	Rows            []HeaderRow
	PromptStyle     lipgloss.Style
	NameInputStyle  lipgloss.Style
	ValueInputStyle lipgloss.Style
	IndexStyle      lipgloss.Style
	BlockStyle      lipgloss.Style
}

func NewHeaderView() *HeaderView {
	return &HeaderView{
		Title:           "headers",
		Rows:            []HeaderRow{},
		PromptStyle:     lipgloss.NewStyle().MarginRight(1).Foreground(lipgloss.Color("#FFFFFF")),
		NameInputStyle:  lipgloss.NewStyle().Width(32),
		ValueInputStyle: lipgloss.NewStyle().Width(60),
		IndexStyle:      lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("202")),
		BlockStyle:      lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("57")),
	}
}

func (h *HeaderView) AddEmptyHeader() *HeaderRow {
	size := len(h.Rows)
	if size > 0 {
		lastRow := h.Rows[size-1]
		if lastRow.nameInput.Value() == "" && lastRow.valInput.Value() == "" {
			return nil
		}
	}
	nameInput := textinput.New()
	nameInput.Prompt = ">"
	nameInput.PromptStyle = h.PromptStyle
	nameInput.Placeholder = "header name"
	nameInput.TextStyle = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("#FFFFFF"))
	nameInput.Width = 30
	valInput := textinput.New()
	valInput.Prompt = ">"
	valInput.PromptStyle = h.PromptStyle
	valInput.Placeholder = "header value"
	valInput.Width = 20
	row := HeaderRow{
		nameInput: &nameInput,
		valInput:  &valInput,
	}
	Context.AddInput(&nameInput)
	Context.AddInput(&valInput)
	h.Rows = append(h.Rows, row)
	return &row
}

func (h *HeaderView) Init() tea.Cmd {
	return nil
}

func (h *HeaderView) Update(msg tea.Msg) (*HeaderView, tea.Cmd) {
	return h, nil
}

func (h *HeaderView) Focus() tea.Cmd {
	if len(h.Rows) > 0 {
		i := len(h.Rows) - 1
		return Context.Focus(h.Rows[i].nameInput)
	}
	return nil
}

func (h *HeaderView) Focused() bool {
	for i := range h.Rows {
		row := h.Rows[i]
		if row.nameInput.Focused() || row.valInput.Focused() {
			return true
		}
	}
	return false
}

func (h *HeaderView) View() string {
	if len(h.Rows) > 0 {
		str := strings.Builder{}
		title := "headers  "
		titleStyle := lipgloss.NewStyle().Bold(true).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#454545")).
			Foreground(lipgloss.Color("#828282"))
		str.WriteString(titleStyle.Render(title))
		str.WriteString("\n")
		for i := range h.Rows {
			row := h.Rows[i]
			str.WriteByte('\n')
			str.WriteString(
				layout.HLeft(
					h.IndexStyle.Render(fmt.Sprintf("%d. ", i+1)),
					h.NameInputStyle.Render(row.nameInput.View()),
					" ",
					h.ValueInputStyle.Render(row.valInput.View()),
				),
			)
			//str.WriteByte('\n')
		}
		return str.String()
	}
	return ""
}
