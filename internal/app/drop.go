package app

import (
	"math/rand"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

const maxLen = 5

var speed = interval[float64]{min: 0.4, max: 1}

type interval[T any] struct {
	min T
	max T
}

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
		speed:  min(rand.Float64()+speed.min, speed.max),
		length: rand.Intn(maxLen),
		symbol: symbols[rand.Intn(len(symbols))],
	}
}

func (d *Drop) Fall(h, w int) {
	d.y += d.speed
	if d.y > float64(h) {
		d.y = -rand.Float64() * float64(2*h)
		d.x = rand.Intn(w)

		d.speed = min(rand.Float64()+speed.min, speed.max)
		d.length = rand.Intn(maxLen)
		d.symbol = symbols[rand.Intn(len(symbols))]
	}
}

func (d *Drop) Draw(s tcell.Screen) {
	_, heigth := s.Size()

	for i := range d.length {
		ty := int(d.y) - i

		if ty < 0 || ty > heigth {
			return
		}

		fade := 1.0 - float64(i)/float64(d.length)
		intensity := remapInterval(d.speed, speed, interval[float64]{0.0, 1.0})

		r := int32(0)
		g := int32((40 + intensity*120) * fade)
		b := int32((20 + intensity*235) * fade)

		style := tcell.StyleDefault.Foreground(color.NewRGBColor(r, g, b))

		if i == 0 {
			style = style.Bold(true)
		}

		s.SetContent(d.x, ty, d.symbol, nil, style)
	}
}

func remapInterval(val float64, old, new interval[float64]) float64 {
	return (val-old.min)*(new.max-new.min)/(old.max-old.min) + new.min
}
