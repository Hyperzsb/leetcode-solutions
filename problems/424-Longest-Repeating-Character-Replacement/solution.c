int getMax(int *flags) {
    int max = 0;
    for(int i = 0; i < 26; i++) {
        if(flags[i] > max) {
            max = flags[i];
        }
    }
    return max;
}

int characterReplacement(char * s, int k){
    int *flags = (int *)malloc(26 * sizeof(int));
    memset(flags, 0, 26 * sizeof(int));
    
    int start = 0, end = 0, max = 0, len = strlen(s);
    
    int shrink = 0;
    while(end < len) {
        if(!shrink) {
            flags[s[end] - 'A']++;
        }
        
        if(getMax(flags) + k >= end - start + 1) {
            if(end - start + 1 > max) {
                max = end - start + 1;
            }
            end++;
            shrink = 0;
        } else {
            flags[s[start] - 'A']--;
            start++;
            shrink = 1;
        }
    }
    
    return max;
}

