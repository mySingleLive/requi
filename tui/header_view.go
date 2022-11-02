package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mySingleLive/requi/tui/layout"
	"strings"
)

type HeaderRow struct {
	nameInput textinput.Model
	valInput  textinput.Model
}

type HeaderView struct {
	Title string
	Rows  []HeaderRow
}

func NewHeaderView() *HeaderView {
	return &HeaderView{
		Title: "headers",
		Rows:  []HeaderRow{},
	}
}

func (h *HeaderView) AddEmptyHeader() *HeaderRow {
	nameInput := textinput.New()
	nameInput.Prompt = "Header Name"
	nameInput.Width = 30
	valInput := textinput.New()
	valInput.Prompt = "Header Value"
	valInput.Width = 30
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

	return h, nil
}

func (h *HeaderView) Focus() tea.Cmd {
	if len(h.Rows) > 0 {
		return h.Rows[0].nameInput.Focus()
	}
	return nil
}

func (h *HeaderView) View() string {
	if len(h.Rows) > 0 {
		strs := strings.Builder{}
		for i := range h.Rows {
			row := h.Rows[i]
			strs.WriteString(
				layout.HLeft(
					row.nameInput.View(),
					" ",
					row.valInput.Value(),
				),
			)
		}
		return strs.String()
	}
	return ""
}
