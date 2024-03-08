package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const game_data_dir = "games"

type PinmanGame struct {
	board_no                  int
	board                     []string
	board_width, board_height int
}

func (g *PinmanGame) load() {
	file_name := fmt.Sprintf("game_%02d.txt", g.board_no)
	f, err := os.Open(filepath.Join(game_data_dir, file_name))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g.board_width = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Clone(scanner.Text())
		g.board_width = max(g.board_width, len(line))
		g.board = append(g.board, line)
	}

	g.board_height = len(g.board)

}
