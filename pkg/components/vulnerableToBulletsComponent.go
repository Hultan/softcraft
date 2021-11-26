package components

import (
	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type vulnerableToBullets struct {
	container *common.Element
	animator  *animator
}

func NewVulnerableToBullets(container *common.Element) *vulnerableToBullets {
	return &vulnerableToBullets{
		container: container,
		animator:  container.GetComponent(&animator{}).(*animator)}
}

func (vtb *vulnerableToBullets) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (vtb *vulnerableToBullets) OnUpdate() error {
	if vtb.animator.finished && vtb.animator.current == "destroy" {
		vtb.container.Active = false
	}

	return nil
}

func (vtb *vulnerableToBullets) OnCollision(other *common.Element) error {
	if other.Tag == "bullet" {
		vtb.animator.setSequence("destroy")
	}
	return nil
}
