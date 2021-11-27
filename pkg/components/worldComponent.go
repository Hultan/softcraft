package components

import (
	"softcraft/pkg/common"
	"softcraft/pkg/world"

	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	common.Element

	renderer *sdl.Renderer
	window   *sdl.Window
	data     [1000][1000]world.Asset
	tex      map[world.Asset]*sdl.Texture
}

func NewWorld(renderer *sdl.Renderer, window *sdl.Window) (*World, error) {
	w := &World{}

	// Position the player in the center of the world.
	w.Element.Position = common.Vector{X: 500 * 32, Y: 500 * 32}
	w.Element.Tag = "world"
	w.renderer = renderer
	w.window = window
	w.tex = make(map[world.Asset]*sdl.Texture, 4)

	gen := world.Generator{}
	w.data = gen.GenerateRandomWorld()

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

func (w *World) loadAssets(renderer *sdl.Renderer) {
	// Load assets
	w.tex[world.AssetGrass] = w.loadAsset("assets/world/grass.bmp", renderer)
	w.tex[world.AssetGround] = w.loadAsset("assets/world/ground.bmp", renderer)
	w.tex[world.AssetPath] = w.loadAsset("assets/world/path.bmp", renderer)
	w.tex[world.AssetSand] = w.loadAsset("assets/world/sand.bmp", renderer)
	w.tex[world.AssetWater] = w.loadAsset("assets/world/water.bmp", renderer)
}

func (w *World) loadAsset(fileName string, renderer *sdl.Renderer) *sdl.Texture {
	t, err := common.LoadTextureFromBMP(fileName, renderer)
	if err != nil {
		// No point in continue if we can't load the assets
		panic(err)
	}
	return t
}
