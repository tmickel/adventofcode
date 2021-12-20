package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	input, _ := fileutil.ScanStrings("input.txt")

	scanners := make([]*Scanner, 0)
	for _, line := range input {
		if strings.HasPrefix(line, "--- scanner") {
			s := Scanner{
				readings: make([]*Point, 0),
			}
			scanners = append(scanners, &s)
			continue
		}
		if line == "" {
			continue
		}
		pts := strings.Split(line, ",")
		x, _ := strconv.Atoi(pts[0])
		y, _ := strconv.Atoi(pts[1])
		z, _ := strconv.Atoi(pts[2])
		pt := &Point{x: x, y: y, z: z}
		scanners[len(scanners)-1].readings = append(scanners[len(scanners)-1].readings, pt)
	}

	scanners[0].happy = true
	// jamScanners(scanners[0], scanners[1])
	// jamScanners(scanners[1], scanners[4])
	// jamScanners(scanners[4], scanners[2])
	// jamScanners(scanners[1], scanners[3])
	for {
		for _, relative := range scanners {
			if relative.happy {
				continue
			}
			for _, origin := range scanners {
				if !origin.happy {
					continue
				}
				jamScanners(origin, relative)
				if relative.happy {
					break
				}
			}
		}

		allHappy := true
		for _, s := range scanners {
			allHappy = allHappy && s.happy
		}
		if allHappy {
			break
		}
	}

	pointSet := make(map[string]bool, 0)
	for i, s := range scanners {
		log.Printf("%d: %d,%d,%d", i, s.absoluteX, s.absoluteY, s.absoluteZ)
		for _, p := range s.readings {
			key := fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
			pointSet[key] = true
		}
	}
	fmt.Println(len(pointSet))

	maxD := 0
	for _, i := range scanners {
		for _, j := range scanners {
			d := abs(i.absoluteX-j.absoluteX) + abs(i.absoluteY-j.absoluteY) + abs(i.absoluteZ-j.absoluteZ)
			if d > maxD {
				maxD = d
			}
		}
	}
	log.Println(maxD)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x int
	y int
	z int
}

func (p *Point) transform(orientation int) *Point {
	x := p.x
	y := p.y
	z := p.z
	// +x
	if orientation == 0 {
		return &Point{x: x, y: y, z: z}
	}
	if orientation == 1 {
		return &Point{x: x, y: z, z: -y}
	}
	if orientation == 2 {
		return &Point{x: x, y: -y, z: -z}
	}
	if orientation == 3 {
		return &Point{x: x, y: -z, z: y}
	}

	// -x
	if orientation == 4 {
		return &Point{x: -x, y: y, z: -z}
	}
	if orientation == 5 {
		return &Point{x: -x, y: z, z: y}
	}
	if orientation == 6 {
		return &Point{x: -x, y: -y, z: z}
	}
	if orientation == 7 {
		return &Point{x: -x, y: -z, z: -y}
	}

	// +y
	if orientation == 8 {
		return &Point{x: y, y: x, z: -z}
	}
	if orientation == 9 {
		return &Point{x: y, y: -x, z: z}
	}
	if orientation == 10 {
		return &Point{x: y, y: z, z: x}
	}
	if orientation == 11 {
		return &Point{x: y, y: -z, z: -x}
	}

	// -y
	if orientation == 12 {
		return &Point{x: -y, y: x, z: z}
	}
	if orientation == 13 {
		return &Point{x: -y, y: -x, z: -z}
	}
	if orientation == 14 {
		return &Point{x: -y, y: z, z: -x}
	}
	if orientation == 15 {
		return &Point{x: -y, y: -z, z: x}
	}

	// +z
	if orientation == 16 {
		return &Point{x: z, y: x, z: y}
	}
	if orientation == 17 {
		return &Point{x: z, y: -x, z: -y}
	}
	if orientation == 18 {
		return &Point{x: z, y: y, z: -x}
	}
	if orientation == 19 {
		return &Point{x: z, y: -y, z: x}
	}

	// -z
	if orientation == 20 {
		return &Point{x: -z, y: x, z: -y}
	}
	if orientation == 21 {
		return &Point{x: -z, y: -x, z: y}
	}
	if orientation == 22 {
		return &Point{x: -z, y: y, z: x}
	}
	if orientation == 23 {
		return &Point{x: -z, y: -y, z: -x}
	}
	return nil
}

type Scanner struct {
	readings    []*Point
	absoluteX   int
	absoluteY   int
	absoluteZ   int
	happy       bool
	orientation int
}

func (s *Scanner) allowedXDeltas(orientation int, other *Scanner) []int {
	allowed := make(map[int]int, 0)
	for _, p := range s.readings {
		t := p.transform(orientation)
		for _, o := range other.readings {
			if _, ok := allowed[o.x-t.x]; !ok {
				allowed[o.x-t.x] = 1
			}
			allowed[o.x-t.x] += 1
		}
	}
	result := make([]int, 0)
	for k, v := range allowed {
		if v >= 12 {
			result = append(result, k)
		}
	}
	return result
}

func (s *Scanner) allowedYDeltas(orientation int, other *Scanner) []int {
	allowed := make(map[int]int, 0)
	for _, p := range s.readings {
		t := p.transform(orientation)
		for _, o := range other.readings {
			if _, ok := allowed[-t.y+o.y]; !ok {
				allowed[-t.y+o.y] = 1
			}
			allowed[-t.y+o.y] += 1
		}
	}
	result := make([]int, 0)
	for k, v := range allowed {
		if v >= 12 {
			result = append(result, k)
		}
	}
	return result
}

func (s *Scanner) allowedZDeltas(orientation int, other *Scanner) []int {
	allowed := make(map[int]int, 0)
	for _, p := range s.readings {
		t := p.transform(orientation)
		for _, o := range other.readings {
			if _, ok := allowed[-t.z+o.z]; !ok {
				allowed[-t.z+o.z] = 1
			}
			allowed[-t.z+o.z] += 1
		}
	}
	result := make([]int, 0)
	for k, v := range allowed {
		if v >= 12 {
			result = append(result, k)
		}
	}
	return result
}

func jamScanners(originScanner, relativeScanner *Scanner) {
	for o := 0; o < 24; o++ {
		allowedX := relativeScanner.allowedXDeltas(o, originScanner)
		allowedY := relativeScanner.allowedYDeltas(o, originScanner)
		allowedZ := relativeScanner.allowedZDeltas(o, originScanner)
		for _, i := range allowedX {
			for _, j := range allowedY {
				for _, k := range allowedZ {
					matches := 0
					for _, p := range relativeScanner.readings {
						pt := p.transform(o)
						for _, q := range originScanner.readings {
							if q.x == pt.x+i && q.y == pt.y+j && q.z == pt.z+k {
								matches++
								break
							}
						}
					}
					if matches >= 12 {
						relativeScanner.absoluteX = i
						relativeScanner.absoluteY = j
						relativeScanner.absoluteZ = k

						newPoints := make([]*Point, 0)
						for _, p := range relativeScanner.readings {
							z := p.transform(o)
							z.x += i
							z.y += j
							z.z += k
							z.transform(originScanner.orientation)
							newPoints = append(newPoints, z)
						}
						relativeScanner.readings = newPoints

						relativeScanner.happy = true
						return
					}
				}
			}
		}
	}
}
