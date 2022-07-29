package adapter

import (
	chip8 "github.com/MarceloMPJR/go-chip-8"
	"github.com/gdamore/tcell"
)

type DisplayOutput struct {
	screen *tcell.Screen
}

func NewDisplay(screen *tcell.Screen) *chip8.StandardDisplay {
	return chip8.NewStandardDisplay(
		&chip8.ConfigDisplay{Output: &DisplayOutput{screen: screen}},
	)
}

func (do *DisplayOutput) Write(p []byte) (n int, err error) {
	i, j := 0, 0

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	for _, r := range string(p) {
		if r == '\n' {
			i++
			j = 0
			continue
		}

		(*do.screen).SetContent(j, i, r, nil, style)
		j++
	}

	return len(p), nil
}
