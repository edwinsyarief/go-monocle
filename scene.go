package monocle

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Begin(*Engine)
	End()

	BeforeUpdate()
	Update()
	AfterUpdate()

	BeforeRender(*ebiten.Image)
	Render(*ebiten.Image)
	AfterRender(*ebiten.Image)
}

type BaseScene struct {
	Paused     bool
	EndOfFrame func()

	engine  *Engine
	focused bool

	entityList EntityList
}

func (bs *BaseScene) NewEntityList() EntityList {
	result := &EntityList{
		Scene:    bs,
		Entities: make([]Entity, 0),
		toAdd:    make([]Entity, 0),
		toAwake:  make([]Entity, 0),
		toRemove: make([]Entity, 0),
	}

	return *result
}

func (bs *BaseScene) Begin(engine *Engine) {
	bs.engine = engine
	bs.focused = true

	for i := 0; i < len(bs.entityList.Entities); i++ {
		bs.entityList.Entities[i].SceneBegin(bs)
	}
}

func (bs *BaseScene) End() {
	bs.focused = false

	for i := 0; i < len(bs.entityList.Entities); i++ {
		bs.entityList.Entities[i].SceneEnd(bs)
	}
}

func (bs *BaseScene) BeforeUpdate() {

}

func (bs *BaseScene) Update() {
	if !bs.Paused {
		bs.entityList.Update()
	}
}

func (bs *BaseScene) AfterUpdate() {

}

func (bs *BaseScene) BeforeRender(screen *ebiten.Image) {

}

func (bs *BaseScene) Render(screen *ebiten.Image) {

}

func (bs *BaseScene) AfterRender(screen *ebiten.Image) {

}
