char * reverseWords(char * s){
    int len = strlen(s);
    char *reversedStr = (char*)malloc((len + 5) * sizeof(char));
    
    int start = len - 1, end = len -1, flag = 0, reversedLen = 0;
    
    for(int i = len - 1; i >= 0; i--) {
        if(s[i] == ' ') {
            if (flag == 0) {
                continue;
            } else {
                start = i + 1;
                for(int j = start; j <= end; j++) {
                    reversedStr[reversedLen++] = s[j];
                }
                reversedStr[reversedLen++] = ' ';
                flag = 0;
            }
        } else if (i == 0) {
            if (flag == 0) {
                reversedStr[reversedLen++] = s[0];
                reversedStr[reversedLen++] = ' ';
                flag = 0;
            } else {
                start = i;
                for(int j = start; j <= end; j++) {
                    reversedStr[reversedLen++] = s[j];
                }
                reversedStr[reversedLen++] = ' ';
                flag = 0;
            }
        } else {
            if (flag == 0) {
                end = i;
                flag = 1;
            } else {
                continue;
            }
        }
    }
    
    reversedLen--;
    reversedStr[reversedLen] = '\0';
    
    return reversedStr;
}

