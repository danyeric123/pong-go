package game

import (
	"github.com/gdamore/tcell/v2"
	"time"
)

type Game struct {
	Screen tcell.Screen
	ball   Ball
}

func NewGame(s tcell.Screen) *Game {
	width, height := s.Size()
	return &Game{
		Screen: s,
		ball:   NewBall(width, height),
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

		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
