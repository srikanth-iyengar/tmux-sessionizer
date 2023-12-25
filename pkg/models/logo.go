package models

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	art = ` _______ __  __ _    ___   __   _____ ______  _____ _____ _____ ____  _   _ _____ ____________ _____
 |__   __|  \/  | |  | \ \ / /  / ____|  ____|/ ____/ ____|_   _/ __ \| \ | |_   _|___  /  ____|  __ \
    | |  | \  / | |  | |\ V /  | (___ | |__  | (___| (___   | || |  | |  \| | | |    / /| |__  | |__) |
    | |  | |\/| | |  | | > <    \___ \|  __|  \___ \\___ \  | || |  | | .   | | |   / / |  __| |  _  /
    | |  | |  | | |__| |/ . \   ____) | |____ ____) |___) |_| || |__| | |\  |_| |_ / /__| |____| | \ \
    |_|  |_|  |_|\____//_/ \_\ |_____/|______|_____/_____/|_____\____/|_| \_|_____/_____|______|_|  \_\
    `
)

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

func Logo() string {
	lines := strings.Split(art, "\n")
	maxLength := 0
	for _, line := range lines {
		if l := len(line); l > maxLength {
			maxLength = l
		}
	}
	colors := colorGrid(maxLength, len(lines))

	var s strings.Builder

	for row, line := range lines {
		for col, char := range line {
			s.WriteString(lipgloss.NewStyle().SetString(string(char)).Foreground(lipgloss.Color(colors[row][col])).String())
		}
		s.WriteRune('\n')
	}
	return s.String()
}
