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

func (g Game) Advance() Game {
	var ret Game

	for i, row := range g.Board {
		for j, _ := range row {
			g.getNeighbours(i, j)

		}
	}
	return ret
}

func (g Game) getNeighbours(i, j int) int {
	ret := 0
	for row := Max(i-1, 0); row < Min(i+2, len(g.Board)); row += 1 {
		for col := Max(j-1, 0); col < Min(j+2, len(g.Board[row])); col += 1 {
			if ((row != i) || (col != j)) && g.Board[row][col] {
				ret += 1
			}
		}
	}
	return ret
}

// Seriously? Go doesnt have a max or min for integers?
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
