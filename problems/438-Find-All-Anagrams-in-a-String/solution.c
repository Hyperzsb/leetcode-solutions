/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

int equal(int *s, int *t, int size) {
    for(int i = 0; i < size; i++) {
        if(s[i] != t[i]) {
            return 0;
        }
    }
    
    return 1;
}

int* findAnagrams(char * s, char * p, int* returnSize){
    int bases[26];
    memset(bases, 0, 26 * sizeof(int));
    int flags[26];
    memset(flags, 0, 26 * sizeof(int));
    
    int slen = strlen(s), plen = strlen(p);
    
    if(slen < plen) {
        *returnSize = 0;
        return NULL;
    }
    
    int *result = (int *)malloc(slen * sizeof(int));
    int size = 0;
    
    for(int i = 0; i < plen; i++) {
        bases[p[i] - 'a']++;
        flags[s[i] - 'a']++;
    }
    
    if(equal(bases, flags, 26)) {
        result[size++] = 0;
    }
    
    for(int i = plen; i < slen; i++) {
        flags[s[i - plen] - 'a']--;
        flags[s[i] - 'a']++;
        
        if(equal(bases, flags, 26)) {
            result[size++] = i - plen + 1;
        }
    }
    
    *returnSize = size;
    
    return result;
}

