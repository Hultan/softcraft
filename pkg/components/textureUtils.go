package components

import (
	"fmt"

	"softcraft/pkg/types"

	"github.com/veandco/go-sdl2/sdl"
)

func drawTexture(
	tex *sdl.Texture,
	position types.Vector,
	rotation float64,
	renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture failed: %v", err)
	}

	// Convert coordinates to the top left of the sprite
	position.X -= float64(width) / 2.0
	position.Y -= float64(height) / 2.0

	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: width, H: height},
		rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
		sdl.FLIP_NONE)
}

func loadTextureFromBMP(filename string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, fmt.Errorf("loading %v failed: %v", filename, err)
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("creating texture from %v failed: %v", filename, err)
	}

	return tex, nil
}
