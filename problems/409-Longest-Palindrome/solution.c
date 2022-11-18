int longestPalindrome(char * s){
    int lowercases[26], uppercases[26];
    memset(lowercases, 0, 26 * sizeof(int));
    memset(uppercases, 0, 26 * sizeof(int));
    
    int len = strlen(s);
    int maxLen = 0, remain = 0;
    for(int i = 0; i < len; i++) {
        if(s[i] >= 'A' && s[i] <= 'Z') {
            if(uppercases[(int)(s[i] - 'A')] == 0) {
                uppercases[(int)(s[i] - 'A')] = 1;
                remain++;
            } else {
                maxLen += 2;
                uppercases[(int)(s[i] - 'A')] = 0;
                remain--;
            }
        } else {
            if(lowercases[(int)(s[i] - 'a')] == 0) {
                lowercases[(int)(s[i] - 'a')] = 1;
                remain++;
            } else {
                maxLen += 2;
                lowercases[(int)(s[i] - 'a')] = 0;
                remain--;
            }
        }
    }
    
    if(remain > 0) {
        return maxLen + 1;
    } else {
        return maxLen;
    }
}

