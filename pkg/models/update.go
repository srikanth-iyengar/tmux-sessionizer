package models

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.IsSearch {
			switch msg.String() {
			case "esc", "q":
				return m, tea.Quit
			case "enter":
				CreateSession(&m, nil)
				return m, tea.Quit
			case "tab":
				m.Toggle = !m.Toggle
				return m, nil
			case "/":
				m.SearchStr = ""
				m.IsSearch = true
				updateSearchResult(&m)
			case "o":
				CreateSession(&m, &m.SearchStr)
				return m, tea.Quit
			}
		} else {
			switch msg.String() {
			case "ctrl+c", "enter", "esc":
				m.IsSearch = false
				return m, nil
			case "backspace":
				if len(m.SearchStr) != 0 {
					m.SearchStr = m.SearchStr[:len(m.SearchStr)-1]
				}
				updateSearchResult(&m)
				return m, nil
			default:
				m.SearchStr += msg.String()
				updateSearchResult(&m)
				return m, nil
			}
		}
	}
	if m.Toggle {
		m.BTable, cmd = m.BTable.Update(msg)
	} else {
		m.DTable, cmd = m.DTable.Update(msg)
	}
	return m, cmd
}

func updateSearchResult(m *Model) {
	if len(m.SearchStr) == 0 {
		LoadDirs(m)
		LoadBookmarks(m)
		return
	}

	rs := []table.Row{}
	for _, v := range m.Bookmarks {
		if strings.Contains(string(v), m.SearchStr) {
			rs = append(rs, table.Row{string(v)})
		}
	}
	m.BTable.SetRows(rs)
	rs = []table.Row{}
	for _, v := range m.Dirs {
		if strings.Contains(v.Dir, m.SearchStr) {
			rs = append(rs, table.Row{
				v.Dir,
				strconv.Itoa(v.Opened),
				strconv.FormatInt(v.Timestamp, 10),
			})
		}
	}
	m.DTable.SetRows(rs)
}
