package app

import (
	"errors"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/talismanch1k/gorain/internal/app/rain"
)

var (
	ErrExitedScreen = errors.New("Screen was closed unexpectedly")
)

func DrawScreen() error {
	s, err := tcell.NewScreen()
	if err != nil {
		slog.Error("Failed to create screen", "err", err)
		os.Exit(1)
	}
	defer s.Fini()

	if err = s.Init(); err != nil {
		slog.Error("Failed to initialize screen", "err", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.Background(color.Default).Foreground(color.Default)
	s.SetStyle(defStyle)
	s.Clear()

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			slog.Error("Recovered from panic", "panic", maybePanic)
			panic(maybePanic)
		}
		slog.Info("Quitting app")
	}
	defer quit()

	width, height := s.Size()
	numDrops := 30
	drops := make([]*rain.Drop, numDrops)
	for i := range drops {
		drops[i] = rain.NewDrop(
			rand.Intn(width),
			-rand.Float64()*float64(height),
		)
	}

	ticker := time.NewTicker(time.Millisecond * 16)

	for {
		select {

		case ev := <-s.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return nil
				}
			}

		case <-ticker.C:
			s.Clear()

			for _, d := range drops {
				d.Fall(height)
				d.Draw(s)
			}

			s.Show()

		}
	}
	return ErrExitedScreen
}
