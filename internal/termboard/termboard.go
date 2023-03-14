package termboard

type Core struct {
	OriginText     string
	Input          []rune
	CursorPosition int
}

func New(text string) *Core {
	c := Core{
		OriginText:     text,
		CursorPosition: 0,
		Input:          make([]rune, 0, len(text)),
	}

	return &c
}

func (c *Core) WriteChar(char rune) {
	if len(c.Input) < len(c.OriginText) {
		c.Input = append(c.Input, char)
		c.MoveRight()
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
	c.OriginText = origin
}

func (c *Core) Validate() bool {
	return (len(c.Input) == len(c.OriginText)) && string(c.Input) == c.OriginText
}
