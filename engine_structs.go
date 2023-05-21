package monocle

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlatformType int

const (
	PC PlatformType = iota
	XBOX
	PS4
	PS5
	NSWITCH
)

type Context struct {
	Title                string
	Version              string
	Platform             PlatformType
	IsDebug              bool
	WindowWidth          int
	WindowHeight         int
	ScreenWidth          int
	ScreenHeight         int
	ExitOnEscapeKeypress bool

	IsFixedUpdate bool
}

type Engine struct {
	Context      *Context
	ClearColor color.Color
	Delta        float64

	previousTime time.Time
	scene     Scene
	nextScene Scene
}

type Renderer interface {
}

type Component interface {
}

type Entity interface {
}

type Scene interface {
	Begin(engine *Engine)
	End()

	BeforeUpdate()
	Update()
	AfterUpdate()

	BeforeRender(screen *ebiten.Image)
	Render(screen *ebiten.Image)
	AfterRender(screen *ebiten.Image)
}
