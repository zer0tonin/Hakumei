package browser

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	files      []FileInfos
	help       help.Model // TODO: put into main
	focusIndex int
	url        string
}

func NewBrowserModel(url string) Model {
	return Model{
		url: url,
		focusIndex: 0,
		files: []FileInfos{},
		help: help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	req := browseRequest{
		target: m.url,
		path:   "/",
		search: "",
	}
	return req.do
}

func (m Model) updateFocus() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.focusIndex {
	case 0:
		return m, cmd
	}
	return m, cmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case browseResponse:
		m.files = msg.FileInfos
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, browserKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, browserKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, browserKeys.Down):
			m.focusIndex = (1 - m.focusIndex) % 2
			return m.updateFocus()
		case key.Matches(msg, browserKeys.Up):
			m.focusIndex = (m.focusIndex + 1) % 2
			return m.updateFocus()
		}
	case error:
		return m, nil
	}

	switch m.focusIndex {
	case 0:
		return m, cmd
	}

	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf("ok")
}
