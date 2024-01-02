package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Page int64

const (
    Login Page = 1
    Browser Page = 2
)

type model struct {
    page Page
    loginModel loginModel
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
	page: Login,
	loginModel: loginModel{
	    username: usernameInput,
	    password: passwordInput,
	    focusIndex: 0,
	},
    }
}

func (m model) Init() tea.Cmd {
    switch m.page {
    case Login:
	return m.loginModel.Init()
    case Browser:
	fallthrough
    default:
	return nil
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch m.page {
    case Login:
	return m.loginModel.Update(msg)
    case Browser:
	fallthrough
    default:
	return nil, nil
    }
}

func (m model) View() string {
    switch m.page {
    case Login:
	return m.loginModel.View()
    case Browser:
	fallthrough
    default:
	return "" 
    }
}


func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
