package main

import (
	"github.com/danyeric123/pong-go/internal/game"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

func main() {
	s, e := tcell.NewScreen()

	if e != nil {
		log.Fatal(e)
	}

	if e = s.Init(); e != nil {
		log.Fatal(e)
	}
	g := game.NewGame(s)
	go g.Run()

	switch event := s.PollEvent().(type) {
	case *tcell.EventResize:
		g.Screen.Sync()
	case *tcell.EventKey:
		if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
			s.Fini()
			os.Exit(0)
		}
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
