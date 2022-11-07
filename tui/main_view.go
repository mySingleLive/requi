package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
	"github.com/mySingleLive/requi/tui/layout"
	url2 "net/url"
)

var (
	mainBoxStyle      = lipgloss.NewStyle().Margin(0, 0, 0, 0)
	reqLineTitleStyle = lipgloss.NewStyle().MarginBottom(1).Foreground(lipgloss.Color("44"))
	reqTypeStyle      = lipgloss.NewStyle().Margin(0, 1, 0, 0).Bold(true).Foreground(lipgloss.Color("170"))
	reqLineBoxStyle   = lipgloss.NewStyle().Margin(0, 0, 1, 0)
	urlStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("32"))
	simpleReqView     = NewSimpleView()
	reqTypeListView   = NewReqTypeListView()
	headerView        = NewHeaderView()
	pendingView       = NewPendingView()
	respView          = NewResView()
	Context           = NewViewContext()
)

type view uint8

const (
	Main view = iota
	ReqTypeList
)

type ViewContext struct {
	Req           *request.Req
	view          view
	SimpleReqView *SimpleReqView
	showHeaders   bool
	focusedIndex  int
	Inputs        []*textinput.Model
}

func NewViewContext() *ViewContext {
	return &ViewContext{
		Req:           request.New(request.GET),
		view:          Main,
		SimpleReqView: simpleReqView,
		showHeaders:   false,
		focusedIndex:  -1,
		Inputs:        []*textinput.Model{},
	}
}

func (c *ViewContext) IsShowHeaders() bool {
	return len(c.Req.Headers) > 0 || c.showHeaders
}

func (c *ViewContext) AddInput(input *textinput.Model) int {
	c.Inputs = append(c.Inputs, input)
	return len(c.Inputs) - 1
}

func (c *ViewContext) GetIndex(input *textinput.Model) int {
	for i := range c.Inputs {
		if c.Inputs[i] == input {
			return i
		}
	}
	return -1
}

func (c *ViewContext) FocusIndex(index int) tea.Cmd {
	if c.focusedIndex >= 0 && len(c.Inputs) > 0 {
		c.Inputs[c.focusedIndex].Blur()
	}
	c.focusedIndex = index
	return c.Inputs[c.focusedIndex].Focus()
}

func (c *ViewContext) Focus(input *textinput.Model) tea.Cmd {
	i := c.GetIndex(input)
	if i != -1 {
		return c.FocusIndex(i)
	}
	return nil
}

func (c *ViewContext) FocusedInput() *textinput.Model {
	if c.focusedIndex == -1 {
		return nil
	}
	return c.Inputs[c.focusedIndex]
}

func (c *ViewContext) Blur() {
	input := c.FocusedInput()
	if input != nil {
		input.Blur()
		c.focusedIndex = -1
	}
}

func (c *ViewContext) FocusNext() tea.Cmd {
	index := c.focusedIndex + 1
	if index >= len(c.Inputs) {
		index = 0
	}
	return c.FocusIndex(index)
}

func (c *ViewContext) FocusPrev() tea.Cmd {
	index := c.focusedIndex - 1
	if index < 0 {
		index = len(c.Inputs) - 1
	}
	return c.FocusIndex(index)
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

	return &SimpleReqView{
		urlInput: url,
		sendBtn:  "Send",
	}
}

func (s *SimpleReqView) Init() tea.Cmd {
	index := Context.AddInput(&s.urlInput)
	Context.FocusIndex(index)
	return pendingView.Start()
}

// VIEW

func (s *SimpleReqView) SendRequest() tea.Cmd {
	urlText := s.urlInput.Value()
	url, err := url2.Parse(urlText)
	if err == nil && url != nil {
		Context.Req.Send()
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
			Context.Blur()
			return s, tea.Quit
		case tea.KeyCtrlX:
			Context.Blur()
			return s, s.SendRequest()
		case tea.KeyCtrlT:
			Context.view = ReqTypeList
			reqTypeListView.SelectCurrentType()
			return s, nil
		case tea.KeyCtrlH:
			headerView.AddEmptyHeader()
			return s, headerView.Focus()
		case tea.KeyTab:
			return s, Context.FocusNext()
		case tea.KeyEnter:

		case tea.KeyShiftTab:
			return s, Context.FocusPrev()
		}
	}
	var headerCmd, pendingCmd, respCmd tea.Cmd
	var inputCmds []tea.Cmd
	focusedInput := Context.FocusedInput()
	if focusedInput != nil {
		var ic tea.Cmd
		*focusedInput, ic = focusedInput.Update(msg)
		inputCmds = append(inputCmds, ic)
	}
	if s.urlInput.Focused() {
		Context.Req.ParseURL(s.urlInput.Value())
	}
	headerView, headerCmd = headerView.Update(msg)
	pendingView, pendingCmd = pendingView.Update(msg)
	respView, respCmd = respView.Update(msg)
	inputCmds = append(inputCmds, headerCmd, pendingCmd, respCmd)
	return s, tea.Batch(inputCmds...)
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
					reqLineBoxStyle.Render(
						layout.HLeft(
							reqTypeStyle.Render(Context.Req.Type.Name()),
							s.urlInput.View())),
					headerView.View(),
					pendingView.View(),
					respView.View(),
				)))
	case ReqTypeList:
		return reqTypeListView.View()
	}
	return ""
}
