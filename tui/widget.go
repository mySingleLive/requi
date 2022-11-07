package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type WidgetModel interface {
	textinput.Model | textarea.Model
}

type OnFocus[T WidgetModel] func(widget *Widget[T])

type OnBlur[T WidgetModel] func(widget *Widget[T])

type Widget[T WidgetModel] struct {
	Model   *T
	onFocus OnFocus[T]
	onBlur  OnBlur[T]
}

func NewWidget[T WidgetModel](model *T) *Widget[T] {
	return &Widget[T]{
		Model: model,
	}
}

func (w *Widget[T]) OnFocus(focus OnFocus[T]) *Widget[T] {
	w.onFocus = focus
	return w
}

func (w *Widget[T]) OnBlur(blur OnBlur[T]) *Widget[T] {
	w.onBlur = blur
	return w
}

func (w *Widget[T]) Focus() tea.Cmd {
	var ti interface{} = w.Model
	if w.onFocus != nil {
		defer w.onFocus(w)
	}
	switch v := ti.(type) {
	case *textinput.Model:
		return v.Focus()
	case *textarea.Model:
		return v.Focus()
	}
	return nil
}
