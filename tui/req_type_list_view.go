package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http/request"
	"github.com/spf13/cast"
	"io"
)

const listHeight = 16

var (
	titleStyle            = lipgloss.NewStyle().MarginLeft(1).Bold(true).Foreground(lipgloss.Color("170"))
	itemStyle             = lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#929292"))
	selectedItemStyle     = lipgloss.NewStyle().Width(27).PaddingLeft(1).Bold(true).Background(lipgloss.Color("170"))
	lastSelectedItemStyle = lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("170"))
	paginationStyle       = list.DefaultStyles().PaginationStyle.PaddingLeft(3)
	helpStyle             = list.DefaultStyles().HelpStyle.PaddingLeft(3).PaddingBottom(1)
	quitTextStyle         = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	reqTypes              = []request.RequestType{
		request.GET,
		request.POST,
		request.PUT,
		request.HEAD,
		request.DELETE,
		request.OPTIONS,
		request.TRACE,
		request.PATCH,
	}
)

type item request.RequestType

func (i item) FilterValue() string {
	reqType := request.RequestType(i)
	return reqType.Name()
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}
	reqType := request.RequestType(i)
	str := fmt.Sprintf("%d. %s", index+1, reqType.Name())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	} else if index == reqTypeListView.lastSelectedIndex {
		fn = func(s string) string {
			return lastSelectedItemStyle.Render(s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type ReqTypeListView struct {
	list              list.Model
	lastSelectedIndex int
}

func NewReqTypeListView() *ReqTypeListView {

	const defaultWidth = 40

	var items []list.Item
	for i := range reqTypes {
		it := reqTypes[i]
		items = append(items, item(it))
	}

	list := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	list.Title = "Choose the request type"
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.KeyMap.CursorUp.SetKeys("up", "k", "ctrl+p")
	list.KeyMap.CursorDown.SetKeys("down", "j", "ctrl+n")
	list.SetShowHelp(false)
	list.Styles.Title = titleStyle
	list.Styles.HelpStyle = helpStyle
	list.Styles.PaginationStyle = paginationStyle

	return &ReqTypeListView{
		list:              list,
		lastSelectedIndex: 0,
	}
}

func (tl *ReqTypeListView) Init() tea.Cmd {
	return nil
}

func (tl *ReqTypeListView) SelectCurrentType() {
	for i, reqType := range reqTypes {
		if reqType.Name() == Context.req.Type.Name() {
			tl.list.Select(i)
			tl.lastSelectedIndex = i
			return
		}
	}
}

func (tl *ReqTypeListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch Context.view {
	case Main:
		return Context.SimpleReqView.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		tl.list.SetWidth(msg.Width)
		return tl, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return tl, tea.Quit
		case tea.KeyEsc, tea.KeyCtrlT:
			Context.view = Main
			simpleReqView.urlInput.Focus()
			return tl, nil
		case tea.KeyEnter, tea.KeySpace:
			i, ok := tl.list.SelectedItem().(item)
			if ok {
				Context.req.Type = request.RequestType(i)
			}
			Context.view = Main
			simpleReqView.urlInput.Focus()
			return tl, nil
		}

		keyStr := msg.String()
		i := cast.ToInt(keyStr)
		if i >= 1 && i <= len(reqTypes) {
			tl.list.Select(i - 1)
		}
	}

	var cmd tea.Cmd
	tl.list, cmd = tl.list.Update(msg)
	return tl, cmd
}

func (tl *ReqTypeListView) View() string {
	switch Context.view {
	case Main:
		return Context.SimpleReqView.View()
	}
	//tl.SelectCurrentType()
	return "\n" + tl.list.View()
}
