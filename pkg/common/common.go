package common

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
}

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Tag        string
	Collisions []Circle
	Components []Component
}

var Elements []*Element

type Circle struct {
	Center Vector
	Radius float64
}

type Vector struct {
	X float64
	Y float64
}

var Delta float64

const (
	ScreenWidth  = 600
	ScreenHeight = 800

	TargetTicksPerSecond = 60
)

const BasicEnemySize = 105


const (
	PlayerSize = 105

	PlayerShotCoolDown = time.Millisecond * 250
)
