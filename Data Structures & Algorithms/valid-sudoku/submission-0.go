func isValidSudoku(board [][]byte) bool {
	// same row, all columns
	rm := make(map[byte]struct{})
	// same column, all rows
	cm := make(map[byte]struct{})
	for r := 0; r < 9; r++ {
		clear(rm)
		clear(cm)
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				if _, ok := rm[board[r][c]]; ok {
					return false
				} else {
					rm[board[r][c]] = struct{}{}
				}
			}

			if board[c][r] != '.' {
				if _, ok := cm[board[c][r]]; ok {
					return false
				} else {
					cm[board[c][r]] = struct{}{}
				}
			}
		}
	}

	// squares
	sm := make(map[byte]struct{})
	for cnt := 0; cnt < 9; cnt += 3 {
		clear(sm)

		// first 3 columns
		for c := 0; c < 3; c++ {
			if !validSquare(cnt, sm, board, c) {
				return false
			}
		}

		clear(sm)
		// next 3 columns
		for c := 3; c < 6; c++ {
			if !validSquare(cnt, sm, board, c) {
				return false
			}
		}

		clear(sm)
		// last 3 columns
		for c := 6; c < 9; c++ {
			if !validSquare(cnt, sm, board, c) {
				return false
			}
		}

	}

	return true
}

func validSquare(cnt int, sm map[byte]struct{}, board [][]byte, c int) bool {
	for r := cnt; r < cnt+3; r++ {
		if board[r][c] != '.' {
			if _, ok := sm[board[r][c]]; ok {
				return false
			} else {
				sm[board[r][c]] = struct{}{}
			}
		}
	}
	return true
}
