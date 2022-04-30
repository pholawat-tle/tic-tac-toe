package border

import "fmt"

const (
	// cellHeight represents how many rows are in a cell
	cellHeight = 2
	// cellWidth represents how many columns are in a cell
	cellWidth = 4

	// marginTop and marginLeft represent the offset of the grid
	// board from the top left of the terminal window.
	marginTop  = 2
	marginLeft  = 1

  // numberOfColumns represents the number of columns in the grid
  numberOfColumns = 3
  // numberOfRows represents the number of rows in the grid
  numberOfRows = 3
)

func Cell(x, y int) int {
  outOfBound := cellWidth * numberOfColumns <= (x - marginLeft) || cellHeight * numberOfRows <= (y - marginTop)
  if outOfBound {
    return -1
  }

	col := (x - marginLeft) / cellWidth
	row := (y - marginTop) / cellHeight
	return (row * 3)+ col
}

func Grid(squares [9]string) string {
  displaySquares := [9]string{}

  for i, square := range squares {
    squareIsEmpty := square == ""

    if squareIsEmpty {
      displaySquares[i] = " "
    } else {
      displaySquares[i] = square
    }

  }

  s := fmt.Sprintf(
` ┌───┬───┬───┐
 │ %s │ %s │ %s │ 
 ├───├───├───┤
 │ %s │ %s │ %s │
 ├───├───├───┤
 │ %s │ %s │ %s │
 └───┴───┴───┘
`, 
    displaySquares[0],
    displaySquares[1],
    displaySquares[2],
    displaySquares[3],
    displaySquares[4],
    displaySquares[5],
    displaySquares[6],
    displaySquares[7],
    displaySquares[8],
  )

  return s
}
