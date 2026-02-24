package rain

import (
	"math/rand"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

type Drop struct {
	x       int
	y       float64
	speed   float64
	length  int
	symbols []rune
}

func NewDrop(x int, y float64) *Drop {
	return &Drop{
		x:       x,
		y:       y,
		speed:   rand.Float64() + 0.1,
		length:  1,
		symbols: []rune{'•', '°', '·'},
	}

}

func (d *Drop) Fall(screenHeight int) {
	d.y += d.speed
	if d.y > float64(screenHeight) {
		d.y = -float64(screenHeight) // maybe randomize?
	}
}

func (d *Drop) Draw(s tcell.Screen) {
	y := int(d.y)

	_, heigth := s.Size()
	if y < 0 || y > heigth {
		return
	}

	style := tcell.StyleDefault.Foreground(color.Blue)

	s.SetContent(d.x, y, d.symbols[0], nil, style)
}
