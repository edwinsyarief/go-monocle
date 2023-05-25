package monocle

import (
	"encoding/hex"
	"image/color"
)

func indexOfEntity(entity *Entity, entityList []Entity) int {
	for i, v := range entityList {
		if entity == &v {
			return i
		}
	}
	return -1
}

func indexOfComponent(component *Component, componentList []Component) int {
	for i, v := range componentList {
		if component == &v {
			return i
		}
	}
	return -1
}

func indexOfRenderer(renderer Renderer, rendererList []Renderer) int {
	for i, v := range rendererList {
		if renderer == v {
			return i
		}
	}
	return -1
}

func removeFromEntityArray(entity *Entity, entityList []Entity) []Entity {
	var index = indexOfEntity(entity, entityList)

	entityList = append(entityList[:index], entityList[index+1:]...)

	return entityList
}

func removeFromComponentArray(component *Component, componentList []Component) []Component {
	var index = indexOfComponent(component, componentList)

	componentList = append(componentList[:index], componentList[index+1:]...)

	return componentList
}

func removeFromRendererArray(renderer Renderer, rendererList []Renderer) []Renderer {
	var index = indexOfRenderer(renderer, rendererList)

	rendererList = append(rendererList[:index], rendererList[index+1:]...)

	return rendererList
}

func HexToColor(str string) color.RGBA {
	result, _ := hex.DecodeString(str)
	return color.RGBA{result[0], result[1], result[2], 0xff}
}
