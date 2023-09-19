type Row struct {
    Idx int
    Soldier int
}

func kWeakestRows(mat [][]int, k int) []int {
    rows := make([]Row, len(mat))
    for i := range rows {
        soldier := len(mat[i])
        for j := range mat[i] {
            if mat[i][j] == 0 {
                soldier = j
                break
            }
        }

        rows[i] = Row{
            Idx: i,
            Soldier: soldier,
        }
    }

    sort.Slice(rows, func(a, b int) bool {
        if rows[a].Soldier < rows[b].Soldier {
            return true
        } else if rows[a].Soldier > rows[b].Soldier {
            return false;
        } else {
            return rows[a].Idx < rows[b].Idx
        }
    })

    result := make([]int, k)
    for i := range result {
        result[i] = rows[i].Idx
    }

    return result
}

