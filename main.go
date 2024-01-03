package main

import (
	"log"
	"os"

	"github.com/zer0tonin/Hakumei/login"
	"github.com/zer0tonin/Hakumei/browser"

	tea "github.com/charmbracelet/bubbletea"
)

type Page int64

const (
	Login   Page = 1
	Browser Page = 2
)

type model struct {
	page       Page
	loginModel login.Model
	browserModel browser.Model
}

func initialModel(url string) model {
	return model{
		page:       Login,
		loginModel: login.NewLoginModel(url),
		browserModel: browser.NewBrowserModel(url),
	}
}

func (m model) Init() tea.Cmd {
	switch m.page {
	case Login:
		return m.loginModel.Init()
	case Browser:
		return m.browserModel.Init()
	default:
		return nil
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case login.LoginResponse:
		m.page = 2
	}

	switch m.page {
	case Login:
		return m.loginModel.Update(msg)
	case Browser:
		return m.browserModel.Update(msg)
	default:
		return nil, nil
	}
}

func (m model) View() string {
	switch m.page {
	case Login:
		return m.loginModel.View()
	case Browser:
		return m.browserModel.View()
	default:
		return ""
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("URL required")
	}

	p := tea.NewProgram(initialModel(os.Args[1]))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
