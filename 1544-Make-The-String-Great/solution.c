char * makeGood(char * s){
    int len = strlen(s);
    
    if(len == 0 || len == 1) {
        return s;
    }
    
    int sIndex = 0, eIndex = 1;
    while(eIndex < len) {
        if((int)s[sIndex] - (int)s[eIndex] == 32 || 
           (int)s[eIndex] - (int)s[sIndex] == 32) {
            s[sIndex--] = '!';
            s[eIndex++] = '!';
            
            while(sIndex >= 0) {
                if(s[sIndex] != '!') {
                    break;
                }
                sIndex--;
            }
            
            if(sIndex == -1) {
                sIndex = eIndex++;
            }
        } else {
            sIndex = eIndex++;
        }
    }
    
    sIndex = 0, eIndex = 0;
    
    while(eIndex < len) {
        if(s[eIndex] != '!') {
            s[sIndex++] = s[eIndex];
        }
        eIndex++;
    }
    
    s[sIndex] = '\0';
    
    return s;
}

