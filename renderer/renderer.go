package renderer

import (
	"github.com/gdamore/tcell/v2"
)

const FilledCellChar = '■'
const EmptyCellCHar = ' '

type Renderer struct {
	screen tcell.Screen
}

func NewRenderer(s tcell.Screen) Renderer {
	s.Init()
	s.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	return Renderer{s}
}

func (r *Renderer) SetCursorPosition(x, y int) {
	r.screen.ShowCursor(x, y)
}

func (r *Renderer) RenderToScreen() {
	r.screen.Show()
}

func (r *Renderer) DrawLine(startX, startY, endX, endY int) {
	dx := fastAbsDiff(startX, endX)
	dy := -fastAbsDiff(startY, endY)

	var stepX, stepY int
	if startX < endX {
		stepX = 1
	} else {
		stepX = -1
	}
	if startY < endY {
		stepY = 1
	} else {
		stepY = -1
	}

	errorRate := dx + dy
	x := startX
	y := startY

	for {
		r.screen.SetContent(x, y, FilledCellChar, nil, tcell.StyleDefault)
		if x == endX && y == endY {
			break
		}
		nextError := 2 * errorRate
		if nextError >= dy {
			if x == endX {
				break
			}
			errorRate += dy
			x += stepX
		}
		if nextError <= dx {
			if y == endY {
				break
			}
			errorRate += dx
			y += stepY
		}
	}
}

func fastAbsDiff(a, b int) int {
	diff := b - a
	if diff < 0 {
		return -diff
	}
	return diff
}
