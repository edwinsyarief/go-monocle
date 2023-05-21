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
	PreviousTime time.Time
	Delta        float64

	Scene     Scene
	NextScene Scene

	ClearColor color.Color
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
