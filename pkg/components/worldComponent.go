package components

import (
	"softcraft/pkg/assetManager"
	"softcraft/pkg/common"
	"softcraft/pkg/world"

	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	common.Element

	renderer *sdl.Renderer
	window   *sdl.Window
	data     [][]assetManager.AssetNumeric
	assets      *assetManager.AssetManager
}

// NewWorld creates a new world
func NewWorld(renderer *sdl.Renderer, window *sdl.Window, am *assetManager.AssetManager) (*World, error) {
	w := &World{}

	// Position the player in the center of the world.
	w.Element.Position = common.Vector{X: 45 * 32, Y: 30 * 32}
	w.Element.Tag = "world"
	w.renderer = renderer
	w.window = window

	gen := world.Loader{}
	w.data = gen.LoadWorld()
	w.assets = am

	mover := NewKeyboardMover(&w.Element, 0.1, w)
	w.AddComponent(mover)

	loc, err := NewLocationComponent(w)
	if err != nil {
		return nil, err
	}
	w.AddComponent(loc)

	return w, nil
}

// OnUpdate updates the world
func (w *World) OnUpdate() error {
	for _, comp := range w.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

// OnDraw draws the world
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
	for i := 0; i < common.CanvasWidth; i++ {
		for j := 0; j < common.CanvasHeight; j++ {
			if yy+j<0 || yy+j>=len(w.data) {
				continue
			}
			if xx+i<0 || xx+i>=len(w.data[yy+j]) {
				continue
			}
			err = common.DrawTexture(
				w.assets.GetWorldAsset(w.data[yy+j][xx+i]),
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

// OnCollision handles collisions
func (w *World) OnCollision(_ *common.Element) error {
	return nil
}

