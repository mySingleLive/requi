package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mySingleLive/requi/http"
	"github.com/spf13/cast"
	"io"
)

const listHeight = 16

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	reqTypes          = []http.RequestType{
		http.GET,
		http.POST,
		http.PUT,
		http.HEAD,
		http.DELETE,
		http.OPTIONS,
		http.TRACE,
		http.PATCH,
	}
)

type item http.RequestType

func (i item) FilterValue() string {
	reqType := http.RequestType(i)
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
	reqType := http.RequestType(i)
	str := fmt.Sprintf("%d. %s", index+1, reqType.Name())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type ReqTypeListView struct {
	list list.Model
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
	list.Styles.Title = titleStyle
	list.Styles.HelpStyle = helpStyle
	list.Styles.PaginationStyle = paginationStyle

	return &ReqTypeListView{
		list: list,
	}
}

func (tl *ReqTypeListView) Init() tea.Cmd {
	return nil
}

func (tl *ReqTypeListView) SelectCurrentType() {
	for i, reqType := range reqTypes {
		if reqType.Name() == Context.reqType.Name() {
			tl.list.Select(i)
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
			return tl, nil
		case tea.KeyEnter:
			i, ok := tl.list.SelectedItem().(item)
			if ok {
				Context.reqType = http.RequestType(i)
			}
			Context.view = Main
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
