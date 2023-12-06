package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input
var input string

const MAXINT = int(^uint(0) >> 1)

type ComponentKey struct {
	From string
	To   string
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println("Parsing seeds...")
	seeds := parseSeeds(lines[0])
	fmt.Println("Parsing sections...")
	sections := parseSections(lines[2:])
	fmt.Println("Getting path...")
	path := getPath(sections)
	fmt.Println("Path is", path)

	for section, m := range sections {
		fmt.Println(section)
		fmt.Printf("%+v\n", m)
	}

	fmt.Println("Starting computation...")
	lowest := MAXINT
	for _, seed := range seeds {
		fmt.Println("Calculating for seed", seed)
		currentMaterial := "seed"
		currentVal := seed
		for _, destMaterial := range path {
			key := ComponentKey{
				From: currentMaterial,
				To:   destMaterial,
			}
			val := sections[key].Match(currentVal)
			// fmt.Printf("DEBUG: match for %s->%s: %d->%d\n", currentMaterial, destMaterial, currentVal, val)
			currentVal = val
			currentMaterial = destMaterial
		}
		lowest = min(lowest, currentVal)
	}

	fmt.Println("ANSWER", lowest)
}

func getPath(sections map[ComponentKey]Section) []string {
	path := []string{}
	current := "seed"
	for {
		for key := range sections {
			if key.From != current {
				continue
			}
			current = key.To
			path = append(path, key.To)
			if current == "location" {
				return path
			}
		}
	}
}

type Section struct {
	ranges []Range
}

func (s Section) Match(i int) int {
	lastRange := s.ranges[len(s.ranges)-1]
	if lastRange.sourceRangeStart+lastRange.rangeLength < i {
		// fmt.Println("DEBUG: exit early because upper range constraint")
		return i
	}

	for _, r := range s.ranges {
		if r.sourceRangeStart > i {
			// fmt.Println("DEBUG: exit early because lower range constraint", r.sourceRangeStart, i)
			return i
		}

		if res, ok := r.Match(i); ok {
			return res
		}
	}
	return i
}

type Range struct {
	destRangeStart, sourceRangeStart, rangeLength int
}

func (r Range) String() string {
	return fmt.Sprintf("Source(%d-%d) -> Dest(%d-%d)", r.sourceRangeStart, r.sourceRangeStart+r.rangeLength, r.destRangeStart, r.destRangeStart+r.rangeLength)
}

func (r Range) Match(i int) (int, bool) {
	if i >= r.sourceRangeStart && i <= r.sourceRangeStart+r.rangeLength {
		// fmt.Println("DEBUG: MATCH!!", r, i)
		return i - r.sourceRangeStart + r.destRangeStart, true
	} else {
		// fmt.Println("DEBUG: no match", i, r)
	}
	return 0, false
}

func parseSections(lines []string) map[ComponentKey]Section {
	var key ComponentKey
	sections := map[ComponentKey]Section{}
	section := Section{}
	addSection := func() {
		sort.Slice(section.ranges, func(i, j int) bool { return section.ranges[i].sourceRangeStart < section.ranges[j].sourceRangeStart })
		sections[key] = section
	}
	for _, line := range lines {
		if line == "" {
			addSection()
			continue
		}

		if strings.HasSuffix(line, "map:") {
			key = parseComponentKey(line)
			section = Section{ranges: []Range{}}
			continue
		}

		res := parseIntList(line)
		section.ranges = append(section.ranges, Range{
			destRangeStart:   res[0],
			sourceRangeStart: res[1],
			rangeLength:      res[2],
		})
	}
	addSection()
	return sections
}

func parseComponentKey(line string) ComponentKey {
	line = strings.TrimSuffix(line, " map:")
	parts := strings.Split(line, "-to-")
	return ComponentKey{
		From: parts[0],
		To:   parts[1],
	}
}

func parseSeeds(seedLine string) []int {
	parts := strings.Split(seedLine, "seeds:")
	return parseIntList(parts[1])
}

func parseIntList(s string) []int {
	intStrs := strings.Fields(s)
	ints := make([]int, len(intStrs))
	for i, s := range intStrs {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("Error: %s", s)
		}
		ints[i] = int(num)
	}
	return ints
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
