package components

import "github.com/veandco/go-sdl2/sdl"

type vulnerableToBullets struct {
	container *Element
	animator  *animator
}

func NewVulnerableToBullets(container *Element) *vulnerableToBullets {
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

func (vtb *vulnerableToBullets) OnCollision(other *Element) error {
	if other.Tag == "bullet" {
		vtb.animator.setSequence("destroy")
	}
	return nil
}
