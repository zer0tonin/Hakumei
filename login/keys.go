package login

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter},
		{k.Help, k.Quit},
	}
}

var loginKeys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys(tea.KeyUp.String(), tea.KeyShiftTab.String()),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys(tea.KeyDown.String(), tea.KeyTab.String()),
		key.WithHelp("↓", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("enter", "submit"),
	),
	Help: key.NewBinding(
		key.WithKeys(tea.KeyCtrlH.String()),
		key.WithHelp("ctrl+h", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String(), tea.KeyCtrlC.String()),
		key.WithHelp("ctrl+c/esc", "quit"),
	),
}
