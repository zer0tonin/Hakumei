package main

import (
    "fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type loginModel struct {
    username    textinput.Model 
    password    textinput.Model
    focusIndex  int
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
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC, tea.KeyEsc:
            return m, tea.Quit
	case tea.KeyTab, tea.KeyDown:
	    m.focusIndex = (1 - m.focusIndex) % 2
	    return m.updateFocus()
	case tea.KeyShiftTab, tea.KeyUp:
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
		"Login\n\n%s\n\n%s\n\n%s",
		m.username.View(),
		m.password.View(),
		"(esc to quit)",
	) + "\n"
}
