package components

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type animator struct {
	container       *common.Element
	sequences       map[string]*Sequence
	current         string
	lastFrameChange time.Time
	finished        bool
}

func NewAnimator(
	container *common.Element,
	sequences map[string]*Sequence,
	defaultSequence string) *animator {
	var an animator

	an.container = container
	an.sequences = sequences
	an.current = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) OnUpdate() error {
	s := an.sequences[an.current]

	frameInterval := float64(time.Second) / s.sampleRate

	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.finished = s.nextFrame()
		an.lastFrameChange = time.Now()
	}

	return nil
}

func (an *animator) OnDraw(renderer *sdl.Renderer) error {
	tex := an.sequences[an.current].texture()

	return common.DrawTexture(
		tex,
		an.container.Position,
		an.container.Rotation,
		renderer,
		true)
}

func (an *animator) OnCollision(_ *common.Element) error {
	return nil
}

func (an *animator) setSequence(name string) {
	an.current = name
	an.lastFrameChange = time.Now()
}

type Sequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64
	loop       bool
}

func NewSequence(
	filepath string,
	sampleRate float64,
	loop bool,
	renderer *sdl.Renderer) (*Sequence, error) {

	var seq Sequence

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v failed: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())

		tex, err := common.LoadTextureFromBMP(filename, renderer)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame failed: %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}

	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *Sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

func (seq *Sequence) nextFrame() bool {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}

	return false
}
