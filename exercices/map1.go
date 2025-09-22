package exercices

import (
	"fmt"
	"math/rand"
)

func RunMapEx1() {
	colors := []string{
		"blue", "black", "yellow", "pink",
	}
	members := []string{
		"dad", "mum", "son", "daughter",
	}

	selection := map[string]string{}

	for _, member := range members {
		selectedColor := colors[rand.Intn(len(colors))]
		selection[member] = selectedColor
	}

	for key, _ := range selection {
		fmt.Println(key)
	}

	fmt.Println(selection)
}
