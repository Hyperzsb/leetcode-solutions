char * getHint(char * secret, char * guess){
    int len = strlen(secret);
    
    int *bases = (int *)malloc(10 * sizeof(int));
    memset(bases, 0, 10 * sizeof(int));
    
    int *flags = (int *)malloc(10 * sizeof(int));
    memset(flags, 0, 10 * sizeof(int));
    
    int bulls = 0, cows = 0;
    
    for(int i = 0; i < len; i++) {
        bases[secret[i] - '0']++;
        flags[guess[i] - '0']++;
        
        if(secret[i] == guess[i]) {
            bulls++;
        }
    }
    
    for(int i = 0; i < 10; i++) {
        cows += bases[i] <= flags[i] ? bases[i] : flags[i];
    }
    
    cows -= bulls;
    
    char *result = (char *)malloc(20 * sizeof(char));
    int index = 0, temp = 0, size = 0;
    
    temp = bulls;
    size = 1;
    temp /= 10;
    while(temp > 0) {
        temp /= 10;
        size++;
    }
    
    index += size;
    for(int i = 1; i <= size; i++) {
        result[index - i] = bulls % 10 + '0';
        bulls /= 10;
    }
    
    result[index++] = 'A';
    
    temp = cows;
    size = 1;
    temp /= 10;
    while(temp > 0) {
        temp /= 10;
        size++;
    }
    
    index += size;
    for(int i = 1; i <= size; i++) {
        result[index - i] = cows % 10 + '0';
        cows /= 10;
    }
    
    result[index++] = 'B';
    
    result[index] = '\0';
    
    free(bases);
    free(flags);
    
    return result;
}

