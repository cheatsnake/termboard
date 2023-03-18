package terminal

import "github.com/gdamore/tcell/v2"

const (
	defaultWords  int     = 5
	defaultLang   string  = "en"
	avgWordLength float64 = 4.7
)

const (
	ShortcutsInfo = "Restart (Esc)   Exit (Ctrl+C)"
)

var (
	styleWrongChar = tcell.StyleDefault.Background(tcell.ColorIndianRed).Foreground(tcell.ColorWhite)
	styleValidChar = tcell.StyleDefault.Foreground(tcell.ColorGreen)
)
