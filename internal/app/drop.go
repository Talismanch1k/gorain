package app

import (
	"math/rand"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

const maxLen = 3

var symbols = []rune{'|', '│', '┃', '╽', '·'}

type Drop struct {
	x      int
	y      float64
	speed  float64
	length int
	symbol rune
}

func NewDrop(x int, y float64) *Drop {
	return &Drop{
		x:      x,
		y:      y,
		speed:  min(rand.Float64()+0.4, 1),
		length: rand.Intn(maxLen),
		symbol: symbols[rand.Intn(len(symbols))],
	}
}

func (d *Drop) Fall(h, w int) {
	d.y += d.speed
	if d.y > float64(h) {
		d.y = -float64(h)
		d.x = rand.Intn(w)
	}
}

func (d *Drop) Draw(s tcell.Screen) {
	y := int(d.y)

	_, heigth := s.Size()
	if y < 0 || y > heigth {
		return
	}

	style := tcell.StyleDefault.Foreground(color.Blue)

	s.SetContent(d.x, y, d.symbol, nil, style)
}
