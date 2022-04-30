package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pholawat-tle/tic-tac-toe/game"
)


func main() {
  p := tea.NewProgram(
    game.NewGame(),
    tea.WithAltScreen(),
    tea.WithMouseCellMotion(),
  )

  if err := p.Start(); err != nil {
    fmt.Printf("Alas, there's been an error: %v", err)
    os.Exit(1)
  }
}

