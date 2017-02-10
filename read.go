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
func ReadBoard(r *bufio.Reader) (*Board, error) {
	var b *Board
	line := 0
	for {
		line++

		s, err := r.ReadString('\n')
		if err == io.EOF {
		} else if err != nil {
			return nil, err
		}

		s = strings.TrimSpace(s)
		if s == "" {
			if b == nil {
				return nil, errors.New("no board")
			}
			return b, nil
		}

		sl := strings.Fields(s)
		if len(sl) == 0 {
			return nil, fmt.Errorf("error line %d: no command", line)
		}

		switch sl[0] {
		case "BOARD":
			if board, err := readBoardBoard(sl[1:], b); err != nil {
				return nil, fmt.Errorf("error line %d: %v", line, err)
			} else {
				b = board
			}
		case "OOB":
			if err := readBoardOOB(sl[1:], b); err != nil {
				return nil, fmt.Errorf("error line %d: %v", line, err)
			}
		case "WALL":
			if err := readBoardWall(sl[1:], b); err != nil {
				return nil, fmt.Errorf("error line %d: %v", line, err)
			}
		case "SINK":
			if err := readBoardSink(sl[1:], b); err != nil {
				return nil, fmt.Errorf("error line %d: %v", line, err)
			}
		case "ROBOT":
			if err := readBoardRobot(sl[1:], b); err != nil {
				return nil, fmt.Errorf("error line %d: %v", line, err)
			}
		default:
			return nil, fmt.Errorf(
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
	return b.AddWall(Wall{pos, dir})
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

func readBoardRobot(tl []string, b *Board) error {
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
	c, err := strconv.Atoi(tl[1])
	if err != nil {
		return err
	}
	return b.AddRobot(Robot{Colour(c)}, pos)
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
