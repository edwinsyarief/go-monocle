package monocle

import (
	"image/color"
	"log"
	"time"

	_ "github.com/silbinarywolf/preferdiscretegpu"

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
	ExitOnEscapeKeypress bool
	IsFixedUpdate        bool

	WindowWidth       int
	WindowHeight      int
	ScreenWidth       int
	ScreenHeight      int
	FullScreen        bool
	AllowWindowResize bool
}

type Engine struct {
	Context    *Context
	ClearColor color.Color
	Delta      float64

	previousTime time.Time
	scene        Scene
	nextScene    Scene
}

func (g *Engine) Update() error {

	var delta float64
	if g.Context.IsFixedUpdate {
		delta = 1.0 / 60.0
	} else {
		now := time.Now()
		delta = now.Sub(g.previousTime).Seconds()
		g.previousTime = now
	}
	g.Delta = delta

	g.scene.BeforeUpdate()
	g.scene.Update()
	g.scene.AfterUpdate()

	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	screen.Clear()

	if g.scene != nil {
		g.scene.BeforeRender(screen)
	}

	screen.Fill(g.ClearColor)

	if g.scene != nil {
		g.scene.Render(screen)
		g.scene.AfterRender(screen)
	}
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Context.ScreenWidth, g.Context.ScreenHeight
}

func NewEngine(context *Context, scene Scene) *Engine {
	ebiten.SetWindowTitle(context.Title)

	switch context.Platform {
	case PS4:
	case PS5:
	case XBOX:
		ebiten.SetWindowSize(1920, 1080)
	case NSWITCH:
		ebiten.SetWindowSize(1280, 720)
	default:
		ebiten.SetWindowSize(context.WindowWidth, context.WindowHeight)
	}

	if context.AllowWindowResize {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	}

	switch context.Platform {
	case XBOX:
	case PS4:
	case PS5:
	case NSWITCH:
		ebiten.SetFullscreen(true)
	default:
		ebiten.SetFullscreen(context.FullScreen)
	}

	engine := &Engine{Context: context, previousTime: time.Now()}

	if scene != nil {
		engine.nextScene = scene
		scene.Begin(engine)
		engine.scene = scene
	}

	return engine
}

func (g *Engine) Run() error {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
