package components

import (
	"os"

	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardQuitter struct {
	container *common.Element
}

func NewKeyboardQuitter(container *common.Element) *keyboardQuitter {
	return &keyboardQuitter{
		container: container,
	}
}

func (q *keyboardQuitter) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (q *keyboardQuitter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_ESCAPE] == 1 || keys[sdl.SCANCODE_Q] == 1 {
		os.Exit(0)
	}

	return nil
}

func (q *keyboardQuitter) OnCollision(_ *common.Element) error {
	return nil
}
