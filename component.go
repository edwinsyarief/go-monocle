package monocle

type ComponentList struct {
	entity     Entity
	components []Component
	toAdd      []Component
	toRemove   []Component
}

type Component interface {
	Added(Entity)
	Removed(Entity)
	EntityAdded(Scene)
	EntityRemoved(Scene)
	SceneEnd(Scene)
	EntityAwake()
}
