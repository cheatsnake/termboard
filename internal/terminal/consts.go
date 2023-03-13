package terminal

import "github.com/gdamore/tcell/v2"

const wordsAmount = 5

var (
	styleWrongChar = tcell.StyleDefault.Background(tcell.ColorIndianRed).Foreground(tcell.ColorWhite)
	styleValidChar = tcell.StyleDefault.Foreground(tcell.ColorGreen)
)
