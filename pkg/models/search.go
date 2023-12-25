package models

import "fmt"

func (m Model) GetSearchStr() string {
	return fmt.Sprintf(`
󰈲 %s
`, m.SearchStr)
}
