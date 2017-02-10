package ricochet

import "testing"

func TestPositionEqual(t *testing.T) {
	if !(Position{0, 0}).Equal(Position{0, 0}) {
		t.Errorf("expected equal")
	}
	if (Position{10, 10}).Equal(Position{10, 11}) {
		t.Errorf("expected not equal")
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

	w := Wall{Position{-1, -1}, DirectionNorth}
	if err := b.AddWall(w); err == nil {
		t.Errorf("expected error")
	}
	w = Wall{Position{0, 0}, DirectionNorth}
	if err := b.AddWall(w); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := b.AddWall(w); err == nil {
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

func TestAddRobot(t *testing.T) {
	b, _ := NewBoard(10)

	r := Robot{ColourRed}
	if err := b.AddRobot(r, Position{-1, -1}); err == nil {
		t.Errorf("expected error")
	}
	if err := b.AddRobot(r, Position{0, 0}); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := b.AddRobot(r, Position{0, 1}); err == nil {
		t.Errorf("expected error")
	}

	r = Robot{ColourBlue}
	if err := b.AddRobot(r, Position{0, 0}); err == nil {
		t.Errorf("expected error")
	}
	if err := b.AddRobot(r, Position{0, 1}); err != nil {
		t.Errorf("expected success, got %v", err)
	}
}

func TestValid(t *testing.T) {
	b, _ := NewBoard(10)

	if b.Valid() {
		t.Errorf("expected invalid")
	}

	i := 0
	for _, s := range allShapes {
		for _, c := range allColours {
			i++
			b.AddSink(Token{s, c}, Position{i / 10, i % 10})
			if b.Valid() {
				t.Errorf("expected invalid")
			}
		}
	}

	for i := 0; i < minRobots; i++ {
		b.AddRobot(Robot{Colour(i)}, Position{i / 10, i % 10})
		if i < minRobots-1 {
			if b.Valid() {
				t.Errorf("expected invalid")
			}
		}
	}

	if !b.Valid() {
		t.Errorf("expected valid")
	}
}
