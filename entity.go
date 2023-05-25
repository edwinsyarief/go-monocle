package monocle

import (
	"golang.org/x/image/math/f64"
)

type EntityList struct {
	Scene    Scene
	Entities []Entity

	toAdd    []Entity
	toAwake  []Entity
	toRemove []Entity
}

func (el *EntityList) UpdateLists() {
	if len(el.toAdd) > 0 {
		for _, e := range el.toAdd {
			index := indexOfEntity(&e, el.Entities)
			if index == -1 {
				el.Entities = append(el.Entities, e)

				if el.Scene != nil {
					e.Added(el.Scene)
				}
			}
		}
	}

	if len(el.toRemove) > 0 {
		for _, e := range el.toRemove {
			index := indexOfEntity(&e, el.Entities)
			if index > -1 {
				removeFromEntityArray(&e, el.Entities)

				if el.Scene != nil {
					e.Removed(el.Scene)
				}
			}
		}

		el.toRemove = nil
	}

	if len(el.toAdd) > 0 {
		el.toAwake = append(el.toAwake, el.toAdd...)
		el.toAdd = nil

		for _, e := range el.toAwake {
			if e.GetScene() == el.Scene {
				e.Awake(el.Scene)
			}
		}

		el.toAwake = nil
	}
}

func (el *EntityList) Update() {
	for _, e := range el.Entities {
		if e.IsActive() {
			e.Update()
		}
	}
}

func (el *EntityList) Add(e Entity) {
	index := indexOfEntity(&e, el.Entities)
	if index == -1 {
		el.toAdd = append(el.toAdd, e)
	}
}

func (el *EntityList) Remove(e Entity) {
	index := indexOfEntity(&e, el.Entities)
	if index > -1 {
		el.toRemove = append(el.toRemove, e)
	}
}

func (el *EntityList) AddRange(arrEntities []Entity) {
	for _, e := range el.Entities {
		el.Add(e)
	}
}

func (el *EntityList) RemoveRange(arrEntities []Entity) {
	for _, e := range el.Entities {
		el.Remove(e)
	}
}

func (el *EntityList) Count() int {
	return len(el.Entities)
}

func (el *EntityList) Get(index int) Entity {
	return el.Entities[index]
}

type Entity interface {
	GetScene() Scene
	IsActive() bool

	SceneBegin(Scene)
	SceneEnd(Scene)

	Awake(Scene)
	Added(Scene)
	Removed(Scene)

	Update()
	Render()
}

type BaseEntity struct {
	Scene      Scene
	Position   f64.Vec2
	Active     bool
	Visible    bool
	Collidable bool
}
