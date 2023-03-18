package terminal

import (
	"fmt"
	"log"
	"os"
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
	timer := termboard.NewTimer(60)
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
				screen.Clear()
				tb = termboard.New(text.RandomText(defaultWords))
				render(screen, tb)
			} else {
				if timer.CurrentSeconds < 0 {
					continue
				}

				tb.WriteChar(event.Rune())
				render(screen, tb)

				if tb.CharsTotal == 1 {
					go drawTimer(timer, screen, tb)
				}
			}
		}
	}
}

func render(s tcell.Screen, c *termboard.Core) {
	if len(c.Input) == len(c.Origin) {
		c.Refresh(text.RandomText(defaultWords))
	}

	w, h := s.Size()
	textPosX, textPosY := w/2-len(c.Origin)/2, h/2-1
	clearLine(textPosY, s)

	drawTask(textPosX, textPosY, c.Origin, string(c.Input), s, tcell.StyleDefault)
	s.ShowCursor(textPosX+c.CursorPosition, textPosY)

	s.Show()
}

func drawTask(x, y int, text, input string, s tcell.Screen, style tcell.Style) {
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

		s.SetContent(posX, y, char, nil, currentStyle)
		currentStyle = style
		posX++
	}
}

func drawText(x, y int, text string, s tcell.Screen, style tcell.Style) {
	posX := x
	for _, char := range text {
		s.SetContent(posX, y, char, nil, style)
		posX++
	}
}

func drawTimer(t *termboard.Timer, s tcell.Screen, tb *termboard.Core) {
	w, _ := s.Size()
	for t.CurrentSeconds >= 0 {
		drawText(w/2-(len(t.Output)/2), 1, t.Output, s, tcell.StyleDefault.Foreground(tcell.ColorYellow))
		s.Show()
		time.Sleep(time.Second)
		t.Tick()
	}
	drawStats(s, tb)
}

func drawStats(s tcell.Screen, tb *termboard.Core) {
	s.Clear()
	s.HideCursor()

	w, h := s.Size()
	speed := fmt.Sprintf("Speed: %.1f CPM", float64(tb.CharsTotal))
	accuracy := fmt.Sprintf("Accuracy: %.1f%s", (1-float64(tb.CharsWrong)/float64(tb.CharsTotal))*100, "%")

	drawText(0, h-1, ShortcutsInfo, s, tcell.StyleDefault.Foreground(tcell.ColorLightGray))
	drawText(w/2-(len(speed)/2), h/2, speed, s, tcell.StyleDefault)
	drawText(w/2-(len(speed)/2), h/2+1, accuracy, s, tcell.StyleDefault)

	s.Show()
}

func clearLine(y int, s tcell.Screen) {
	w, _ := s.Size()
	for i := 0; i < w; i++ {
		s.SetContent(i, y, ' ', nil, tcell.StyleDefault)
	}
}
