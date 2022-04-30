package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pholawat-tle/tic-tac-toe/border"
)
type Game struct {
  squares [9]string

  status string
}

func NewGame() *Game {
  return &Game{
    squares: [9]string{},
  }
}

func (m *Game) Init() tea.Cmd {
 return nil
}

func (m *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  winner := calculateWinner(m.squares)
  nextValue := calculateNextValue(m.squares)

  switch msg := msg.(type) {

  case tea.MouseMsg:
    if msg.Type != tea.MouseLeft &&  winner == "" {
      selectedSquare := border.Cell(msg.X, msg.Y)
      return m.Select(selectedSquare, nextValue)
    }
    return m, nil

  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    case "r":
      m.squares = [9]string{}
      return m, nil
    }

  }
  return m, nil
}

func (m *Game) Select(square int, nextValue string) (tea.Model, tea.Cmd){
  if square < 0 || square >= 9 {
    return m, nil
  }

  if m.squares[square] == "" {
    m.squares[square] = nextValue
  }

  return m, nil
}

func (m *Game) View() string {
  winner := calculateWinner(m.squares)
  nextValue := calculateNextValue(m.squares)
  m.status = calculateStatus(winner, m.squares, nextValue)

  s := fmt.Sprintf("\n %s\n%s\n %s", 
    m.status,
    border.Grid(m.squares),
    "Press q to Exit",
  )

  return s
}

func calculateWinner(squares [9]string) string {
  lines := [][]int{
    {0, 1, 2},
    {3, 4, 5},
    {6, 7, 8},
    {0, 3, 6},
    {1, 4, 7},
    {2, 5, 8},
    {0, 4, 8},
    {2, 4, 6},
  }

  for _, line := range lines {
    a, b, c := line[0], line[1], line[2]
    if squares[a] != "" && squares[a] == squares[b] && squares[a] == squares[c] {
      return squares[a]
    }
  }

  return ""
}

func turnsPlayed(squares [9]string) int {
  turns := 0

  for _, square := range squares {
    if square != "" {
      turns += 1
    }
  }

  return turns
}

func calculateNextValue(squares [9]string) string {
  turnsPlayed := turnsPlayed(squares)
  if turnsPlayed % 2 == 0 {
    return "X"
  }
  return "O"
}

func calculateStatus(winner string, squares [9]string, nextValue string) string {
  if winner != "" {
    return fmt.Sprintf("%s Wins! Press R to Restart", winner)
  }
  if turnsPlayed(squares) == 9 {
    return "Draw! Press R to Restart"
  }

  return fmt.Sprintf("Next Player: %s", nextValue)
}
