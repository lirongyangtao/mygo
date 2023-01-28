package example

import (
	"awesomeProject5/base/path_finding"
	"strings"
)

func gridFromString(str string) {
	rows := strings.Split(str, "\n")
	height := len(rows)
	width := len(rows[0])
	grid := path_finding.NewGrid(width, height, path_finding.WithDontCrossCorners(true))
	startX, startY := 0, 0
	endX, endY := 0, 0
	for i, row := range rows {
		for j, v := range row {
			if v == '1' {
				grid.SetWalkableAt(j, i, false)
			} else if v == 'e' {
				endX, endY = j, i
			} else if v == 's' {
				startX, startY = j, i
			}
		}
	}
	_ = startX
	_ = startY
	_ = endX
	_ = endY
	grid.PrintMap()
	grid.PathFindingPrint(startX, startY, endX, endY)
}

// a*算法
func AStar() {
	gridFromString(map1)
}
