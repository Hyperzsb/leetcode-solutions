/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* spiralOrder(int** matrix, int matrixSize, int* matrixColSize, int* returnSize){
    int *result = (int *)malloc(matrixSize * matrixColSize[0] * sizeof(int));
    memset(result, 0, matrixSize * matrixColSize[0] * sizeof(int));
    *returnSize = matrixSize * matrixColSize[0];
    
    int x = 0, y = 0, index = 0, direction = 1;
    
    result[index++] = matrix[x][y];
    matrix[x][y] = 10000;
    while(true) {
        if(direction == 1) {
            if(y < matrixColSize[0] - 1 && matrix[x][y + 1] != 10000) {
                result[index++] = matrix[x][y + 1];
                matrix[x][y + 1] = 10000;
                y++;
            } else if(x < matrixSize - 1 && matrix[x + 1][y] != 10000) {
                result[index++] = matrix[x + 1][y];
                matrix[x + 1][y] = 10000;
                x++;
                direction = 2;
            } else {
                break;
            }
        } else if(direction == 2) {
            if(x < matrixSize - 1 && matrix[x + 1][y] != 10000) {
                result[index++] = matrix[x + 1][y];
                matrix[x + 1][y] = 10000;
                x++;
            } else if(y > 0 && matrix[x][y - 1] != 10000) {
                result[index++] = matrix[x][y - 1];
                matrix[x][y - 1] = 10000;
                direction = 3;
                y--;
            } else {
                break;
            }
        } else if(direction == 3) {
            if(y > 0 && matrix[x][y - 1] != 10000) {
                result[index++] = matrix[x][y - 1];
                matrix[x][y - 1] = 10000;
                y--;
            } else if(x > 0 && matrix[x - 1][y] != 10000) {
                result[index++] = matrix[x - 1][y];
                matrix[x - 1][y] = 10000;
                x--;
                direction = 4;
            } else {
                break;
            }
        } else {
            if(x > 0 && matrix[x - 1][y] != 10000) {
                result[index++] = matrix[x - 1][y];
                matrix[x - 1][y] = 10000;
                x--;
            } else if(y < matrixColSize[0] - 1 && matrix[x][y + 1] != 10000) {
                result[index++] = matrix[x][y + 1];
                matrix[x][y + 1] = 10000;
                y++;
                direction = 1;
            } else {
                break;
            }
        }
    }
    
    return result;
}

