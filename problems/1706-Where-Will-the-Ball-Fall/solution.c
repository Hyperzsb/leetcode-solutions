/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int fall(int **grid, int **flags, int m, int n, int x, int y) {
    if(flags[x][y] != -2) {
        return flags[x][y];
    }
    
    if(grid[x][y] == 1) {
        if(y == n -1) {
            flags[x][y] = -1;
        } else if(grid[x][y + 1] == -1) {
            flags[x][y] = -1;
        } else if(x == m - 1) {
            flags[x][y] = y + 1;
        } else {
            flags[x][y] = fall(grid, flags, m, n, x + 1, y + 1);
        }
    } else {
        if(y == 0) {
            flags[x][y] = -1;
        } else if(grid[x][y - 1] == 1) {
            flags[x][y] = -1;
        } else if(x == m - 1) {
            flags[x][y] = y - 1;
        } else {
            flags[x][y] = fall(grid, flags, m, n, x + 1, y - 1);
        }
    }
    
    return flags[x][y];
}

int* findBall(int** grid, int gridSize, int* gridColSize, int* returnSize){
    int **flags = (int **)malloc(gridSize * sizeof(int *));
    for(int i = 0; i < gridSize; i++) {
        flags[i] = (int *)malloc(gridColSize[0] * sizeof(int));
        for(int j = 0; j < gridColSize[0]; j++) {
            flags[i][j] = -2;
        }
    }
    
    int *result = (int *)malloc(gridColSize[0] * sizeof(int));
    for(int i = 0; i < gridColSize[0]; i++) {
        result[i] = fall(grid, flags, gridSize, gridColSize[0], 0, i);
    }
    
    *returnSize = gridColSize[0];
    return result;
}

