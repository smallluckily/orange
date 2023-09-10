package main

import "fmt"

type Node struct {
	Row   int
	Joe   int
	Value int
}

var ChessMap1 []Node

func main() {
	var ChessMap [11][11]int
	ChessMap[1][2] = 1
	ChessMap[2][3] = 1
	for i, v1 := range ChessMap {
		fmt.Println(v1)
		for j, v2 := range v1 {
			if v2 != 0 {
				chessMap := Node{
					Row:   i,
					Joe:   j,
					Value: v2,
				}
				ChessMap1 = append(ChessMap1, chessMap)

			}

		}

	}
	for i, Node := range ChessMap1 {
		fmt.Printf("%d: %d %d %d\n ", i, Node.Row, Node.Joe, Node.Value)
	}
}
