class Solution {
    public int minFallingPathSum(int[][] matrix) {
        if (matrix.length == 1) {
            return matrix[0][0];
        }

        for (int i = 1; i < matrix.length; i++) {
            matrix[i][0] = Math.min(matrix[i-1][0], matrix[i-1][1]) + matrix[i][0];
            
            for (int j = 1; j < matrix[i].length-1; j++) {
                matrix[i][j] = Math.min(Math.min(matrix[i-1][j-1], matrix[i-1][j+1]), matrix[i-1][j]) + matrix[i][j];
            }

            matrix[i][matrix[i].length-1] = Math.min(matrix[i-1][matrix[i-1].length-2], matrix[i-1][matrix[i-1].length-1]) + matrix[i][matrix[i].length-1];
        }

        int result = Integer.MAX_VALUE;
        for (int i = 0; i < matrix[matrix.length-1].length; i++) {
            result = Math.min(result, matrix[matrix.length-1][i]);
        }

        return result;
    }
}

