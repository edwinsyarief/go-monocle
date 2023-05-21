package monocle

import (
	"log"
	"time"

	_ "github.com/silbinarywolf/preferdiscretegpu"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Engine) Update() error {

	var delta float64
	if g.Context.IsFixedUpdate {
		delta = 1.0 / 60.0
	} else {
		now := time.Now()
		delta = now.Sub(g.PreviousTime).Seconds()
		g.PreviousTime = now
	}
	g.Delta = delta

	g.Scene.BeforeUpdate()
	g.Scene.Update()
	g.Scene.AfterUpdate()

	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	screen.Clear()

	if g.Scene != nil {
		g.Scene.BeforeRender(screen)
	}

	screen.Fill(g.ClearColor)

	if g.Scene != nil {
		g.Scene.Render(screen)
		g.Scene.AfterRender(screen)
	}
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	switch g.Context.Platform {
	case PS4:
	case PS5:
	case XBOX:
		return 1920, 1280
	case NSWITCH:
		return 1280, 720
	}

	return g.Context.ScreenWidth, g.Context.ScreenHeight
}

func NewEngine(context *Context, scene Scene) *Engine {
	ebiten.SetWindowSize(context.WindowWidth, context.WindowHeight)
	ebiten.SetWindowTitle(context.Title)

	engine := &Engine{Context: context, PreviousTime: time.Now()}

	if scene != nil {
		engine.NextScene = scene
		scene.Begin(engine)
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
