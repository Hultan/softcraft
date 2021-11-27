package components

import (
	"math/rand"

	"softcraft/pkg/common"

	"github.com/aquilax/go-perlin"
	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	common.Element

	renderer *sdl.Renderer
	window   *sdl.Window
	data     [1000][1000]int
	tex      map[int]*sdl.Texture
}

func NewWorld(renderer *sdl.Renderer, window *sdl.Window) (*World, error) {
	w := &World{}

	// Position the player in the center of the world.
	w.Element.Position = common.Vector{X: 500 * 32, Y: 500 * 32}
	w.Element.Tag = "world"
	w.renderer = renderer
	w.window = window
	w.tex = make(map[int]*sdl.Texture, 4)

	w.generateRandomWorld()

	w.loadAssets(renderer)

	mover := NewKeyboardMover(&w.Element, 5, w)
	w.AddComponent(mover)

	loc, err := NewLocationComponent(w)
	if err != nil {
		return nil, err
	}
	w.AddComponent(loc)

	return w, nil
}

func (w *World) OnUpdate() error {
	for _, comp := range w.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *World) OnDraw(renderer *sdl.Renderer) error {
	var err error

	// Upper left corner coordinates
	x := w.Element.Position.X - common.ScreenWidth/2
	y := w.Element.Position.Y - common.ScreenHeight/2
	// Upper left corner in block coordinates
	xx := int(x / common.BlockWidth)
	yy := int(y / common.BlockHeight)
	// Position inside block
	dx := x - float64(xx)*common.BlockWidth
	dy := y - float64(yy)*common.BlockHeight

drawing:
	for i := -1; i < common.CanvasWidth+1; i++ {
		for j := -1; j < common.CanvasHeight+1; j++ {
			err = common.DrawTexture(
				w.tex[w.data[xx+i][yy+j]],
				common.Vector{
					X: float64(i*common.BlockWidth) - dx,
					Y: float64(j*common.BlockHeight) - dy,
				},
				0.0,
				renderer,
				false)

			if err != nil {
				break drawing
			}
		}
	}

	// Draw components
	for _, comp := range w.Components {
		err = comp.OnDraw(renderer)
		if err != nil {
			return err
		}
	}

	return err
}

func (w *World) OnCollision(_ *common.Element) error {
	return nil
}

//
// Helper functions
//

// generateRandomWorld generates a random world
func (w *World) generateRandomWorld() {
	const (
		alpha       = 2.
		beta        = 2.
		n           = 3
		seed  int64 = 100
	)
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	var r float64
	var data int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			r = p.Noise2D(float64(i)/10, float64(j)/10)
			switch {
			case r <= 0.0:
				data = 0
			case r <= 0.25:
				data = 1
			case r <= 0.5:
				data = 2
			case r <= 0.6:
				data = 3
			case r <= 1:
				data = 4
			}
			w.data[i][j] = data
		}
	}
}

func (w *World) loadAssets(renderer *sdl.Renderer) {
	// Load assets
	w.tex[0] = w.loadAsset("assets/world/grass.bmp", renderer)
	w.tex[1] = w.loadAsset("assets/world/ground.bmp", renderer)
	w.tex[2] = w.loadAsset("assets/world/path.bmp", renderer)
	w.tex[3] = w.loadAsset("assets/world/sand.bmp", renderer)
	w.tex[4] = w.loadAsset("assets/world/water.bmp", renderer)
}

func (w *World) loadAsset(fileName string, renderer *sdl.Renderer) *sdl.Texture {
	t, err := common.LoadTextureFromBMP(fileName, renderer)
	if err != nil {
		// No point in continue if we can't load the assets
		panic(err)
	}
	return t
}
