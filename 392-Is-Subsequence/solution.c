bool isSubsequence(char * s, char * t){
    int sLen = strlen(s), tLen = strlen(t);
    
    if(sLen > tLen) {
        return false;
    }
    
    int sIndex = 0, tIndex = 0;
    
    while(sIndex < sLen && tIndex < tLen) {
        while (tIndex < tLen && s[sIndex] != t[tIndex]) {
            tIndex++;
        }
        if (tIndex < tLen && s[sIndex] == t[tIndex]) {
            sIndex++;
            tIndex++;
        }
    }
    
    if(sIndex == sLen) {
        return true;
    } else {
        return false;
    }
}

