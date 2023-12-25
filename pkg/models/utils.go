package models

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"

	"github.com/charmbracelet/bubbles/table"
	"gopkg.in/yaml.v2"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetWidth() uint {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(ws.Col)
}

func UpdateHistory(m *Model, dir string) {
	idx := -1
	for k, v := range m.Dirs {
		if v.Dir == dir {
			idx = k
			break
		}
	}
	if idx >= 0 {
		m.Dirs[idx].Timestamp = time.Now().Unix()
		m.Dirs[idx].Opened += 1
	} else {
		m.Dirs = append(m.Dirs, Dir{
			Dir:       dir,
			Timestamp: time.Now().Unix(),
			Opened:    1,
		})
	}
	bytes, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(m.ConfigStr, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func CreateSession(m *Model, dir *string) bool {
	var selected table.Row
	if m.Toggle {
		selected = m.BTable.SelectedRow()
	} else {
		selected = m.DTable.SelectedRow()
	}
	if dir == nil {
		dir = &[]string(selected)[0]
	}
	sessionName := filepath.Base(*dir)
	cmd := exec.Command("tmux", "new-session", "-d", "-c", *dir, "-s", sessionName)

	if err := cmd.Run(); err != nil {
		return false
	} else {
		UpdateHistory(m, *dir)
		return true
	}
}
