package main

import (
	"sort"
	"strings"
)

func transform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
func main() {

	plants := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	println(strings.Join(plants, ","))

	sort.StringSlice(plants).Sort()
	println(strings.Join(plants, ","))

	twoWorlds := transform("New", "Earth", "Mars")
	println(strings.Join(twoWorlds, ","))
}
