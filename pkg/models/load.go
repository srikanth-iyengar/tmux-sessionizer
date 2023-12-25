package models

import (
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	bookmarkCol = []table.Column{
		{Title: "Directory", Width: 94},
	}
	dirCol = []table.Column{
		{Title: "Directory", Width: 50},
		{Title: "Frequency", Width: 10},
		{Title: "Last Opened", Width: 30},
	}
	tableStyles = loadTableStyles()
)

func LoadBookmarks(m *Model) {
	rows := []table.Row{}
	for _, v := range m.Bookmarks {
		rows = append(rows, table.Row{string(v)})
	}
	t := table.New(
		table.WithColumns(bookmarkCol),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)
	t.SetStyles(tableStyles)
	m.BTable = t
}

func LoadDirs(m *Model) {
	rows := []table.Row{}
	for _, v := range m.Dirs {
		tm := time.Unix(v.Timestamp, 0)
		rows = append(rows, table.Row{
			v.Dir,
			strconv.Itoa(v.Opened),
			tm.String(),
		})
	}
	t := table.New(
		table.WithColumns(dirCol),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)
	t.SetStyles(tableStyles)
	m.DTable = t
}

func loadTableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	return s
}
