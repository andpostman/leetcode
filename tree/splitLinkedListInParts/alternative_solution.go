package splitLinkedListInParts

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)

	// create matrix and fill by defaul -1
	for i := 0; i < m; i++ {
		r := make([]int, n)
		for j := 0; j < n; j++ {
			r[j] = -1
		}
		res[i] = r
	}

	// set starting direction
	dir := 'r'

	// set starting position
	i, j := 0, 0

	for head != nil {
		res[i][j] = head.Val

		// check if need change direction
		if dir == 'r' {
			if j+1 >= len(res[i]) || res[i][j+1] != -1 {
				dir = 'd'
			}
		} else if dir == 'd' {
			if i+1 >= len(res) || res[i+1][j] != -1 {
				dir = 'l'
			}
		} else if dir == 'l' {
			if j-1 < 0 || res[i][j-1] != -1 {
				dir = 't'
			}
		} else if dir == 't' {
			if i-1 < 0 || res[i-1][j] != -1 {
				dir = 'r'
			}
		}

		head = head.Next

		// change position in the matrix
		switch {
		case dir == 'r':
			j++
		case dir == 'd':
			i++
		case dir == 'l':
			j--
		case dir == 't':
			i--
		}
	}

	return res
}
