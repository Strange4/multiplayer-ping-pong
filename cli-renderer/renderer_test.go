package cli_renderer_test

import (
	"math"
	cli_renderer "multiplayer-ping-pong/cli-renderer"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRendererCursor(t *testing.T) {
	screen := tcell.NewSimulationScreen("")
	renderer := cli_renderer.Renderer{screen}
	renderer.SetCursorBlock()
	renderer.SetCursorPosition(10, 10)
	x, y, _ := screen.GetCursor()
	assert.Equal(t, 10, x, "x position is not what was expected")
	assert.Equal(t, 10, y, "y position is not what was expected")
}

func TestRendererDisplay(t *testing.T) {
	t.Run("straight horizontal line", func(t *testing.T) {
		var wantedScreen = `■         
 ■        
  ■       
   ■      
    ■      `
		wantedScreen = strings.ReplaceAll(wantedScreen, "\n", "")
		nextCharByte := 0
		getNextChar := func() rune {
			runeValue, width := utf8.DecodeRuneInString(wantedScreen[nextCharByte:])
			nextCharByte += width
			return runeValue
		}

		screen := tcell.NewSimulationScreen("")
		screen.Init()
		screen.SetSize(10, 5)
		renderer := cli_renderer.Renderer{screen}
		renderer.DrawLine(0, 0, 4, 4)
		renderer.RenderToScreen()
		contents, width, _ := screen.GetContents()

		for i, cell := range contents {
			got := string(cell.Runes)
			want := string(getNextChar())
			require.Equal(t, want, got, "At position {x: %d, y: %f}", i%width, math.Floor(float64(i)/float64(width)))
		}
	})
}
