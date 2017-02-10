package ricochet

import "errors"

type Direction int

const (
	DirectionNorth Direction = 0
	DirectionEast  Direction = 1
	DirectionSouth Direction = 2
	DirectionWest  Direction = 3
)

type Shape int

const (
	ShapeCircle   Shape = 0
	ShapeTriangle Shape = 1
	ShapeDiamond  Shape = 2
	ShapeHexagon  Shape = 3
)

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

type Wall struct {
	Position  Position
	Direction Direction
}

type Robot struct {
	Colour Colour
}

const minRobots = 4

type Board struct {
	size   int                // the width or height of the board
	walls  []Wall             // positions and directions of walls on the board
	sinks  map[Token]Position // positions and types of tokens on the board
	oob    []Position         // positions which are out of bounds
	robots map[Robot]Position // positions of robots
}

func NewBoard() *Board {
	return &Board{
		sinks:  make(map[Token]Position),
		robots: make(map[Robot]Position),
	}
}

func (b *Board) SetSize(size int) error {
	if size < 1 || size > 100 {
		return errors.New("invalid board size")
	}
	b.size = size
	return nil
}

func (b *Board) InBounds(p Position) bool {
	if p.X < 0 || p.X >= b.size {
		return false
	}
	if p.Y < 0 || p.Y >= b.size {
		return false
	}
	for _, o := range b.oob {
		if o.Equal(p) {
			return false
		}
	}
	return true
}

func (b *Board) SetOOB(pos Position) error {
	if !b.InBounds(pos) {
		return errors.New("position already oob")
	}
	b.oob = append(b.oob, pos)
	return nil
}

func (b *Board) AddWall(wall Wall) error {
	if !b.InBounds(wall.Position) {
		return errors.New("wall out of bounds")
	}
	for _, w := range b.walls {
		if w.Position.Equal(wall.Position) && w.Direction == wall.Direction {
			return errors.New("duplicate wall")
		}
	}
	b.walls = append(b.walls, wall)
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

func (b *Board) AddRobot(robot Robot, pos Position) error {
	if !b.InBounds(pos) {
		return errors.New("position is oob")
	}

	for r, p := range b.robots {
		if r.Colour == robot.Colour {
			return errors.New("robot already added")
		}
		if p.Equal(pos) {
			return errors.New("position already has a robot")
		}
	}

	b.robots[robot] = pos
	return nil
}

// Valid returns true if the board is correctly configured.
func (b *Board) Valid() bool {
	// Make sure all tokens are placed.
	for _, s := range allShapes {
		for _, c := range allColours {
			t := Token{s, c}
			if _, ok := b.sinks[t]; !ok {
				return false
			}
		}
	}

	// Make sure there are at least four robots.
	if len(b.robots) < minRobots {
		return false
	}

	return true
}
