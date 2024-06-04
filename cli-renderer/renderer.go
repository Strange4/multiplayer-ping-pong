package cli_renderer

import (
	"github.com/gdamore/tcell/v2"
)

const cellStyle = '■'

type Renderer struct {
	Screen tcell.Screen
}

func (r *Renderer) SetCursorBlock() {
	r.Screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
}

func (r *Renderer) SetCursorPosition(x, y int) {
	r.Screen.ShowCursor(x, y)
}

func (r *Renderer) RenderToScreen() {
	r.Screen.Show()
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
		r.Screen.SetContent(x, y, cellStyle, nil, tcell.StyleDefault)
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
