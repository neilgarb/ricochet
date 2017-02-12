package ricochet

import "errors"

type Direction int

const (
	DirectionNorth Direction = 0
	DirectionEast  Direction = 1
	DirectionSouth Direction = 2
	DirectionWest  Direction = 3
)

var allDirections = []Direction{DirectionNorth, DirectionEast, DirectionSouth,
	DirectionWest}

func (d Direction) Valid() bool {
	return d >= 0 && d <= 3
}

// Flip returns the opposite direction to `d`.
func (d Direction) Flip() Direction {
	return (d + 2) % 4
}

type Shape int

const (
	ShapeCircle   Shape = 0
	ShapeTriangle Shape = 1
	ShapeDiamond  Shape = 2
	ShapeHexagon  Shape = 3
)

func (s Shape) Valid() bool {
	return s >= 0 && s <= 3
}

var allShapes = []Shape{ShapeCircle, ShapeTriangle, ShapeDiamond, ShapeHexagon}

type Colour int

const (
	ColourBlue   Colour = 0
	ColourYellow Colour = 1
	ColourGreen  Colour = 2
	ColourRed    Colour = 3

	// Additional robot colours:
	ColourSilver Colour = 10
)

func (c Colour) ValidForToken() bool {
	return c >= 0 && c <= 3
}

var allColours = []Colour{ColourBlue, ColourYellow, ColourGreen, ColourRed}

type Token struct {
	Shape  Shape
	Colour Colour
}

type Position struct {
	X int
	Y int
}

func (p Position) Equal(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p Position) Next(dir Direction) Position {
	switch dir {
	case DirectionNorth:
		return Position{p.X, p.Y - 1}
	case DirectionEast:
		return Position{p.X + 1, p.Y}
	case DirectionSouth:
		return Position{p.X, p.Y + 1}
	case DirectionWest:
		return Position{p.X - 1, p.Y}
	}
	return p
}

type Block struct {
	oob   bool
	walls map[Direction]bool
}

func NewBlock() Block {
	return Block{walls: make(map[Direction]bool)}
}

type Robot struct {
	Colour Colour
}

type State struct {
	board  *Board
	robots map[Position]Robot
}

func (s *State) AddRobot(pos Position, robot Robot) error {
	if !s.board.InBounds(pos) {
		return errors.New("position out of bounds")
	}

	for p, r := range s.robots {
		if p.Equal(pos) {
			return errors.New("position already has a robot")
		}
		if r.Colour == robot.Colour {
			return errors.New("robot already added")
		}
	}

	s.robots[pos] = robot
	return nil
}

// CanMove returns true if a robot can move from the given position in the given
// direction.
func (s *State) CanMove(pos Position, dir Direction) bool {
	next := pos.Next(dir)

	// Next block is OOB...
	if !s.board.InBounds(next) {
		return false
	}

	// There's a wall in the way...
	if s.board.blocks[pos].walls[dir] {
		return false
	}
	if s.board.blocks[next].walls[dir.Flip()] {
		return false
	}

	// There's a robot in the way...
	if _, ok := s.robots[next]; ok {
		return false
	}

	return true
}

// Move returns the position a robot would end up in if it started in `pos` and
// moved in direction `dir`.
func (s *State) Move(pos Position, dir Direction) Position {
	for {
		if !s.CanMove(pos, dir) {
			return pos
		}
		pos = pos.Next(dir)
	}
}

const minRobots = 4

type Board struct {
	size   int                // the width or height of the board
	blocks map[Position]Block // positions of blocks of interest
	sinks  map[Token]Position // positions and types of tokens on the board
}

func NewBoard(size int) (*Board, error) {
	if size < 1 || size > 100 {
		return nil, errors.New("invalid board size")
	}
	return &Board{
		size:   size,
		blocks: make(map[Position]Block),
		sinks:  make(map[Token]Position),
	}, nil
}

func (b *Board) NewState() *State {
	return &State{b, make(map[Position]Robot)}
}

func (b *Board) getBlock(pos Position) Block {
	block, ok := b.blocks[pos]
	if !ok {
		return NewBlock()
	}
	return block
}

func (b *Board) InBounds(p Position) bool {
	if p.X < 0 || p.X >= b.size {
		return false
	}
	if p.Y < 0 || p.Y >= b.size {
		return false
	}
	for pos, block := range b.blocks {
		if pos.Equal(p) && block.oob {
			return false
		}
	}
	return true
}

func (b *Board) SetOOB(pos Position) error {
	if !b.InBounds(pos) {
		return errors.New("position already oob")
	}
	block := b.getBlock(pos)
	block.oob = true
	b.blocks[pos] = block
	return nil
}

func (b *Board) AddWall(pos Position, dir Direction) error {
	if !b.InBounds(pos) {
		return errors.New("wall out of bounds")
	}

	block := b.getBlock(pos)
	if block.walls[dir] {
		return errors.New("duplicate wall")
	}

	block.walls[dir] = true
	b.blocks[pos] = block

	return nil
}

func (b *Board) AddSink(token Token, pos Position) error {
	if _, ok := b.sinks[token]; ok {
		return errors.New("token is already on board")
	}
	if !b.InBounds(pos) {
		return errors.New("position is oob")
	}
	for _, p := range b.sinks {
		if p.Equal(pos) {
			return errors.New("position already has a sink")
		}
	}
	b.sinks[token] = pos
	return nil
}

// Valid returns true if the board is correctly configured.
func (b *Board) Valid() bool {
	for _, s := range allShapes {
		for _, c := range allColours {
			t := Token{s, c}
			if _, ok := b.sinks[t]; !ok {
				return false
			}
		}
	}
	return true
}
