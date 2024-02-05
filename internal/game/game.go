package game

import (
	"github.com/gdamore/tcell/v2"
	"time"
)

type Game struct {
	Screen  tcell.Screen
	ball    Ball
	Player1 Paddle
	Player2 Paddle
}

func NewGame(s tcell.Screen) *Game {
	width, height := s.Size()
	return &Game{
		Screen:  s,
		ball:    NewBall(width, height),
		Player1: NewPaddle(5, 3, 3),
		Player2: NewPaddle(width-5, 3, 3),
	}
}

func (g *Game) Run() {

	g.Screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack))

	for {
		g.Screen.Clear()

		g.ball.CheckEdges(g.Screen.Size())

		g.ball.Move()

		g.Screen.SetContent(g.ball.X, g.ball.Y, g.ball.Display(), nil, tcell.StyleDefault)

		paddleStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

		// repeat this for Player 2
		drawSprite(g.Screen,
			g.Player1.X,
			g.Player1.Y,
			g.Player1.X+g.Player1.width,
			g.Player1.Y+g.Player1.height,
			paddleStyle,
			g.Player1.Display())

		drawSprite(g.Screen,
			g.Player2.X,
			g.Player2.Y,
			g.Player2.X+g.Player2.width,
			g.Player2.Y+g.Player2.height,
			paddleStyle,
			g.Player2.Display())

		if g.ball.intersects(g.Player1) || g.ball.intersects(g.Player2) {
			g.ball.reverseX()
			g.ball.reverseY()
		}

		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}

func drawSprite(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
