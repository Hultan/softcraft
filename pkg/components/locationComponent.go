package components

import (
	"fmt"

	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type locationComponent struct {
	world   *World
	window  *sdl.Window
	surface *sdl.Surface
	font    *ttf.Font
}

func NewLocationComponent(world *World) (*locationComponent, error) {
	var err error
	lc := &locationComponent{
		world:  world,
	}

	if err = ttf.Init(); err != nil {
		return nil, err
	}

	lc.window = world.window

	if lc.surface, err = world.window.GetSurface(); err != nil{
		return nil, err
	}
	// Load the font for our text
	if lc.font, err = ttf.OpenFont("assets/nerdfont.ttf", 15); err != nil {
		return nil, err
	}

	return lc, nil
}

func (lc *locationComponent) OnDraw(_ *sdl.Renderer) error {
	var err error
	var text *sdl.Surface

	x:= int64(lc.world.Position.X/common.BlockWidth)
	y:= int64(lc.world.Position.Y/common.BlockHeight)
	// Create a white text with the font
	message := fmt.Sprintf("Pos : %d,%d", x,y)
	if text, err = lc.font.RenderUTF8Blended(message, sdl.Color{R: 255, G: 255, B: 255, A: 255}); err != nil {
		return err
	}
	defer text.Free()

	// Draw the text around the center of the window
	if err = text.Blit(nil, lc.surface, &sdl.Rect{X: 10, Y: 10, W: 0, H: 0}); err != nil {
		return err
	}

	// Update the window surface with what we have drawn
	lc.window.UpdateSurfaceRects([]sdl.Rect{sdl.Rect{X: 10, Y: 10, W: text.W, H: text.H}})

	return nil
}

func (lc *locationComponent) OnUpdate() error {
	return nil
}

func (lc *locationComponent) OnCollision(_ *common.Element) error {
	return nil
}
