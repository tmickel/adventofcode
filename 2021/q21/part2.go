package main

import (
	"fmt"
)

func main() {
	p1 := 10
	p2 := 8
	cache := make(map[string]int64, 0)
	p1Wins, p2Wins := turn(p1, 0, p2, 0, cache)
	if p1Wins > p2Wins {
		fmt.Println(p1Wins)
	} else {
		fmt.Println(p2Wins)
	}
}

func turn(p1 int, scoreP1 int, p2 int, scoreP2 int, cache map[string]int64) (int64, int64) {
	winsP1 := int64(0)
	winsP2 := int64(0)
	ds := []int{3, 4, 5, 4, 5, 6, 5, 6, 7, 4, 5, 6, 5, 6, 7, 6, 7, 8, 5, 6, 7, 6, 7, 8, 7, 8, 9}
	for _, dP1 := range ds {
		newp1Pos := position(p1, dP1)
		newp1Score := scoreP1 + newp1Pos

		if newp1Score >= 21 {
			winsP1++
			continue
		}

		for _, dP2 := range ds {
			newp2Pos := position(p2, dP2)
			newp2Score := scoreP2 + newp2Pos

			if newp2Score >= 21 {
				winsP2++
				continue
			}

			if v, ok := cache[fmt.Sprintf("%d:%d:%d:%d:1", newp1Pos, newp1Score, newp2Pos, newp2Score)]; ok {
				winsP1 += v
				winsP2 += cache[fmt.Sprintf("%d:%d:%d:%d:2", newp1Pos, newp1Score, newp2Pos, newp2Score)]
			} else {
				a, b := turn(newp1Pos, newp1Score, newp2Pos, newp2Score, cache)
				cache[fmt.Sprintf("%d:%d:%d:%d:1", newp1Pos, newp1Score, newp2Pos, newp2Score)] = a
				cache[fmt.Sprintf("%d:%d:%d:%d:2", newp1Pos, newp1Score, newp2Pos, newp2Score)] = b
				winsP1 += a
				winsP2 += b
			}
		}
	}

	return winsP1, winsP2
}

func position(oldPos int, dieRoll int) int {
	return 1 + (dieRoll+oldPos-1)%(10)
}
