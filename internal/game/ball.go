package game

import (
	"math/rand"
)

type Ball struct {
	X, Y                   int
	XDirection, YDirection int
}

func NewBall(maxWidth, maxHeight int) Ball {
	return Ball{
		X:          rand.Intn(maxWidth-5) + 5,
		Y:          rand.Intn(maxHeight-5) + 5,
		XDirection: []int{1,-1}[rand.Intn(2)],
		YDirection: []int{1,-1}[rand.Intn(2)],
	}
}

func (b *Ball) Display() rune {
	return '\u25CF'
}

func (b *Ball) Move() {
	b.X += b.XDirection
	b.Y += b.YDirection
}


func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
			b.XDirection *= -1
	}

	if b.Y <= 0 || b.Y >= maxHeight {
			b.YDirection *= -1
	}
}