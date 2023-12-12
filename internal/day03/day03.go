package day03

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func Star1(input string) error {
	grid := parseInput(input)
	iter := grid.Iterater()

	sum := 0
	for n, ok := iter.NextNumber(); ok; n, ok = iter.NextNumber() {
		for _, surround := range surroundingCoords(n.coords) {
			b := grid.Get(surround.x, surround.y)
			if b != byte('.') && !isDigit(b) {
				sum += n.value
				break
			}
		}
	}

	fmt.Println(sum)
	return nil
}

func Star2(input string) error {
	grid := parseInput(input)
	iter := grid.Iterater()

	gears := map[Coords][]int{}
	for n, ok := iter.NextNumber(); ok; n, ok = iter.NextNumber() {
		for _, surround := range surroundingCoords(n.coords) {
			b := grid.Get(surround.x, surround.y)
			if b == byte('*') {
				// store all numbers we find next to each gear
				pos := Coords{surround.x, surround.y}
				ns := gears[pos]
				ns = append(ns, n.value)
				gears[pos] = ns
			}
		}
	}

	sum := 0
	for _, gearNs := range gears {
		if len(gearNs) == 2 {
			sum += gearNs[0] * gearNs[1]
		}
	}

	fmt.Println(sum)
	return nil
}

type Coords struct {
	x int
	y int
}

type Grid struct {
	items map[Coords]byte
	maxX  int
	maxY  int
}

func (grid *Grid) Iterater() GridIterator {
	return GridIterator{grid, 0, 0}
}

func (grid *Grid) Get(x int, y int) byte {
	b, found := grid.items[Coords{x, y}]
	if !found {
		b = byte('.')
	}
	return b
}

type GridIterator struct {
	grid *Grid
	x    int
	y    int
}

// Being used to iterators in Rust, I wanted to see what
// a similar pattern would look like in Go..
func (iter *GridIterator) Next() (GridIteratorItem, bool) {
	x := iter.x
	y := iter.y
	item := iter.grid.Get(x, y)

	iter.x += 1
	if iter.x > iter.grid.maxX {
		iter.x = 0
		iter.y += 1
	}
	if iter.y > iter.grid.maxY {
		return GridIteratorItem{}, false
	}

	return GridIteratorItem{x, y, item}, true
}

type GridIteratorItem struct {
	x    int
	y    int
	item byte
}

func (iter *GridIterator) NextNumber() (GridIteratorNumber, bool) {
	digits := []byte{}
	coords := []Coords{}
	for {
		item, found := iter.Next()

		itemIsDigit := found && isDigit(item.item)

		if itemIsDigit {
			digits = append(digits, item.item)
			coords = append(coords, Coords{item.x, item.y})
		}

		if !found || !itemIsDigit || item.x == iter.grid.maxX {
			if len(digits) > 0 {
				// We've collected some digits so return them.
				n, err := strconv.Atoi(string(digits))
				if err != nil {
					panic("Cannot parse digits to number")
				}
				return GridIteratorNumber{n, coords}, true
			} else if !found {
				// We've finished iterating.
				return GridIteratorNumber{}, false
			} else {
				// No digits seen yet so keep looking.
				continue
			}
		}
	}
}

type GridIteratorNumber struct {
	value  int
	coords []Coords
}

func surroundingCoords(coords []Coords) []Coords {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	for _, c := range coords {
		if c.x < minX {
			minX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	minX -= 1
	minY -= 1
	maxX += 1
	maxY += 1

	cs := []Coords{}
	for x := minX; x <= maxX; x++ {
		cs = append(cs, Coords{x, minY})
		cs = append(cs, Coords{x, maxY})
	}
	for y := minY + 1; y < maxY; y++ {
		cs = append(cs, Coords{minX, y})
		cs = append(cs, Coords{maxX, y})
	}
	return cs
}

func parseInput(input string) Grid {
	lines := bytes.Split(bytes.TrimSpace([]byte(input)), []byte("\n"))
	grid := Grid{
		items: map[Coords]byte{},
		maxX:  0,
		maxY:  0,
	}

	for y, line := range lines {
		for x, b := range line {
			if b != byte('.') {
				grid.items[Coords{x, y}] = b
			}
			if x > grid.maxX {
				grid.maxX = x
			}
			if y > grid.maxY {
				grid.maxY = y
			}
		}
	}

	return grid
}

func isDigit(b byte) bool {
	return b >= byte('0') && b <= byte('9')
}
