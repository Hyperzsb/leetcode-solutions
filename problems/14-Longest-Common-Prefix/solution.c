char * longestCommonPrefix(char ** strs, int strsSize){    
    char *prefix = (char *)malloc(210 * sizeof(char));
    
    int minlen = 300;
    for(int i = 0; i < strsSize; i++) {
        int len = strlen(strs[i]);
        minlen = minlen <= len ? minlen : len;
    }
    
    int index = 0;
    int flag = 0;
    while(index < minlen) {
        prefix[index] = strs[0][index];
        for(int i = 1; i < strsSize; i++) {
            if(strs[i][index] != prefix[index]) {
                flag = 1;
                break;
            }
        }
        
        if(flag) {
            break;
        } else {
            index++;
        }
    }
    
    prefix[index] = '\0';
    
    return prefix;
}

