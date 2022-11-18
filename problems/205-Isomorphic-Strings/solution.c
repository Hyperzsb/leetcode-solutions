bool isIsomorphic(char * s, char * t){
    int tFlags[128];
    memset(tFlags, 0, 128 * sizeof(int));
    int sFlags[128];
    memset(sFlags, 0, 128 * sizeof(int));
    
    int len = strlen(s);
    
    int sIndex = 1, tIndex = 1;
    
    for(int i = 0; i < len; i++) {
        if (sFlags[(int)s[i]] == 0) {
            sFlags[(int)s[i]] = sIndex++;
        }
        s[i] = (char)sFlags[(int)s[i]];
        
        if (tFlags[(int)t[i]] == 0) {
            tFlags[(int)t[i]] = tIndex++;
        }
        t[i] = (char)tFlags[(int)t[i]];
        
        if(s[i] != t[i]) {
            return false;
        }
    }
    
    return true;
}

