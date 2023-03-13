package main

import (
	"math/rand"
	"time"

	"github.com/cheatsnake/termboard/internal/terminal"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	terminal.MainScreen()
}
