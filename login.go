package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type loginKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k loginKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k loginKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter},
		{k.Help, k.Quit},
	}
}

var loginKeys = loginKeyMap{
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

type loginModel struct {
	username   textinput.Model
	password   textinput.Model
	help       help.Model
	focusIndex int
}

func newLoginModel() loginModel {
	usernameInput := textinput.New()
	usernameInput.Placeholder = "Username"
	usernameInput.Focus()
	usernameInput.CharLimit = 156
	usernameInput.Width = 20

	passwordInput := textinput.New()
	passwordInput.Placeholder = "Password"
	passwordInput.Blur()
	passwordInput.CharLimit = 156
	passwordInput.Width = 20

	return loginModel{
		username:   usernameInput,
		password:   passwordInput,
		focusIndex: 0,
		help:       help.New(),
	}
}

func (m loginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m loginModel) updateFocus() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.focusIndex {
	case 0:
		cmd = m.username.Focus()
		m.password.Blur()
	case 1:
		cmd = m.password.Focus()
		m.username.Blur()
	}
	return m, cmd
}

func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, loginKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, loginKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, loginKeys.Down):
			m.focusIndex = (1 - m.focusIndex) % 2
			return m.updateFocus()
		case key.Matches(msg, loginKeys.Up):
			m.focusIndex = (m.focusIndex + 1) % 2
			return m.updateFocus()
		}
	case error:
		return m, nil
	}

	switch m.focusIndex {
	case 0:
		m.username, cmd = m.username.Update(msg)
		return m, cmd
	case 1:
		m.password, cmd = m.password.Update(msg)
		return m, cmd
	}

	return m, cmd
}

func (m loginModel) View() string {
	return fmt.Sprintf(
		"Login\n\n%s\n\n%s\n\n",
		m.username.View(),
		m.password.View(),
	) + m.help.View(loginKeys)
}
