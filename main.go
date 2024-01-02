package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Page int64

const (
	Login   Page = 1
	Browser Page = 2
)

type model struct {
	page       Page
	loginModel loginModel
}

func initialModel() model {
	return model{
		page:       Login,
		loginModel: newLoginModel(),
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
