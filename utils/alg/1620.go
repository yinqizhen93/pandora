package main

import (
	"fmt"
	"math"
)

func bestCoordinate(towers [][]int, radius int) []int {
	var maxSignal int
	p := make([]int, 2)
	minX, maxX, minY, maxY := narrowRange(towers)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			var total int
			for _, tower := range towers {
				d := calcDistance(tower, x, y)
				var s int
				if d <= float64(radius) {
					s = int(math.Floor(float64(tower[2]) / (1 + d)))
				}
				total += s
			}
			fmt.Println(x, y, total)
			if total > maxSignal {
				maxSignal = total
				p[0] = x
				p[1] = y
			} else {
				if total == maxSignal {
					if x < p[0] {
						p[0] = x
						p[1] = y
					} else {
						if x == p[0] && y < p[1] {
							p[0] = x
							p[1] = y
						}
					}
				}
			}
		}
	}
	fmt.Println("p", p)
	return p
}

// 缩小范围
func narrowRange(towers [][]int) (int, int, int, int) {
	minX := towers[0][0]
	maxX := towers[0][0]
	minY := towers[0][1]
	maxY := towers[0][1]
	for i := 1; i < len(towers); i++ {
		if minX > towers[i][0] {
			minX = towers[i][0]
		}
		if maxX < towers[i][0] {
			maxX = towers[i][0]
		}
		if minY > towers[i][1] {
			minY = towers[i][1]
		}
		if maxY < towers[i][1] {
			maxY = towers[i][1]
		}
	}
	return minX, maxX, minY, maxY
}

// 计算距离
func calcDistance(tower []int, x2, y2 int) float64 {
	x1 := tower[0]
	y1 := tower[1]
	return math.Sqrt(math.Pow(float64(x1)-float64(x2), 2) + math.Pow(float64(y1)-float64(y2), 2))
}

func main() {
	c := [][]int{{42, 0, 0}}
	bestCoordinate(c, 2)
}
