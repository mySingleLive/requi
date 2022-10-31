package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
	"github.com/mySingleLive/requi/http/response"
	"github.com/mySingleLive/requi/tui/layout"
	url2 "net/url"
)

var (
	mainBoxStyle      = lipgloss.NewStyle().Margin(0, 0, 0, 2)
	reqLineTitleStyle = lipgloss.NewStyle().MarginBottom(1).Foreground(lipgloss.Color("44"))
	reqTypeStyle      = lipgloss.NewStyle().Margin(0, 1, 0, 0).Bold(true).Foreground(lipgloss.Color("170"))
	reqLineBoxStyle   = lipgloss.NewStyle().Margin(0, 0, 1, 0)
	urlStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("33"))
	simpleReqView     = NewSimpleView()
	reqTypeListView   = NewReqTypeListView()
	pendingView       = NewPendingView()
	resView           = NewResView()
	Context           = NewViewContext()
)

type view uint8

const (
	Main view = iota
	ReqTypeList
)

type ViewContext struct {
	req           *request.Req
	res           *response.Resp
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

// NewSimpleView create a new simple view
func NewSimpleView() *SimpleReqView {
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

func (s *SimpleReqView) Init() tea.Cmd {
	Context.req.OnEnd(func(req *request.Req, resp *response.Resp) {
		//fmt.Println("xxxxx")
	})
	return pendingView.Start()
}

// VIEW

func (s *SimpleReqView) SendRequest() tea.Cmd {
	urlText := s.urlInput.Value()
	url, err := url2.Parse(urlText)
	if err == nil && url != nil {
		Context.req.Send()
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
	Context.req.ParseURL(s.urlInput.Value())
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
				layout.VTop(
					reqLineTitleStyle.Render("http request line"),
					reqLineBoxStyle.Render(
						layout.HLeft(
							reqTypeStyle.Render(Context.req.Type.Name()),
							s.urlInput.View())),
					pendingView.View(),
					resView.View(),
				)))
	case ReqTypeList:
		return reqTypeListView.View()
	}
	return ""
}
