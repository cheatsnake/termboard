package terminal

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/cheatsnake/termboard/internal/termboard"
	"github.com/cheatsnake/termboard/internal/text"
	"github.com/gdamore/tcell/v2"
)

func MainScreen() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v\n", err)
	}

	tb := termboard.New(text.RandomText(defaultWords))
	screen.SetStyle(tcell.StyleDefault)

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			render(screen, tb)

		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyBackspace2 || event.Key() == tcell.KeyBackspace {
				tb.RemoveChar()
				render(screen, tb)
			} else if event.Key() == tcell.KeyEscape {
				tb.Refresh(text.RandomText(defaultWords))
				render(screen, tb)
			} else {
				tb.WriteChar(event.Rune())
				render(screen, tb)
			}
		}
	}
}

func render(s tcell.Screen, c *termboard.Core) {
	s.Clear()

	if c.Validate() {
		c.Refresh(text.RandomText(defaultWords))
	}

	w, h := s.Size()
	textPosX, textPosY := w/2-len(c.OriginText)/2, h/2-1

	if len(c.Input) < 1 {
		drawText(1, h-1, ShortcutsInfo, s, tcell.StyleDefault.Foreground(tcell.ColorDimGray))
	}

	drawTask(textPosX, textPosY, c.OriginText, string(c.Input), s, tcell.StyleDefault)
	s.ShowCursor(textPosX+c.CursorPosition, textPosY)

	s.Show()
}

func drawTask(x, y int, text, input string, screen tcell.Screen, style tcell.Style) {
	posX := x
	currentStyle := style

	for i, char := range text {
		if i < len(input) {
			if rune(input[i]) == char {
				currentStyle = styleValidChar
			} else {
				currentStyle = styleWrongChar
			}
		}

		screen.SetContent(posX, y, char, nil, currentStyle)
		currentStyle = style
		posX++
	}
}

func drawText(x, y int, text string, screen tcell.Screen, style tcell.Style) {
	posX := x
	for _, char := range text {
		screen.SetContent(posX, y, char, nil, style)
		posX++
	}
}

func drawTimer(sec int, screen tcell.Screen) {
	count := sec - 1
	w, _ := screen.Size()
	for count >= 0 {
		drawText(w/2-1, 1, strconv.Itoa(count), screen, tcell.StyleDefault.Foreground(tcell.ColorYellow))
		screen.Show()
		time.Sleep(time.Second)
		count--
	}
}
