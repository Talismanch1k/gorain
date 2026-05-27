// Package app render/draw drops on screen
package app

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

var ErrExitedScreen = errors.New("screen was closed unexpectedly")

func DrawScreen() error {
	s, err := tcell.NewScreen()
	if err != nil {
		return fmt.Errorf("create screen: %w", err)
	}
	defer s.Fini()

	if err = s.Init(); err != nil {
		return fmt.Errorf("initialize screen: %w", err)
	}

	defStyle := tcell.StyleDefault.Background(color.Default).Foreground(color.Default)
	s.SetStyle(defStyle)
	s.Clear()

	width, height := s.Size()
	numDrops := (height + width) / 2
	drops := make([]*Drop, numDrops)

	for i := range drops {
		drops[i] = NewDrop(
			rand.Intn(width),
			-rand.Float64()*float64(2*height),
		)
	}

	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		case ev := <-s.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
				width, height = s.Size()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return nil
				}
			}

		case <-ticker.C:
			s.Clear()

			for _, d := range drops {
				d.Fall(height, width)
				d.Draw(s)
			}

			s.Show()
		}
	}
}
