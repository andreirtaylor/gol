package gol

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshaling(t *testing.T) {
	cases := []struct {
		input, output []byte
	}{
		{
			[]byte(`{"Board":[[true],[false]]}`),
			[]byte(`{"Board":[[true],[false]]}`),
		},
		{
			[]byte(`{"shouldBeNull": "shit"}`),
			[]byte(`{"Board":null}`),
		},
	}

	for _, c := range cases {
		var g Game
		err := json.Unmarshal(c.input, &g)
		if err != nil {
			t.Errorf("Error while unmarshaling: %s", err.Error())
			continue
		}
		unmarshaled, err := json.Marshal(g)
		if err != nil {
			t.Errorf("Error while marshaling: %s", err.Error())
			continue
		}
		if string(unmarshaled) != string(c.output) {
			t.Errorf("Expected %s to be %s", unmarshaled, c.output)
		}
	}
}

func TestPrinting(t *testing.T) {
	cases := []struct {
		input  []byte
		output string
	}{
		{
			[]byte(`{"Board":[[true],[false]]}`),
			fmt.Sprintf("%s%s%s", Alive, Sep, Dead),
		},
		{
			[]byte(`{"Board":[[true, false],[true, false]]}`),
			fmt.Sprintf("%s%s%s%s%s", Alive, Dead, Sep, Alive, Dead),
		},
		{
			[]byte(`{"Board":[]}`),
			"",
		},
	}

	for _, c := range cases {
		var g Game
		err := json.Unmarshal(c.input, &g)
		if err != nil {
			t.Errorf("Error while unmarshaling: %s", err.Error())
			continue
		}
		if g.String() != c.output {
			t.Errorf("Expected '%s' to be '%s'", g.String(), c.output)
		}
	}
}

func TestGetNeighbours(t *testing.T) {
	cases := []struct {
		input  [][]bool
		output int
	}{
		{
			[][]bool{
				[]bool{true, false},
				[]bool{false, false}},
			1,
		},
		{
			[][]bool{
				[]bool{false, false, false},
				[]bool{false, false, false},
				[]bool{false, false, false}},
			0,
		},
		{
			[][]bool{
				[]bool{true, true, true},
				[]bool{true, false, true},
				[]bool{true, true, true}},
			8,
		},
		{
			[][]bool{
				[]bool{true, true, true},
				[]bool{true, true, true},
				[]bool{true, true, true}},
			8,
		},
	}

	for i, c := range cases {
		var g Game
		g.Board = c.input
		neighbours := g.getNeighbours(1, 1)
		if neighbours != c.output {
			t.Errorf("Expected '%d' to be '%d'. Test: %d", neighbours, c.output, i)
		}
	}
}

func TestAdvance(t *testing.T) {
	cases := []struct {
		input, output [][]bool
	}{
		{
			[][]bool{
				[]bool{true, false},
				[]bool{false, false}},
			[][]bool{
				[]bool{false, false},
				[]bool{false, false}},
		},
		{
			[][]bool{
				[]bool{true, true},
				[]bool{true, false}},
			[][]bool{
				[]bool{true, true},
				[]bool{true, true}},
		},
		{
			[][]bool{
				[]bool{true, true, false},
				[]bool{true, true, false},
				[]bool{true, false, true}},
			[][]bool{
				[]bool{true, true, false},
				[]bool{false, false, true},
				[]bool{true, false, false}},
		},
	}

	for i, c := range cases {
		var g Game
		g.Board = c.input
		g2 := g.Advance()
		g.Board = c.output
		if g.String() != g2.String() {
			t.Errorf(
				"Expected \n'%s'\n to be \n'%s'\n. Test: %d",
				g2.String(),
				g.String(),
				i)
		}
	}
}
