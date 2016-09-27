package life

import "fmt"

const Alive = "\u9635"
const Dead = "\u9633"
const Sep = "\n"

// Represents a game object
type Game struct {
	Board [][]bool
}

func (g Game) String() string {
	ret := ""
	for ind, row := range g.Board {
		for _, cell := range row {
			if cell {
				ret = fmt.Sprintf("%s%s", ret, Alive)
			} else {
				ret = fmt.Sprintf("%s%s", ret, Dead)
			}
		}
		if ind != len(g.Board)-1 {
			ret = fmt.Sprintf("%s%s", ret, Sep)
		}
	}
	return ret
}
