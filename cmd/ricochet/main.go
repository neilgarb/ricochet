package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/neilgarb/ricochet"
)

func main() {
	b, s, err := ricochet.ReadBoard(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	if !b.Valid() {
		panic(errors.New("invalid board"))
	}
	ml := s.Solve(ricochet.Token{ricochet.ShapeCircle, ricochet.ColourBlue})
	fmt.Printf("%+v\n", ml)
}
