package solves

func IsToeplitzMatrix(matrix [][]int) bool {
	subLen := len(matrix[0])
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < subLen-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
