package components

import (
	"softcraft/pkg/types"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *Element
	speed     float64

	sr *spriteRenderer
}

func NewKeyboardMover(container *Element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.GetComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.Position

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if int(pos.X)-(mover.sr.width/2.0) > 0 {
			mover.container.Position.X -= mover.speed * types.Delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if int(pos.X)+(mover.sr.width/2.0) < types.ScreenWidth {
			mover.container.Position.X += mover.speed * types.Delta
		}
	}

	return nil
}

func (mover *keyboardMover) OnCollision(_ *Element) error {
	return nil
}

