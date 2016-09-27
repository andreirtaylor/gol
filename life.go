package life

import "fmt"

const Alive = "A"
const Dead = "D"
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

// generates a empty board the same size as the receiver
func NewGame(g Game) Game {
	ret := Game{make([][]bool, len(g.Board))}
	for i, row := range g.Board {
		ret.Board[i] = make([]bool, len(row))
	}
	return ret
}

func (g Game) Advance() Game {
	ret := NewGame(g)
	for i, row := range g.Board {
		for j, _ := range row {
			n := g.getNeighbours(i, j)
			switch {
			case n < 2:
				ret.Board[i][j] = false
			case (n == 2 || n == 3) && g.Board[i][j]:
				ret.Board[i][j] = true
			case n > 3:
				ret.Board[i][j] = false
			// reproduction
			case n == 3 && !g.Board[i][j]:
				ret.Board[i][j] = true
			}
		}
	}
	return ret
}

func (g Game) getNeighbours(i, j int) int {
	ret := 0
	for row := max(i-1, 0); row < min(i+2, len(g.Board)); row += 1 {
		for col := max(j-1, 0); col < min(j+2, len(g.Board[row])); col += 1 {
			if ((row != i) || (col != j)) && g.Board[row][col] {
				ret += 1
			}
		}
	}
	return ret
}

// Seriously? Go doesnt have a max or min for integers?
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
