package ricochet

import "testing"

func TestDirectionFlip(t *testing.T) {
	if f := DirectionNorth.Flip(); f != DirectionSouth {
		t.Errorf("expected %d, got %d", DirectionSouth, f)
	}
	if f := DirectionEast.Flip(); f != DirectionWest {
		t.Errorf("expected %d, got %d", DirectionWest, f)
	}
	if f := DirectionSouth.Flip(); f != DirectionNorth {
		t.Errorf("expected %d, got %d", DirectionNorth, f)
	}
	if f := DirectionWest.Flip(); f != DirectionEast {
		t.Errorf("expected %d, got %d", DirectionEast, f)
	}
}

func TestPositionEqual(t *testing.T) {
	if !(Position{0, 0}).Equal(Position{0, 0}) {
		t.Errorf("expected equal")
	}
	if (Position{10, 10}).Equal(Position{10, 11}) {
		t.Errorf("expected not equal")
	}
}

func TestPositionNext(t *testing.T) {
	pos := Position{5, 5}

	exp := Position{5, 4}
	if act := pos.Next(DirectionNorth); !act.Equal(exp) {
		t.Errorf("expected %v, got %v", exp, act)
	}
	exp = Position{6, 5}
	if act := pos.Next(DirectionEast); !act.Equal(exp) {
		t.Errorf("expected %v, got %v", exp, act)
	}
	exp = Position{5, 6}
	if act := pos.Next(DirectionSouth); !act.Equal(exp) {
		t.Errorf("expected %v, got %v", exp, act)
	}
	exp = Position{4, 5}
	if act := pos.Next(DirectionWest); !act.Equal(exp) {
		t.Errorf("expected %v, got %v", exp, act)
	}
}

func TestStateAddRobot(t *testing.T) {
	b, _ := NewBoard(10)
	s := b.NewState()

	r := Robot{ColourRed}
	if err := s.AddRobot(Position{-1, -1}, r); err == nil {
		t.Errorf("expected error")
	}
	if err := s.AddRobot(Position{0, 0}, r); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := s.AddRobot(Position{0, 1}, r); err == nil {
		t.Errorf("expected error")
	}

	r = Robot{ColourBlue}
	if err := s.AddRobot(Position{0, 0}, r); err == nil {
		t.Errorf("expected error")
	}
	if err := s.AddRobot(Position{0, 1}, r); err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestNewBoard(t *testing.T) {
	if _, err := NewBoard(-1); err == nil {
		t.Errorf("expected error")
	}
	if _, err := NewBoard(0); err == nil {
		t.Errorf("expected error")
	}
	if _, err := NewBoard(10); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if _, err := NewBoard(101); err == nil {
		t.Errorf("expected error")
	}
}

func TestBoardOOB(t *testing.T) {
	b, _ := NewBoard(10)

	if err := b.SetOOB(Position{-1, -1}); err == nil {
		t.Errorf("expected error")
	}
	if err := b.SetOOB(Position{0, 0}); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := b.SetOOB(Position{0, 0}); err == nil {
		t.Errorf("expected error")
	}
	if err := b.SetOOB(Position{10, 10}); err == nil {
		t.Errorf("expected error")
	}
	if ib := b.InBounds(Position{-1, -1}); ib {
		t.Errorf("expected oob")
	}
	if ib := b.InBounds(Position{0, 0}); ib {
		t.Errorf("expected oob")
	}
	if ib := b.InBounds(Position{1, 0}); !ib {
		t.Errorf("expected in bounds")
	}
}

func TestAddWall(t *testing.T) {
	b, _ := NewBoard(10)

	if err := b.AddWall(Position{-1, -1}, DirectionNorth); err == nil {
		t.Errorf("expected error")
	}
	if err := b.AddWall(Position{0, 0}, DirectionNorth); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := b.AddWall(Position{0, 0}, DirectionNorth); err == nil {
		t.Errorf("expected error")
	}
}

func TestAddSink(t *testing.T) {
	b, _ := NewBoard(10)

	tok := Token{ShapeCircle, ColourBlue}
	if err := b.AddSink(tok, Position{-1, -1}); err == nil {
		t.Errorf("expected error")
	}
	if err := b.AddSink(tok, Position{0, 0}); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := b.AddSink(tok, Position{0, 0}); err == nil {
		t.Errorf("expected error")
	}

	tok = Token{ShapeCircle, ColourRed}
	if err := b.AddSink(tok, Position{0, 0}); err == nil {
		t.Errorf("expected error")
	}
	if err := b.AddSink(tok, Position{1, 0}); err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestValid(t *testing.T) {
	b, _ := NewBoard(10)

	if b.Valid() {
		t.Errorf("expected invalid")
	}

	tot := len(allShapes) * len(allColours)
	i := 0
	for _, s := range allShapes {
		for _, c := range allColours {
			i++
			b.AddSink(Token{s, c}, Position{i / 10, i % 10})
			if i < tot-1 && b.Valid() {
				t.Errorf("expected invalid")
			}
		}
	}

	if !b.Valid() {
		t.Errorf("expected valid")
	}
}

type canMoveTest struct {
	Position  Position
	Direction Direction
	CanMove   bool
}

func TestBoardCanMove(t *testing.T) {
	b, _ := NewBoard(10)
	b.AddWall(Position{1, 1}, DirectionNorth)
	b.SetOOB(Position{5, 5})

	tests := []canMoveTest{
		{Position{0, 0}, DirectionNorth, false}, // oob
		{Position{0, 0}, DirectionEast, true},
		{Position{0, 0}, DirectionSouth, true},
		{Position{0, 0}, DirectionWest, false}, // oob

		{Position{1, 1}, DirectionNorth, false}, // wall in the way
		{Position{1, 1}, DirectionEast, true},
		{Position{1, 1}, DirectionSouth, true},
		{Position{1, 1}, DirectionWest, true},

		{Position{1, 0}, DirectionNorth, false}, // oob
		{Position{1, 0}, DirectionEast, true},
		{Position{1, 0}, DirectionSouth, false}, // wall in the way
		{Position{1, 0}, DirectionWest, true},

		{Position{9, 9}, DirectionNorth, true},
		{Position{9, 9}, DirectionEast, false},  // oob
		{Position{9, 9}, DirectionSouth, false}, // oob
		{Position{9, 9}, DirectionWest, true},

		{Position{6, 5}, DirectionNorth, true},
		{Position{6, 5}, DirectionEast, true},
		{Position{6, 5}, DirectionSouth, true},
		{Position{6, 5}, DirectionWest, false}, // oob
	}

	for _, test := range tests {
		canMove := b.CanMove(test.Position, test.Direction)
		if canMove != test.CanMove {
			t.Errorf("expected CanMove(%v, %d) = %v",
				test.Position, test.Direction, test.CanMove)
		}
	}
}

type moveTest struct {
	Start     Position
	Direction Direction
	End       Position
}

func TestBoardMove(t *testing.T) {
	b, _ := NewBoard(10)
	b.AddWall(Position{1, 1}, DirectionNorth)
	b.SetOOB(Position{5, 5})

	tests := []moveTest{
		{Position{0, 0}, DirectionNorth, Position{0, 0}},
		{Position{0, 0}, DirectionEast, Position{9, 0}},
		{Position{0, 0}, DirectionSouth, Position{0, 9}},
		{Position{0, 0}, DirectionWest, Position{0, 0}},

		{Position{1, 0}, DirectionSouth, Position{1, 0}}, // wall in the way
		{Position{1, 9}, DirectionNorth, Position{1, 1}}, // wall in the way

		{Position{1, 5}, DirectionEast, Position{4, 5}}, // oob in the way
	}

	for _, test := range tests {
		end := b.Move(test.Start, test.Direction)
		if !end.Equal(test.End) {
			t.Errorf("expected Move(%v, %d) = %v, got %v",
				test.Start, test.Direction, test.End, end)
		}
	}
}
