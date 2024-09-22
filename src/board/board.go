package board

import (
	"reflect"
)

const (
	Empty = iota
	Black
	White
)

type Board struct {
	size int
	grid [][]int
}

func (b *Board) PlaceStone(x, y, color int) bool {
	if b.grid[x][y] != Empty {
		return false
	}
	b.grid[x][y] = color
	return true
}

func (b *Board) HasLiberty(x, y, color int) bool {
	visited := make(map[[2]int]bool)
	return b.checkLiberty(x, y, color, visited)
}

func (b *Board) checkLiberty(x, y, color int, visited map[[2]int]bool) bool {
	if b.grid[x][y] == Empty {
		return true
	}
	visited[[2]int{x, y}] = true

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < b.size && ny < b.size && !visited[[2]int{nx, ny}] {
			if b.grid[nx][ny] == Empty || (b.grid[nx][ny] == color && b.checkLiberty(nx, ny, color, visited)) {
				return true
			}
		}
	}
	return false
}

type Game struct {
	board   Board
	history [][][]int
}

func (g *Game) SaveState() {
	snapshot := make([][]int, g.board.size)
	for i := range g.board.grid {
		snapshot[i] = append([]int{}, g.board.grid[i]...)
	}
	g.history = append(g.history, snapshot)
}

func (g *Game) IsKo() bool {
	current := g.board.grid
	for _, past := range g.history {
		if reflect.DeepEqual(current, past) {
			return true
		}
	}
	return false
}
