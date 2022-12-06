struct Character {
    char val;
    int count;
};


int cmp(const void *a, const void *b)  {
    return ((struct Character *)b)->count - ((struct Character *)a)->count;
}

char * frequencySort(char * s){
    struct Character chars[65];
    
    for(int i = 0; i < 10; i++) {
        chars[i].val = '0' + i;
        chars[i].count = 0;
    }
    for(int i = 0; i < 26; i++) {
        chars[10 + i].val = 'a' + i;
        chars[10 + i].count = 0;
    }
    for(int i = 0; i < 26; i++) {
        chars[10 + 26 + i].val = 'A' + i;
        chars[10 + 26 + i].count = 0;
    }
    
    int len = strlen(s);
    for(int i = 0; i < len; i++) {
        if(s[i] >= '0' && s[i] <= '9') {
            chars[s[i] - '0'].count++;
        } else if(s[i] >= 'a' && s[i] <= 'z') {
            chars[10 + (s[i] - 'a')].count++;
        } else {
            chars[10 + 26 + (s[i] - 'A')].count++;
        }
    }
    
    qsort((void *)chars, 62, sizeof(struct Character), cmp);
    
    char *result = (char *)malloc((len + 1) * sizeof(char));
    int index = 0;
    
    for(int i = 0; i < 62; i++) {
        if(chars[i].count > 0) {
            for(int j = 0; j < chars[i].count; j++) {
                result[index++] = chars[i].val;
            }
        }
    }
    
    result[index] = '\0';
    return result;
}

