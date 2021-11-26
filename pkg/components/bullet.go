package components

import (
	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 10
)

func newBullet(renderer *sdl.Renderer) *common.Element {
	bullet := &common.Element{}

	sr := NewSpriteRenderer(bullet, renderer, "assets/sprites/player_bullet.bmp")
	bullet.AddComponent(sr)

	mover := newBulletMover(bullet)
	bullet.AddComponent(mover)

	col := common.Circle{
		Center: bullet.Position,
		Radius: 8,
	}
	bullet.Collisions = append(bullet.Collisions, col)

	bullet.Tag = "bullet"

	return bullet
}

var bulletPool []*common.Element

func InitBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, bul)
		common.Elements = append(common.Elements, bul)
	}
}

func bulletFromPool() (*common.Element, bool) {
	for _, bul := range bulletPool {
		if !bul.Active {
			return bul, true
		}
	}

	return nil, false
}
