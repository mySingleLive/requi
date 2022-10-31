package layout

import "github.com/charmbracelet/lipgloss"

func VLayout(pos lipgloss.Position, strs ...string) string {
	var array []string
	for i := range strs {
		str := strs[i]
		if str != "" {
			array = append(array, str)
		}
	}
	return lipgloss.JoinVertical(pos, array...)
}

func VTop(strs ...string) string {
	return VLayout(lipgloss.Top, strs...)
}

func VBottom(strs ...string) string {
	return VLayout(lipgloss.Bottom, strs...)
}

func VCenter(strs ...string) string {
	return VLayout(lipgloss.Center, strs...)
}

func VLeft(strs ...string) string {
	return VLayout(lipgloss.Left, strs...)
}

func VRight(strs ...string) string {
	return VLayout(lipgloss.Right, strs...)
}

func HLayout(pos lipgloss.Position, strs ...string) string {
	var array []string
	for i := range strs {
		str := strs[i]
		if str != "" {
			array = append(array, str)
		}
	}
	return lipgloss.JoinHorizontal(pos, array...)
}

func HTop(strs ...string) string {
	return HLayout(lipgloss.Top, strs...)
}

func HBottom(strs ...string) string {
	return HLayout(lipgloss.Bottom, strs...)
}

func HCenter(strs ...string) string {
	return HLayout(lipgloss.Center, strs...)
}

func HLeft(strs ...string) string {
	return HLayout(lipgloss.Left, strs...)
}

func HRight(strs ...string) string {
	return HLayout(lipgloss.Right, strs...)
}
