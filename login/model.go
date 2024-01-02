package login

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	username   textinput.Model
	password   textinput.Model
	help       help.Model
	focusIndex int
	url        string
}

func NewLoginModel(url string) Model {
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

	return Model{
		username:   usernameInput,
		password:   passwordInput,
		focusIndex: 0,
		help:       help.New(),
		url:        url,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) updateFocus() (tea.Model, tea.Cmd) {
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case key.Matches(msg, loginKeys.Enter):
			request := loginRequest{
				target: m.url,
				Username: m.username.Value(),
				Password: m.password.Value(),
			}
			return m, request.do
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

func (m Model) View() string {
	return fmt.Sprintf(
		"Login\n\n%s\n\n%s\n\n",
		m.username.View(),
		m.password.View(),
	) + m.help.View(loginKeys)
}
