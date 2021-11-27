package components

import (
	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *common.Element
	speed     float64
}

func NewKeyboardMover(container *common.Element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
	}
}

func (mover *keyboardMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// fmt.Println("left")
		mover.container.Position.X -= mover.speed * common.Delta
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// fmt.Println("right")
		mover.container.Position.X += mover.speed * common.Delta
	} else if keys[sdl.SCANCODE_UP] == 1 {
		// fmt.Println("up")
		mover.container.Position.Y -= mover.speed * common.Delta
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		// fmt.Println("down")
		mover.container.Position.Y += mover.speed * common.Delta
	}

	return nil
}

func (mover *keyboardMover) OnCollision(_ *common.Element) error {
	return nil
}
