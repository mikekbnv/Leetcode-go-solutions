package count_servers_that_communicate_1267

func countServers(grid [][]int) int {
	rez := 0
	a := make([][]bool, len(grid))
	for i := range a {
		a[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				for k := 0; k < len(grid); k++ {
					if grid[k][j] == 1 && k != i {
						a[i][j] = true
						break
					}
				}
				for r := 0; r < len(grid[0]); r++ {
					if grid[i][r] == 1 && r != j {
						a[i][j] = true
						break
					}
				}
			}
		}
	}
	for _, v := range a {
		for _, k := range v {
			if k == true {
				rez++
			}
		}
	}
	return rez
}
