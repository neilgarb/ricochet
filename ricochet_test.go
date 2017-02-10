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

func TestBoardSetSize(t *testing.T) {
	b := NewBoard()

	if err := b.SetSize(-1); err == nil {
		t.Errorf("expected error")
	}
	if err := b.SetSize(0); err == nil {
		t.Errorf("expected error")
	}
	if err := b.SetSize(1); err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if err := b.SetSize(101); err == nil {
		t.Errorf("expected error")
	}
}

func TestBoardOOB(t *testing.T) {
	b := NewBoard()
	b.SetSize(10)

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
	b := NewBoard()
	b.SetSize(10)

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
	b := NewBoard()
	b.SetSize(10)

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
	b := NewBoard()
	b.SetSize(10)

	if b.Valid() {
		t.Errorf("expected invalid")
	}

	tot := len(allShapes) * len(allColours)
	i := 0
	for _, s := range allShapes {
		for _, c := range allColours {
			i++
			b.AddSink(Token{s, c}, Position{i / 10, i % 10})
			if i < tot {
				if b.Valid() {
					t.Errorf("expected invalid")
				}
			}
		}
	}

	if !b.Valid() {
		t.Errorf("expected valid")
	}
}
