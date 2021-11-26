package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.position

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if int(pos.x)-(mover.sr.width/2.0) > 0 {
			mover.container.position.x -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if int(pos.x)+(mover.sr.width/2.0) < screenWidth {
			mover.container.position.x += mover.speed * delta
		}
	}

	return nil
}

func (mover *keyboardMover) onCollision(_ *element) error {
	return nil
}

type keyboardShooter struct {
	container *element
	coolDown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, coolDown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		coolDown:  coolDown}
}

func (shooter *keyboardShooter) onDraw(_ *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := shooter.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.coolDown {
			shooter.shoot(pos.x+25, pos.y-20)
			shooter.shoot(pos.x-25, pos.y-20)

			shooter.lastShot = time.Now()
		}
	}

	return nil
}

func (shooter *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = 270 * (math.Pi / 180)
	}
}

func (shooter *keyboardShooter) onCollision(_ *element) error {
	return nil
}
