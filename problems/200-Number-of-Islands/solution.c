void fill(char **grid, int rows, int cols, int x, int y) {
    grid[x][y] = '0';
    
    if(x > 0 && grid[x - 1][y] == '1') {
        fill(grid, rows, cols, x - 1, y);
    }
    
    if(x < rows - 1 && grid[x + 1][y] == '1') {
        fill(grid, rows, cols, x + 1, y);
    }
    
    if(y > 0 && grid[x][y - 1] == '1') {
        fill(grid, rows, cols, x, y - 1);
    }
    
    if(y < cols - 1 && grid[x][y + 1] == '1') {
        fill(grid, rows, cols, x, y + 1);
    }
}

int numIslands(char** grid, int gridSize, int* gridColSize){
    int rows = gridSize, cols = gridColSize[0];
    int islands = 0;
    
    for(int i = 0; i < rows; i++) {
        for(int j = 0; j < cols; j++) {
            if(grid[i][j] == '1') {
                islands++;
                fill(grid, rows, cols, i, j);
            }
        }
    }
    
    return islands;
}

