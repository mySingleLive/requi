package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
)

var (
	requestLineBoxStyle = lipgloss.NewStyle().Margin(0, 0, 0, 1)
	reqTypeStyle        = lipgloss.NewStyle().Margin(0, 1, 0, 1).Padding(0, 1, 0, 1).Bold(true).Foreground(lipgloss.Color("170"))
	urlStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))
	simpleReqView       = NewSimpleView()
	reqTypeListView     = NewReqTypeListView()
	Context             = NewViewContext()
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
	reqType         request.RequestType
}

func NewViewContext() *ViewContext {
	return &ViewContext{
		view:            Main,
		SimpleReqView:   simpleReqView,
		ReqTypeListView: reqTypeListView,
		reqType:         request.GET,
	}
}

type SimpleReqView struct {
	url     textinput.Model
	err     error
	sendBtn string
	width   int
}

func NewSimpleView() *SimpleReqView {

	// request type list view

	// url text input
	url := textinput.New()
	url.Prompt = ""
	url.Focus()
	url.TextStyle = urlStyle
	url.Width = 80

	return &SimpleReqView{
		url:     url,
		sendBtn: "Send",
		width:   80,
	}
}

func (s SimpleReqView) BuildRequest() *request.Request {
	req := request.New(Context.reqType)
	return req
}

func (s *SimpleReqView) Init() tea.Cmd {
	return nil
}

// VIEW

func (s *SimpleReqView) View() string {
	switch Context.view {
	case Main:
		return fmt.Sprintf(
			"\n%s\n",
			requestLineBoxStyle.Render(
				lipgloss.JoinHorizontal(
					lipgloss.Center,
					reqTypeStyle.Render(Context.reqType.Name()),
					s.url.View(),
				)),
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
	case tea.WindowSizeMsg:
		s.width = msg.Width
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
	if s.url.Width > 80 {
		s.url.Width = 80
	}
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
