class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int rS = 0, rE = matrix.length - 1;
        while (rS < rE) {
            int rH = (rS + rE) / 2;

            if (target <= matrix[rH][0]) {
                rE = rH;
            } else {
                rS = rH + 1;
            }
        }

        if (target < matrix[rS][0]) {
            rS--;
        }

        if (rS < 0) {
            return false;
        }

        int cS = 0, cE = matrix[rS].length - 1;
        while (cS < cE) {
            int cH = (cS + cE) / 2;

            if (target <= matrix[rS][cH]) {
                cE = cH;
            } else {
                cS = cH + 1;
            }
        }

        if (target != matrix[rS][cS]) {
            return false;
        } else {
            return true;
        }
    }
}

