package components

import (
	"softcraft/pkg/assetManager"
	"softcraft/pkg/common"
	"softcraft/pkg/world"

	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	*common.Element

	data   [][]assetManager.AssetMap
	assets *assetManager.AssetManager
}

// NewWorld creates a new world
func NewWorld(am *assetManager.AssetManager) (*World, error) {
	w := &World{}

	// Position the player in the center of the world.
	w.Element = &common.Element{
		Position: common.Vector{X: 45 * 32, Y: 30 * 32},
		Tag:      "world",
		Active:   true,
	}

	gen := world.Loader{}
	w.data = gen.LoadWorld()
	w.assets = am

	mover := NewKeyboardMover(w.Element, 0.35, w)
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
	startX := w.Element.Position.X - common.ScreenWidth/2
	startY := w.Element.Position.Y - common.ScreenHeight/2
	// Upper left corner in block coordinates
	xx := int(startX / common.BlockWidth)
	yy := int(startY / common.BlockHeight)
	// Position inside block
	dx := startX - float64(xx)*common.BlockWidth
	dy := startY - float64(yy)*common.BlockHeight

drawing:
	for y := 0; y < common.CanvasHeight; y++ {
		for x := 0; x < common.CanvasWidth; x++ {
			if yy+y < 0 || yy+y >= len(w.data) {
				continue
			}
			if xx+x < 0 || xx+x >= len(w.data[yy+y]) {
				continue
			}
			err = common.DrawTexture(
				w.assets.GetWorldAsset(w.data[yy+y][xx+x]),
				common.Vector{
					X: float64(x*common.BlockWidth) - dx,
					Y: float64(y*common.BlockHeight) - dy,
				},
				0.0,
				renderer,
				false,
			)

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
