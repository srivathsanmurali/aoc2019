package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type loc struct {
	x, y int
}

type line struct {
	a, b loc
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func (l *line) has_point(p loc) bool {
	maxX := max(l.a.x, l.b.x)
	minX := min(l.a.x, l.b.x)
	maxY := max(l.a.y, l.b.y)
	minY := min(l.a.y, l.b.y)
	if p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY {
		return true
	}
	return false
}

func (lineA *line) parallel(lineB line) bool {
	return (lineA.a.y == lineA.b.y && lineB.a.y == lineB.b.y) || (lineA.a.x == lineA.b.x && lineB.a.x == lineB.b.x)
}

func (lineA *line) intersect(lineB line) loc {
	if lineA.parallel(lineB) {
		return loc{}
	}

	var inter loc
	if lineA.a.y == lineA.b.y {
		inter.x = lineB.a.x
		inter.y = lineA.a.y
	}

	if lineA.a.x == lineA.b.x {
		inter.x = lineA.a.x
		inter.y = lineB.a.y
	}

	if lineA.has_point(inter) && lineB.has_point(inter) {
		return inter
	}

	return loc{}
}

func parse_route(reader *bufio.Reader) []line {
	cable, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error while getting line: %v", err)
	}
	f := func(c rune) bool {
		return c == ','
	}

	cur_loc := loc{x: 0, y: 0}
	var path []line
	for _, val := range strings.FieldsFunc(cable, f) {
		var dir string
		var mag int

		_, err := fmt.Sscanf(val, "%1s%d", &dir, &mag)
		if err != nil {
			log.Fatalf("Error while parsing: %v", err)
		}

		var diff loc
		switch dir {
		case "U":
			diff = loc{x: 0, y: mag}
		case "D":
			diff = loc{x: 0, y: -mag}
		case "R":
			diff = loc{x: mag, y: 0}
		case "L":
			diff = loc{x: -mag, y: 0}
		}

		new_loc := loc{x: cur_loc.x + diff.x, y: cur_loc.y + diff.y}
		path = append(path, line{cur_loc, new_loc})
		cur_loc = new_loc
	}

	return path
}

func get_intersections(pathA, pathB []line) []loc {
	var intersections []loc
	for _, lineA := range pathA {
		for _, lineB := range pathB {
			inter := lineA.intersect(lineB)
			if inter.x != 0 && inter.y != 0 {
				intersections = append(intersections, inter)
			}
		}
	}

	return intersections
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pathA := parse_route(reader)
	pathB := parse_route(reader)

	intersections := get_intersections(pathA, pathB)

	minDist := 1000000
	for _, inter := range intersections {
		dist := abs(inter.x) + abs(inter.y)
		if dist < minDist {
			minDist = dist
		}
	}

	fmt.Println("Result:", minDist)
	return
}
