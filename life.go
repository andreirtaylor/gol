package main

import (
	"encoding/json"
	"fmt"
)

type Game struct {
	Board [][]bool
}

func main() {
	j := []byte(`{"Board": [[true], [false]]}`)
	var g Game
	json.Unmarshal(j, &g)
	output, err := json.Marshal(g)
	if err != nil {
		fmt.Println("error: %s", err.Error())
	}
	fmt.Printf("output: %s", output)
}
