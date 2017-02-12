package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/neilgarb/ricochet"
)

func main() {
	b, _, err := ricochet.ReadBoard(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b.Valid())
}
