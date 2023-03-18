package termboard

import (
	"fmt"
)

type Timer struct {
	InitialSeconds int
	CurrentSeconds int
	Output         string
}

func NewTimer(s int) *Timer {
	t := &Timer{
		InitialSeconds: s,
		CurrentSeconds: s,
	}
	t.renderOutput()
	return t
}

func (t *Timer) Tick() {
	t.CurrentSeconds--
	t.renderOutput()
}

func (t *Timer) renderOutput() {
	min := int(t.CurrentSeconds / 60)
	sec := t.CurrentSeconds % 60

	t.Output = fmt.Sprintf("%s:%s", padZero(min), padZero(sec))
}

func padZero(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprintf("%d", n)
}
