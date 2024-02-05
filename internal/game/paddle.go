package game

import (
	"strings"
)

type Paddle struct {
	X, Y, Yspeed  int
	width, height int
}

func NewPaddle(x, y, yspeed int) Paddle {
	return Paddle{
		X:      x,
		Y:      y,
		Yspeed: yspeed,
		width:  1,
		height: 6,
	}
}

func (p *Paddle) Display() string {
	return strings.Repeat(" ", p.height)
}

func (p *Paddle) MoveUp() {
	if p.Y > 0 {
		p.Y -= p.Yspeed
	}
}

func (p *Paddle) MoveDown(windowHeight int) {
	if p.Y < windowHeight-p.height {
		p.Y += p.Yspeed
	}
}
