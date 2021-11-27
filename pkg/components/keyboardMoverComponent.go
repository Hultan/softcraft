package components

import (
	"fmt"

	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container   *common.Element
	speed       float64
	movingSpeed float64
	world       *World
}

func NewKeyboardMover(container *common.Element, speed float64, world *World) *keyboardMover {
	return &keyboardMover{
		container:   container,
		speed:       speed,
		world:       world,
	}
}

func (mover *keyboardMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if mover.movingSpeed == 0 {
		mover.movingSpeed = mover.speed * common.Delta
	}

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if mover.canMove(-1, 0) {
			mover.world.Position.X -= mover.movingSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if mover.canMove(1, 0) {
			mover.world.Position.X += mover.movingSpeed
		}
	} else if keys[sdl.SCANCODE_UP] == 1 {
		if mover.canMove(0, -1) {
			mover.world.Position.Y -= mover.movingSpeed
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		if mover.canMove(0, 1) {
			mover.world.Position.Y += mover.movingSpeed
		}
	}

	return nil
}

func (mover *keyboardMover) OnCollision(_ *common.Element) error {
	return nil
}

func (mover *keyboardMover) canMove(dx, dy float64) bool {
	x := int64((mover.world.Position.X + dx*mover.movingSpeed) / common.BlockWidth)
	y := int64((mover.world.Position.Y + dy*mover.movingSpeed) / common.BlockHeight)
	fmt.Println(x,y)
	if mover.world.data[x][y] == 3 {
		return false
	}
	return true
}
