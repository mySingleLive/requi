package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
	url2 "net/url"
)

var (
	mainBoxStyle        = lipgloss.NewStyle().Margin(0, 0, 0, 2)
	requestLineBoxStyle = lipgloss.NewStyle().Margin(0, 0, 1, 0)
	reqTypeStyle        = lipgloss.NewStyle().Margin(0, 1, 0, 0).Padding(0, 1, 0, 1).Bold(true).Foreground(lipgloss.Color("170"))
	urlStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))
	simpleReqView       = NewSimpleView()
	reqTypeListView     = NewReqTypeListView()
	pendingView         = NewPendingView()
	Context             = NewViewContext()
)

type view uint8

const (
	Main view = iota
	ReqTypeList
)

type ViewContext struct {
	req           *request.Request
	view          view
	SimpleReqView *SimpleReqView
}

func NewViewContext() *ViewContext {
	return &ViewContext{
		req:           request.New(request.GET),
		view:          Main,
		SimpleReqView: simpleReqView,
	}
}

type SimpleReqView struct {
	urlInput textinput.Model
	err      error
	sendBtn  string
}

func NewSimpleView() *SimpleReqView {

	// request type list view

	// urlInput text input
	url := textinput.New()
	url.Prompt = ""
	url.TextStyle = urlStyle
	url.Width = 80
	url.CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("120"))
	url.Focus()

	return &SimpleReqView{
		urlInput: url,
		sendBtn:  "Send",
	}
}

func (s SimpleReqView) BuildRequest() *request.Request {
	req := request.New(Context.req.Type)
	return req
}

func (s *SimpleReqView) Init() tea.Cmd {
	return pendingView.spinner.Tick
}

// VIEW

func (s SimpleReqView) SendRequest() tea.Cmd {
	urlText := s.urlInput.Value()
	url, err := url2.Parse(urlText)
	if err == nil && url != nil {
		Context.req.State = request.Sending
		return pendingView.Start()
	}
	return nil
}

// UPDATE

func (s *SimpleReqView) UpdateMainView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			s.urlInput.Blur()
			return s, tea.Quit
		case tea.KeyEnter, tea.KeyEsc:
			s.urlInput.Blur()
			return s, s.SendRequest()
		case tea.KeyCtrlT:
			Context.view = ReqTypeList
			reqTypeListView.SelectCurrentType()
			return s, nil
		}
	}
	var urlCmd, pendingCmd tea.Cmd
	s.urlInput, urlCmd = s.urlInput.Update(msg)
	pendingView, pendingCmd = pendingView.Update(msg)
	return s, tea.Batch(urlCmd, pendingCmd)
}

func (s *SimpleReqView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch Context.view {
	case ReqTypeList:
		reqTypeListView.SelectCurrentType()
		return reqTypeListView.Update(msg)
	}
	return s.UpdateMainView(msg)
}

func (s *SimpleReqView) View() string {
	switch Context.view {
	case Main:
		return fmt.Sprintf(
			"\n%s\n",
			mainBoxStyle.Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					requestLineBoxStyle.Render(
						lipgloss.JoinHorizontal(
							lipgloss.Center,
							reqTypeStyle.Render(Context.req.Type.Name()),
							s.urlInput.View())),
					pendingView.View())),
		) + "\n"
	case ReqTypeList:
		return reqTypeListView.View()
	}
	return ""
}
