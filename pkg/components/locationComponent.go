package components

import (
	"fmt"

	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type locationComponent struct {
	world   *World
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

	// Load the font for our text
	if lc.font, err = ttf.OpenFont("assets/nerdfont.ttf", 20); err != nil {
		return nil, err
	}

	return lc, nil
}

func (lc *locationComponent) OnDraw(renderer *sdl.Renderer) error {
	// Get player position
	x, y := lc.getPosition()

	// Draw a black position information
	message := fmt.Sprintf("Pos : %d,%d", x, y)
	surface, err := lc.font.RenderUTF8Blended(message, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		return err
	}
	defer surface.Free()

	// Create a texture for the text position message
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer func(texture *sdl.Texture) {
		err = texture.Destroy()
	}(texture)

	// Draw a white box, where the position message should be shown
	err = renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: surface.W + 10, H: surface.H})
	if err != nil {
		return err
	}

	// Copy the position texture to the screen
	err = renderer.Copy(texture,
		&sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H},
		&sdl.Rect{X: 5, Y: 0, W: surface.W, H: surface.H})
	if err != nil {
		return err
	}

	return nil
}

func (lc *locationComponent) OnUpdate() error {
	return nil
}

func (lc *locationComponent) OnCollision(_ *common.Element) error {
	return nil
}

func (lc *locationComponent) getPosition() (int64, int64) {
	x := int64(lc.world.Position.X / common.BlockWidth)
	y := int64(lc.world.Position.Y / common.BlockHeight)
	return x, y
}
