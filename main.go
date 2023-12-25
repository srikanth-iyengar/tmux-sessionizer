package main

import (
	"flag"
	"fmt"

	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/srikanth-iyengar/tmux-sessionizer/pkg/models"
	"gopkg.in/yaml.v2"
)

var (
	FlConf     = flag.String("c", "config.yml", "path to the config/store file")
	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#556B2F"))
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
)

func initialModel() models.Model {
	config, err := os.ReadFile(*FlConf)
	if err != nil {
		fmt.Println(errorStyle.Render(err.Error()))
		os.Exit(1)
	}
	var m models.Model
	err = yaml.Unmarshal(config, &m)
	if err != nil {
		fmt.Println(errorStyle.Render(err.Error()))
		os.Exit(1)
	}
	models.LoadDirs(&m)
	models.LoadBookmarks(&m)
	m.Toggle = false
	m.SearchStr = ""
	return m
}

func main() {
	flag.Parse()
	m := initialModel()
	m.ConfigStr = *FlConf
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println(errorStyle.Render(fmt.Sprintln("Error in starting program:", err)))
		os.Exit(1)
	}
}
