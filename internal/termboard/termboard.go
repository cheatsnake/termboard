package termboard

type Core struct {
	Origin         string
	Input          []rune
	CharsWrong     int
	CharsTotal     int
	CursorPosition int
}

func New(text string) *Core {
	return &Core{
		Origin:         text,
		Input:          make([]rune, 0, len(text)),
		CharsWrong:     0,
		CharsTotal:     0,
		CursorPosition: 0,
	}
}

func (c *Core) WriteChar(char rune) {
	// Input capacity if full
	if len(c.Input) >= len(c.Origin) {
		return
	}

	c.Input = append(c.Input, char)
	c.MoveRight()

	if char != rune(c.Origin[len(c.Input)-1]) {
		c.CharsWrong++
	} else {
		c.CharsTotal++
	}
}

func (c *Core) RemoveChar() {
	if len(c.Input) > 0 {
		c.Input = c.Input[:len(c.Input)-1]
		c.CursorPosition = len(c.Input)
	}
}

func (c *Core) MoveLeft() {
	if c.CursorPosition > 0 {
		c.CursorPosition--
	}
}

func (c *Core) MoveRight() {
	if c.CursorPosition < len(c.Input) {
		c.CursorPosition++
	}
}

func (c *Core) Refresh(origin string) {
	c.CursorPosition = 0
	c.Input = nil
	c.Origin = origin
}
