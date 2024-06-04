package renderer_test

import (
	"multiplayer-ping-pong/renderer"
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRendererCursor(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	r := renderer.NewRenderer(screen)
	r.SetCursorPosition(10, 10)
	x, y, _ := screen.GetCursor()
	assert.Equal(t, 10, x, "x position is not what was expected")
	assert.Equal(t, 10, y, "y position is not what was expected")
}

func TestLineDrawer(t *testing.T) {

	requireFilledOrEmpty := func(got rune, shouldFill bool, index, width int) {
		x, y := index%width, index/width
		errorMessage := "unexpeted rune at position {x: %d, y: %d}"
		var want string
		if shouldFill {
			want = string(renderer.FilledCellChar)
		} else {
			want = string(renderer.EmptyCellCHar)
		}
		require.Equal(t, want, string(got), errorMessage, x, y)

	}

	t.Run("Horizontal line", func(t *testing.T) {
		screen := tcell.NewSimulationScreen("")
		r := renderer.NewRenderer(screen)
		const lineLength = 10

		r.DrawLine(0, 0, lineLength, 0)
		r.RenderToScreen()

		contents, width, _ := screen.GetContents()

		// the first cells should be filled
		for i := 0; i < lineLength; i++ {
			got := contents[i].Runes[0]
			requireFilledOrEmpty(got, true, i, width)
		}

		// the rest should be empty
		for i := lineLength + 1; i < len(contents); i++ {
			got := contents[i].Runes[0]
			requireFilledOrEmpty(got, false, i, width)
		}
	})

	t.Run("Diagonal line", func(t *testing.T) {
		screen := tcell.NewSimulationScreen("")
		r := renderer.NewRenderer(screen)

		const lineLength = 10

		r.DrawLine(0, 0, lineLength, lineLength)
		r.RenderToScreen()

		contents, width, _ := screen.GetContents()
		nextExpectedCell := 0
		for i := 0; i < lineLength*lineLength; i++ {
			got := contents[i].Runes[0]

			// when it hits  the diagonal
			if i == nextExpectedCell {
				requireFilledOrEmpty(got, true, i, width)

				nextExpectedCell += width + 1
			} else {
				requireFilledOrEmpty(got, false, i, width)
			}
		}
	})

	t.Run("Vertical line", func(t *testing.T) {
		screen := tcell.NewSimulationScreen("")
		r := renderer.NewRenderer(screen)

		const lineLength = 10

		r.DrawLine(0, 0, 0, lineLength)
		r.RenderToScreen()

		contents, width, _ := screen.GetContents()
		for i := 0; i < lineLength*lineLength; i++ {
			got := contents[i].Runes[0]

			// when it hits  the diagonal
			if i%width == 0 {
				requireFilledOrEmpty(got, true, i, width)
			} else {
				requireFilledOrEmpty(got, false, i, width)
			}
		}
	})
}
