package main

type Space struct {
	Value  int
	Marked bool
}

type Board struct {
	Spaces [][]*Space
}

func (b *Board) IsWinner() bool {
	// Across
	for _, row := range b.Spaces {
		win := true
		for _, space := range row {
			if !space.Marked {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Down
	for col := 0; col < len(b.Spaces[0]); col++ {
		win := true
		for row := 0; row < len(b.Spaces); row++ {
			if !b.Spaces[row][col].Marked {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Diagonal
	for i := 0; i < len(b.Spaces); i++ {
		if !b.Spaces[i][i].Marked {
			return false
		}
	}

	return true
}

func (b *Board) Mark(num int) bool {
	for _, row := range b.Spaces {
		for _, space := range row {
			if space.Value == num {
				space.Marked = true
				return true
			}
		}
	}

	return false
}

func (b *Board) Score() int {
	score := 0
	for _, row := range b.Spaces {
		for _, space := range row {
			if !space.Marked {
				score += space.Value
			}
		}
	}

	return score
}
