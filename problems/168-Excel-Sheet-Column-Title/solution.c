char * convertToTitle(int columnNumber){
    char *result = (char *)malloc(8 * sizeof(char));

    int index = 0, end = 0;
    while(columnNumber > 0) {
        columnNumber--;
        end = columnNumber % 26;
        result[index++] = (char)('A' + end);
        columnNumber -= end;
        columnNumber /= 26;
    }

    result[index] = '\0';

    for(int i = 0; i < index / 2; i++) {
        char temp = result[i];
        result[i] = result[index - 1 - i];
        result[index - 1 - i] = temp;
    }

    return result;
}

