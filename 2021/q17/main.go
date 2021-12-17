package main

import "fmt"

func main() {
	all_max_y := 0
	matches := 0

	for ivx := -5000; ivx <= 5000; ivx++ {
		for ivy := -5000; ivy <= 5000; ivy++ {
			x := 0
			y := 0
			vx := ivx
			vy := ivy
			i := 0
			matched := false
			local_max_y := 0

			for i < 500 {
				x += vx
				y += vy
				if y > local_max_y {
					local_max_y = y
				}
				if vx > 0 {
					vx -= 1
				} else if vx < 0 {
					vx += 1
				}
				vy -= 1
				if x >= 34 && x <= 67 && y >= -215 && y <= -186 {
					// if x >= 20 && x <= 30 && y >= -10 && y <= -5 {
					matches += 1
					matched = true
					break
				}
				i++
			}
			if matched && local_max_y > all_max_y {
				all_max_y = local_max_y
			}
		}
	}
	fmt.Println(all_max_y) // 23005
	fmt.Println(matches)   // 2040
}
