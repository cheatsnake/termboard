package terminal

import (
	"log"
	"os"

	"github.com/cheatsnake/termboard/internal/termboard"
	"github.com/cheatsnake/termboard/internal/text"
	"github.com/gdamore/tcell/v2"
	"golang.org/x/exp/slices"
)

func MainScreen() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v\n", err)
	}

	tb := termboard.New(text.RandomText(wordsAmount))
	alphabet := text.Alphabet(tb.Lang)

	screen.SetStyle(tcell.StyleDefault)

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			render(screen, tb)
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyBackspace2 || event.Key() == tcell.KeyBackspace {
				tb.RemoveChar()
				render(screen, tb)
			} else if slices.Contains(alphabet, event.Rune()) {
				tb.WriteChar(event.Rune())
				render(screen, tb)
			}
		}
	}
}

func render(s tcell.Screen, c *termboard.Core) {
	s.Clear()

	if c.Validate() {
		c.Refresh(text.RandomText(wordsAmount))
	}

	w, h := s.Size()
	textPosX, textPosY := w/2-len(c.OriginText)/2, h/2-1

	drawText(textPosX, textPosY, c.OriginText, string(c.Input), s, tcell.StyleDefault)
	s.ShowCursor(textPosX+c.CursorPosition, textPosY)

	s.Show()
}

func drawText(x, y int, text, input string, screen tcell.Screen, style tcell.Style) {
	posX, posY := x, y
	currentStyle := style

	for i, char := range text {
		if i < len(input) {
			if rune(input[i]) == char {
				currentStyle = styleValidChar
			} else {
				currentStyle = styleWrongChar
			}
		}
		screen.SetContent(posX, posY, char, nil, currentStyle)
		currentStyle = style
		posX++
	}
}
