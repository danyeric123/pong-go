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
		XDirection: []int{1, -1}[rand.Intn(2)],
		YDirection: []int{1, -1}[rand.Intn(2)],
	}
}

func (b *Ball) Display() rune {
	return '\u25CF'
}

func (b *Ball) Move() {
	b.X += b.XDirection
	b.Y += b.YDirection
}

func (b *Ball) reverseX() {
	b.XDirection *= -1
}

func (b *Ball) reverseY() {
	b.YDirection *= -1
}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
		b.reverseX()
	}

	if b.Y <= 0 || b.Y >= maxHeight {
		b.reverseY()
	}
}

func (b *Ball) intersects(p Paddle) bool {
	return b.X >= p.X && b.X <= p.X+p.width && b.Y >= p.Y && b.Y <= p.Y+p.height
}
