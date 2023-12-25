package models

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Dir struct {
	Dir       string `yaml:"dir"`
	Timestamp int64  `yaml:"timestamp"`
	Opened    int    `yaml:"opened"`
}

type Bookmark string

type Model struct {
	Bookmarks []Bookmark  `yaml:"bookmarks"`
	Dirs      []Dir       `yaml:"dirs"`
	BTable    table.Model `yaml:"-"`
	DTable    table.Model `yaml:"-"`
	Toggle    bool        `yaml:"-"`
	IsSearch  bool        `yaml:"-"`
	SearchStr string      `yaml:"-"`
	STable    table.Model `yaml:"-"`
	ConfigStr string      `yaml:"-"`
}

func (m Model) Init() tea.Cmd {
	return nil
}
