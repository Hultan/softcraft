package common

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

// Draw draws all the components of the element
func (elem *Element) Draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.Components {
		err := comp.OnDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update updates all the components of the element
func (elem *Element) Update() error {
	for _, comp := range elem.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

// Collision checks all the components of the element for collisions
func (elem *Element) Collision(other *Element) error {
	for _, comp := range elem.Components {
		err := comp.OnCollision(other)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddComponent adds a component to the element
func (elem *Element) AddComponent(new Component) {
	for _, existing := range elem.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v failed",
				reflect.TypeOf(new)))
		}
	}
	elem.Components = append(elem.Components, new)
}

// GetComponent gets a component of the specific type from the element
func (elem *Element) GetComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.Components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v failed", reflect.TypeOf(withType)))
}
