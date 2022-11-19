/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */

void fill(int** image, int rows, int cols, int x, int y, int color) {
    if(image[x][y] == color) {
        return;
    }
    
    int originalColor = image[x][y];
    image[x][y] = color;
    
    if(x > 0 && image[x - 1][y] == originalColor) {
        fill(image, rows, cols, x - 1, y, color);
    }
    
    if(x < rows - 1 && image[x + 1][y] == originalColor) {
        fill(image, rows, cols, x + 1, y, color);
    }
    
    if(y > 0 && image[x][y - 1] == originalColor) {
        fill(image, rows, cols, x, y - 1, color);
    }
    
    if(y < cols - 1 && image[x][y + 1] == originalColor) {
        fill(image, rows, cols, x, y + 1, color);
    }
    
}

int** floodFill(int** image, int imageSize, int* imageColSize, int sr, int sc, int color, int* returnSize, int** returnColumnSizes){
    
    int **returnImage = (int **)malloc(imageSize * sizeof(int *));
    
    for(int i = 0; i < imageSize; i++) {
        returnImage[i] = (int *)malloc(imageColSize[i] * sizeof(int));
        for(int j = 0; j < imageColSize[i]; j++) {
            returnImage[i][j] = image[i][j];
        }
    }

    *returnSize = imageSize;
    
    int *returnImageColSize = (int *)malloc(imageSize * sizeof(int));
    for(int i = 0; i < imageSize; i++) {
        returnImageColSize[i] = imageColSize[i];
    }
    *returnColumnSizes = returnImageColSize;
    
    fill(returnImage, imageSize, imageColSize[0], sr, sc, color);
    
    return returnImage;
}

