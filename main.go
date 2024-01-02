package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    username    textinput.Model 
    password    textinput.Model
    focusIndex  int
}

func initialModel() model {
    usernameInput:= textinput.New()
    usernameInput.Placeholder = "Username"
    usernameInput.Focus()
    usernameInput.CharLimit = 156
    usernameInput.Width = 20

    passwordInput:= textinput.New()
    passwordInput.Placeholder = "Password"
    passwordInput.Blur()
    passwordInput.CharLimit = 156
    passwordInput.Width = 20

    return model{
	    username: usernameInput,
	    password: passwordInput,
	    focusIndex: 0,
    }
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) updateFocus() (tea.Model, tea.Cmd) {
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	return fmt.Sprintf(
		"Login\n\n%s\n\n%s\n\n%s",
		m.username.View(),
		m.password.View(),
		"(esc to quit)",
	) + "\n"
}


func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
