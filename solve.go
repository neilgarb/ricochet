package ricochet

import (
	"fmt"
	"sort"
	"strconv"
)

func (s *State) Solve(tok Token) []Move {
	sinkPos := s.board.sinks[tok]
	queue := []*State{s}
	tried := map[string]bool{s.String(): true}

	depth := 0
	for len(queue) > 0 {
		qs := queue[0]
		queue = queue[1:]

		if len(qs.path) > 1 {
			if _, ok := s.robots[sinkPos]; ok {
				return qs.path
			}
		}

		if len(qs.path) > depth {
			depth = len(qs.path)
			fmt.Printf("Trying depth %d (queue %d)\n", depth, len(queue))
		}

		for p, r := range s.robots {
			for _, d := range allDirections {
				if s.CanMove(p, d) {
					newState := qs.Clone()
					next := qs.Move(p, d)

					delete(newState.robots, p)
					newState.robots[next] = r
					newState.path = append(qs.path, Move{r, next})

					// If we've already tried this state, ignore.
					hash := newState.String()
					if tried[hash] {
						continue
					}

					tried[hash] = true
					queue = append(queue, newState)
				}
			}
		}
	}

	return nil
}

func (s *State) String() string {
	var sl []int
	for p := range s.robots {
		sl = append(sl, p.X*s.board.size+p.Y)
	}
	sort.Ints(sl)
	var str string
	for _, i := range sl {
		str = str + strconv.Itoa(i) + ","
	}
	return str
}
