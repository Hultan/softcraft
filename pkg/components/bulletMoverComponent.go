package components

import (
	"math"

	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *common.Element
	speed     float64
}

func newBulletMover(container *common.Element) *bulletMover {
	return &bulletMover{container: container}
}

func (mover *bulletMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) OnUpdate() error {
	c := mover.container

	c.Position.X += common.BulletSpeed * math.Cos(c.Rotation) * common.Delta
	c.Position.Y += common.BulletSpeed * math.Sin(c.Rotation) * common.Delta

	if c.Position.X > common.ScreenWidth || c.Position.X < 0 ||
		c.Position.Y > common.ScreenHeight || c.Position.Y < 0 {
		c.Active = false
	}

	c.Collisions[0].Center = c.Position

	return nil
}

func (mover *bulletMover) OnCollision(_ *common.Element) error {
	mover.container.Active = false
	return nil
}
