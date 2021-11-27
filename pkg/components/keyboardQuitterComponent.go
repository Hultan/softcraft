package components

import (
	"os"

	"softcraft/pkg/common"

"github.com/veandco/go-sdl2/sdl"
)

type keyboardQuitter struct {
	container *common.Element
}

func NewKeyboardQuitter(container *common.Element) *keyboardMover {
	return &keyboardMover{
		container: container,
	}
}

func (mover *keyboardQuitter) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *keyboardQuitter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_ESCAPE] == 1 {
		os.Exit(0)
	}

	return nil
}

func (mover *keyboardQuitter) OnCollision(_ *common.Element) error {
	return nil
}

