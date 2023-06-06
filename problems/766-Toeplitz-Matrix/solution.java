class Solution {
    public boolean check(int x, int y, int[][]matrix) {
        int target = matrix[x][y];
        while (x < matrix.length && y < matrix[0].length) {
            if (matrix[x][y] != target) {
                return false;
            }
            x++;
            y++;
        }

        return true;
    }

    public boolean isToeplitzMatrix(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            if (!check(i, 0, matrix)) {
                return false;
            }
        }

        for (int i = 0; i < matrix[0].length; i++) {
            if (!check(0, i, matrix)) {
                return false;
            }
        }

        return true;
    }
}

