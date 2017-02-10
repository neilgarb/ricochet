package ricochet

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadBoardBlank(t *testing.T) {
	s := ``
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardNoBoard(t *testing.T) {
	s := `OOB 10,2`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardDoubleBoard(t *testing.T) {
	s := `BOARD 20
BOARD 10`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardSimple(t *testing.T) {
	s := `BOARD 20`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestReadBoardOOB(t *testing.T) {
	s := `BOARD 10
OOB`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
OOB apple`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
OOB -1,-1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
OOB 1,1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
OOB 1,1
OOB 1,1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardWall(t *testing.T) {
	s := `BOARD 10
WALL`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
WALL apple`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
WALL -1,-1 0`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
WALL 1,1 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
WALL 1,1 1
WALL 1,1 2`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
WALL 1,1 1
WALL 1,1 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardSink(t *testing.T) {
	s := `BOARD 10
SINK`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
SINK apple`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
SINK -1,-1 0 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
SINK 1,1 0 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
SINK 1,1 1 1
SINK 1,2 2 2`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
SINK 1,1 1 1
SINK 1,1 2 2`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardRobot(t *testing.T) {
	s := `BOARD 10
ROBOT`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
ROBOT apple`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
ROBOT -1,-1 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}

	s = `BOARD 10
ROBOT 1,1 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
ROBOT 1,1 1
ROBOT 1,2 2`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}

	s = `BOARD 10
SINK 1,1 1
SINK 1,2 1`
	_, err = ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestReadBoardValid(t *testing.T) {
	s := `BOARD 32
OOB 7,7
OOB 7,8
OOB 8,7
OOB 8,8
WALL 3,0 1
WALL 9,0 1
WALL 5,1 1
WALL 6,1 2
WALL 12,1 0
WALL 12,1 1
WALL 9,2 1
WALL 9,2 2
WALL 2,3 0
WALL 2,3 1
WALL 8,3 1
WALL 4,4 1
WALL 5,4 0
WALL 2,5 1
WALL 2,5 2
WALL 7,5 1
WALL 7,5 2
WALL 10,6 1
WALL 11,6 2
WALL 0,9 2
WALL 11,9 1
WALL 12,8 2
WALL 15,9 2
WALL 3,10 1
WALL 3,10 2
WALL 5,10 2
WALL 10,10 1
WALL 10,10 2
WALL 5,11 0
WALL 5,11 1
WALL 14,11 2
WALL 2,12 2
WALL 2,12 3
WALL 4,12 2
WALL 14,12 1
WALL 3,13 1
WALL 10,14 1
WALL 11,14 2
WALL 3,15 1
WALL 13,15 1
SINK 6,1 0 0
SINK 12,1 3 0
SINK 9,2 2 1
SINK 1,3 1 1
SINK 9,3 0 2
SINK 5,4 2 2
SINK 2,5 3 3
SINK 11,6 1 3
SINK 12,9 0 3
SINK 3,10 0 1
SINK 10,10 1 2
SINK 5,11 2 3
SINK 2,12 1 0
SINK 14,12 3 1
SINK 4,13 3 2
SINK 11,14 2 0
ROBOT 3,2 0
ROBOT 10,5 1
ROBOT 1,14 2
ROBOT 13,11 3`
	_, err := ReadBoard(bufio.NewReader(strings.NewReader(s)))
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}
}
