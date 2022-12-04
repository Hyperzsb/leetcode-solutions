void reverse(char *s) {
    int len = strlen(s);
    
    char temp;
    for(int i = 0; i < len / 2; i ++) {
        temp = s[i];
        s[i] = s[len - 1 - i];
        s[len - 1 - i] = temp;
    }
}

void zeroify(char *s) {
    int len = strlen(s);
    
    int flag = 1;
    for(int i = 0; i < len; i++) {
        if(s[i] != '0') {
            flag = 0;
            break;
        }
    }
    
    if(flag) {
        s[0] = '0';
        s[1] = '\0';
    }
} 

void singleMultiply(char *num1, char* num2, int factor) {
    int num = num1[0] - '0';
    int index = 0;
    
    for(int i = 0; i < factor; i++) {
        num1[index++] = '0';
    }
    
    int len = strlen(num2);
    int additive = 0;
    
    for(int i = 0; i < len; i++) {
        int result = num * (num2[i] - '0') + additive;
        num1[index++] = (result % 10) + '0';
        additive = result / 10;
    }
    
    if(additive > 0) {
        num1[index++] = (additive % 10) + '0';
    }
    
    num1[index] = '\0';
}

void add(char *num1, char *num2) {
    int len1 = strlen(num1), len2 = strlen(num2);
    int index1 = 0, index2 = 0;
    
    int additive = 0;
    while(index1 < len1 && index2 < len2) {
        int result = (num1[index1++] - '0') + (num2[index2++] - '0') + additive;
        num1[index1 - 1] = (result % 10) + '0';
        additive = result / 10;
    }
    
    while(index2 < len2) {
        int result = (num2[index2++] - '0') + additive;
        num1[index1++] = (result % 10) + '0';
        additive = result / 10;
    }
    
    if(additive > 0) {
        num1[index1++] = (additive % 10) + '0';
    }
    
    num1[index1] = '\0';
}

char * multiply(char * num1, char * num2){
    int len1 = strlen(num1), len2 = strlen(num2);
    char *result = (char *)malloc((len1 + 1) * (len2 + 1) * sizeof(char));
    result[0] = '\0';
    char *temp = (char *)malloc((len1 + 1) * (len2 + 1) * sizeof(char));
    temp[0] = '\0';
    
    reverse(num1);
    reverse(num2);
    
    for(int i = 0; i < len2; i++) {
        temp[0] = num2[i];
        singleMultiply(temp, num1, i);
        // printf("%s\n", temp);
        
        if(strlen(result) == 0) {
            strcpy(result, temp);
        } else {
            add(result, temp);
        }
    }
    
    reverse(result);
    zeroify(result);
    return result;
}

