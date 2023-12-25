package models

import (
	"github.com/charmbracelet/bubbles/table"
)

func (m Model) View() string {
	var t table.Model
	if m.Toggle {
		t = m.BTable
	} else {
		t = m.DTable
	}
	return Logo() +
		m.GetSearchStr() +
		t.View()
}
