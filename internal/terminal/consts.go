package terminal

import "github.com/gdamore/tcell/v2"

const (
	defaultWords = 5
	defaultLang  = "en"
)

const (
	ShortcutsInfo = "Restart (Esc)  Setup (Ctrl+S)  Exit (Ctrl+C)"
)

var (
	styleWrongChar = tcell.StyleDefault.Background(tcell.ColorIndianRed).Foreground(tcell.ColorWhite)
	styleValidChar = tcell.StyleDefault.Foreground(tcell.ColorGreen)
)
