package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// ============================================================
// DEMO VARIABLE - Change this color to show fast recompile!
// ============================================================
var ornamentColor = color.RGBA{0, 255, 255, 255} // RED - try: {50, 100, 255, 255} for BLUE
// ============================================================

const (
	screenWidth  = 400
	screenHeight = 500
)

var (
	treeColor  = color.RGBA{34, 139, 34, 255}   // Forest green
	trunkColor = color.RGBA{139, 69, 19, 255}   // Saddle brown
	starColor  = color.RGBA{255, 215, 0, 255}   // Gold
	snowColor  = color.RGBA{255, 255, 255, 200} // White
)

type Snowflake struct {
	x, y  float64
	speed float64
	size  float64
}

type Game struct {
	snowflakes []Snowflake
	twinkle    float64
}

func NewGame() *Game {
	g := &Game{
		snowflakes: make([]Snowflake, 100),
	}
	for i := range g.snowflakes {
		g.snowflakes[i] = Snowflake{
			x:     rand.Float64() * screenWidth,
			y:     rand.Float64() * screenHeight,
			speed: 0.5 + rand.Float64()*1.5,
			size:  1 + rand.Float64()*2,
		}
	}
	return g
}

func (g *Game) Update() error {
	// Update snowflakes
	for i := range g.snowflakes {
		g.snowflakes[i].y += g.snowflakes[i].speed
		g.snowflakes[i].x += math.Sin(g.snowflakes[i].y/30) * 0.5
		if g.snowflakes[i].y > screenHeight {
			g.snowflakes[i].y = 0
			g.snowflakes[i].x = rand.Float64() * screenWidth
		}
	}
	// Twinkle effect for ornaments
	g.twinkle += 0.1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Dark blue night sky
	screen.Fill(color.RGBA{10, 10, 40, 255})

	// Draw snowflakes (behind tree)
	for _, sf := range g.snowflakes {
		drawCircle(screen, int(sf.x), int(sf.y), int(sf.size), snowColor)
	}

	// Draw tree (triangle made of 3 layers)
	centerX := screenWidth / 2

	// Layer 1 (top) - smaller triangle
	drawTriangle(screen, centerX, 80, 60, 80, treeColor)
	// Layer 2 (middle)
	drawTriangle(screen, centerX, 130, 90, 100, treeColor)
	// Layer 3 (bottom) - largest
	drawTriangle(screen, centerX, 190, 120, 120, treeColor)

	// Draw trunk
	trunkWidth := 40
	trunkHeight := 60
	for y := 310; y < 310+trunkHeight; y++ {
		for x := centerX - trunkWidth/2; x < centerX+trunkWidth/2; x++ {
			screen.Set(x, y, trunkColor)
		}
	}

	// Draw ornaments with twinkle effect
	ornaments := []struct{ x, y int }{
		{centerX - 25, 120},
		{centerX + 20, 135},
		{centerX - 40, 175},
		{centerX + 35, 185},
		{centerX, 160},
		{centerX - 55, 240},
		{centerX + 50, 250},
		{centerX - 20, 220},
		{centerX + 25, 280},
		{centerX - 45, 290},
	}

	for i, o := range ornaments {
		// Add slight twinkle/glow variation
		brightness := 0.8 + 0.2*math.Sin(g.twinkle+float64(i)*0.5)
		c := color.RGBA{
			uint8(float64(ornamentColor.R) * brightness),
			uint8(float64(ornamentColor.G) * brightness),
			uint8(float64(ornamentColor.B) * brightness),
			255,
		}
		drawCircle(screen, o.x, o.y, 8, c)
		// Small highlight
		drawCircle(screen, o.x-2, o.y-2, 2, color.RGBA{255, 255, 255, 150})
	}
}

func drawTriangle(screen *ebiten.Image, cx, topY, halfWidth, height int, c color.Color) {
	for y := 0; y < height; y++ {
		// Width increases as we go down
		w := int(float64(halfWidth) * float64(y) / float64(height))
		for x := cx - w; x <= cx+w; x++ {
			screen.Set(x, topY+y, c)
		}
	}
}

func drawCircle(screen *ebiten.Image, cx, cy, r int, c color.Color) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				screen.Set(cx+x, cy+y, c)
			}
		}
	}
}

func drawStar(screen *ebiten.Image, cx, cy, size int, c color.Color) {
	// Simple 4-pointed star
	for i := -size; i <= size; i++ {
		screen.Set(cx+i, cy, c)
		screen.Set(cx, cy+i, c)
		// Diagonal arms (smaller)
		if i >= -size/2 && i <= size/2 {
			screen.Set(cx+i, cy+i, c)
			screen.Set(cx+i, cy-i, c)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Christmas Tree")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
