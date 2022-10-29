package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http"
)

var (
	reqTypeStyle    = lipgloss.NewStyle().PaddingLeft(2).Bold(true).Foreground(lipgloss.Color("170"))
	urlStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#49adff")).UnderlineSpaces(true).Blink(true).Faint(true)
	simpleReqView   = NewSimpleView()
	reqTypeListView = NewReqTypeListView()
	Context         = NewViewContext()
)

type view uint8

const (
	Main view = iota
	ReqTypeList
)

type ViewContext struct {
	view            view
	SimpleReqView   *SimpleReqView
	ReqTypeListView *ReqTypeListView
	reqType         http.RequestType
}

func NewViewContext() *ViewContext {
	return &ViewContext{
		view:            Main,
		SimpleReqView:   simpleReqView,
		ReqTypeListView: reqTypeListView,
		reqType:         http.GET,
	}
}

type SimpleReqView struct {
	url textinput.Model
	err error
}

func NewSimpleView() *SimpleReqView {

	// request type list view

	// url text input
	url := textinput.New()
	url.Prompt = ""
	url.Focus()
	url.Width = 120
	url.TextStyle = urlStyle

	return &SimpleReqView{
		url: url,
	}
}

func (s *SimpleReqView) Init() tea.Cmd {
	return textinput.Blink
}

// VIEW

func (s *SimpleReqView) View() string {
	switch Context.view {
	case Main:
		return fmt.Sprintf(
			"\n%s %s\n",
			reqTypeStyle.Render(Context.reqType.Name()),
			s.url.View(),
		) + "\n"
	case ReqTypeList:
		return Context.ReqTypeListView.View()
	}
	return ""
}

// UPDATE

func (s *SimpleReqView) UpdateMainView(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			s.url.Blur()
			return s, tea.Quit
		case tea.KeyCtrlT:
			Context.view = ReqTypeList
			reqTypeListView.SelectCurrentType()
			return s, nil
		}
	}

	s.url, cmd = s.url.Update(msg)
	return s, cmd
}

func (s *SimpleReqView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch Context.view {
	case ReqTypeList:
		reqTypeListView.SelectCurrentType()
		return reqTypeListView.Update(msg)
	}
	return s.UpdateMainView(msg)
}
