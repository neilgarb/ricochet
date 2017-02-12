package ricochet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ReadBoard reads a board configuration. The syntax is as follows:
//
// `BOARD <size>`
// `OOB <position>`
// `WALL <position> <direction>`
// `SINK <position> <colour> <shape>`
// `ROBOT <position> <colour>`
//
// `size` is a number from 1 - 100.
// `position` is a 0-indexed coordinated in the form `col,row`, e.g. `4,5`.
// `direction` is a number from 0 - 3, where 0 is north, 1 is east, etc.
// `colour` is a number from 0 - 3 for `SINK`, and any number for `ROBOT`.
// `shape` is a number from 0 - 3.
func ReadBoard(r *bufio.Reader) (*Board, *State, error) {
	var (
		board *Board
		state *State
		line  int
	)
	for {
		line++

		s, err := r.ReadString('\n')
		if err == io.EOF {
		} else if err != nil {
			return nil, nil, err
		}

		s = strings.TrimSpace(s)
		if s == "" {
			if board == nil {
				return nil, nil, errors.New("no board")
			}
			return board, state, nil
		}

		sl := strings.Fields(s)
		if len(sl) == 0 {
			return nil, nil, fmt.Errorf("error line %d: no command", line)
		}

		switch sl[0] {
		case "BOARD":
			if b, err := readBoardBoard(sl[1:], board); err != nil {
				return nil, nil, fmt.Errorf("error line %d: %v", line, err)
			} else {
				board = b
				state = b.NewState()
			}
		case "OOB":
			if err := readBoardOOB(sl[1:], board); err != nil {
				return nil, nil, fmt.Errorf("error line %d: %v", line, err)
			}
		case "WALL":
			if err := readBoardWall(sl[1:], board); err != nil {
				return nil, nil, fmt.Errorf("error line %d: %v", line, err)
			}
		case "SINK":
			if err := readBoardSink(sl[1:], board); err != nil {
				return nil, nil, fmt.Errorf("error line %d: %v", line, err)
			}
		case "ROBOT":
			if err := readBoardRobot(sl[1:], state); err != nil {
				return nil, nil, fmt.Errorf("error line %d: %v", line, err)
			}
		default:
			return nil, nil, fmt.Errorf(
				"error line %d: unknown command %q", line, sl[0])
		}
	}
}

func readBoardBoard(tl []string, b *Board) (*Board, error) {
	if b != nil {
		return nil, errors.New("already have a board")
	}
	if len(tl) != 1 {
		return nil, errors.New("bad syntax")
	}
	size, err := strconv.Atoi(tl[0])
	if err != nil {
		return nil, err
	}
	return NewBoard(size)
}

func readBoardOOB(tl []string, b *Board) error {
	if b == nil {
		return errors.New("don't have a board yet")
	}
	if len(tl) != 1 {
		return errors.New("bad syntax")
	}
	pos, err := readPos(tl[0])
	if err != nil {
		return err
	}
	return b.SetOOB(pos)
}

func readBoardWall(tl []string, b *Board) error {
	if b == nil {
		return errors.New("don't have a board yet")
	}
	if len(tl) != 2 {
		return errors.New("bad syntax")
	}
	pos, err := readPos(tl[0])
	if err != nil {
		return err
	}
	d, err := strconv.Atoi(tl[1])
	if err != nil {
		return err
	}
	dir := Direction(d)
	if !dir.Valid() {
		return errors.New("bad direction")
	}
	return b.AddWall(pos, dir)
}

func readBoardSink(tl []string, b *Board) error {
	if b == nil {
		return errors.New("don't have a board yet")
	}
	if len(tl) != 3 {
		return errors.New("bad syntax")
	}
	pos, err := readPos(tl[0])
	if err != nil {
		return errors.New("bad position")
	}
	c, err := strconv.Atoi(tl[1])
	if err != nil {
		return err
	}
	col := Colour(c)
	if !col.ValidForToken() {
		return errors.New("bad colour")
	}
	s, err := strconv.Atoi(tl[2])
	if err != nil {
		return err
	}
	shape := Shape(s)
	if !shape.Valid() {
		return errors.New("bad shape")
	}
	return b.AddSink(Token{shape, col}, pos)
}

func readBoardRobot(tl []string, s *State) error {
	if s == nil {
		return errors.New("don't have a state yet")
	}
	if len(tl) != 2 {
		return errors.New("bad syntax")
	}
	pos, err := readPos(tl[0])
	if err != nil {
		return err
	}
	c, err := strconv.Atoi(tl[1])
	if err != nil {
		return err
	}
	return s.AddRobot(pos, Robot{Colour(c)})
}

func readPos(pos string) (Position, error) {
	parts := strings.SplitN(pos, ",", 2)
	if len(parts) != 2 {
		return Position{}, errors.New("bad position")
	}
	col, err := strconv.Atoi(parts[0])
	if err != nil {
		return Position{}, err
	}
	row, err := strconv.Atoi(parts[1])
	if err != nil {
		return Position{}, err
	}
	return Position{col, row}, nil
}
