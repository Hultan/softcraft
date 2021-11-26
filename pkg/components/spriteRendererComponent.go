package components

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *Element
	tex           *sdl.Texture
	width, height int
}

func NewSpriteRenderer(container *Element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	sr := &spriteRenderer{}
	var err error

	sr.tex, err = loadTextureFromBMP(filename, renderer)
	if err != nil {
		panic(err)
	}

	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture failed: %v", err))
	}
	sr.width = int(width)
	sr.height = int(height)

	sr.container = container

	return sr
}

func (sr *spriteRenderer) start() {
}

func (sr *spriteRenderer) OnUpdate() error {
	return nil
}

func (sr *spriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	return drawTexture(
		sr.tex,
		sr.container.Position,
		sr.container.Rotation,
		renderer)
}

func (sr *spriteRenderer) OnCollision(_ *Element) error {
	return nil
}
