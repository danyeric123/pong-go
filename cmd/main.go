package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danyeric123/pong-go/internal/game"
	"github.com/gdamore/tcell/v2"
)

func main() {
	s, e := tcell.NewScreen()

	if e != nil {
		log.Fatal(e)
	}

	if e = s.Init(); e != nil {
		log.Fatal(e)
	}

	_, height := s.Size() // get the size of the terminal

	g := game.NewGame(s)
	go g.Run()

	switch event := s.PollEvent().(type) {
	case *tcell.EventResize:
		g.Screen.Sync()
	case *tcell.EventKey:
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
			s.Fini()
			os.Exit(0)
		} else if event.Key() == tcell.KeyUp {
			g.Player2.MoveUp()
		} else if event.Key() == tcell.KeyDown {
			g.Player2.MoveDown(height)
		} else if event.Rune() == 'w' {
			g.Player1.MoveUp()
		} else if event.Rune() == 's' {
			g.Player1.MoveDown(height)
		}
	default:
		fmt.Println("Unknown event:", event)
	}

	s.Clear()

	s.Show()
	s.Sync()

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}
