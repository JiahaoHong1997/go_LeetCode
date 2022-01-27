package ci04

func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    m, n := len(matrix), len(matrix[0])
    if target < matrix[0][0] || target > matrix[m-1][n-1] {
        return false
    }

    var findFirstRow func() int
    findFirstRow = func() int {
        if matrix[0][n-1] >= target {
            return 0
        }
        l, r := 0, m-1
        for l < r {
            mid := l + (r-l)/2
            if matrix[mid][n-1] == target {
                l = mid + 1
            } else if matrix[mid][n-1] < target {
                l = mid + 1
            } else {
                r = mid
            }
        }
        if matrix[l][n-1] == target {
            return l
        } 
        return l-1
    }

    first := findFirstRow()
    if matrix[first][n-1] == target {
        return true
    }

    var findInRow func(int) bool 
    findInRow = func(i int) bool {
        l, r := 0, n-1
        for l <= r {
            mid := l + (r-l)/2
            if matrix[i][mid] == target {
                return true
            } else if matrix[i][mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        return false
    }

    for i:=first; i < m && matrix[i][0] <= target; i++ {
        if findInRow(i) {
            return true
        }
    }
    return false
}