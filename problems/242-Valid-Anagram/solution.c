bool isAnagram(char * s, char * t){
    int lenS = strlen(s), lenT = strlen(t);
    if (lenS != lenT) {
        return false;
    }

    int chars[26];
    memset(chars, 0, 26 * sizeof(int));

    for(int i = 0; i < lenS; i++) {
        chars[s[i] - 'a']++;
        chars[t[i] - 'a']--;
    }

    for(int i = 0; i < 26; i++) {
        if(chars[i] != 0) {
            return false;
        }
    }

    return true;
}

