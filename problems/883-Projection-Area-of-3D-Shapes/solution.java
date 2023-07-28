class Solution {
    public int projectionArea(int[][] grid) {
        int xArea = 0, yArea = 0, zArea = 0;

        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] != 0) {
                    xArea++;
                }
            }
        }

        for (int i = 0; i < grid.length; i++) {
            int highest = 0;
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] > highest) {
                    highest = grid[i][j];
                }
            }

            yArea += highest;
        }

        for (int j = 0; j < grid[0].length; j++) {
            int highest = 0;
            for (int i = 0; i < grid.length; i++) {
                if (grid[i][j] > highest) {
                    highest = grid[i][j];
                }
            }

            xArea += highest;
        }

        return xArea + yArea + zArea;
    }
}

