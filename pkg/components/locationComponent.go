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
	if lc.font, err = ttf.OpenFont("assets/nerdfont.ttf", 20); err != nil {
		return nil, err
	}

	return lc, nil
}

func (lc *locationComponent) OnDraw(renderer *sdl.Renderer) error {
	var err error
	var surface *sdl.Surface

	x:= int64(lc.world.Position.X/common.BlockWidth)
	y:= int64(lc.world.Position.Y/common.BlockHeight)
	// Create a white text with the font
	message := fmt.Sprintf("Pos : %d,%d", x,y)
	if surface, err = lc.font.RenderUTF8Solid(message, sdl.Color{R: 0, G: 0, B: 0, A: 255}); err != nil {
		return err
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: surface.W+10, H: surface.H})
	renderer.Copy(texture,
		&sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H},
		&sdl.Rect{X: 5, Y: 0, W: surface.W, H: surface.H})

	renderer.Present()

	return nil
}

func (lc *locationComponent) OnUpdate() error {
	return nil
}

func (lc *locationComponent) OnCollision(_ *common.Element) error {
	return nil
}
