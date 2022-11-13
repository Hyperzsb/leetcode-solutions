int maximum69Number (int num){
    int temp = num;
    int length = 0;
    
    int digits[5];
    
    while(temp) {
        digits[length] = temp % 10;
        length += 1;
        temp /= 10;
    }
    
    for(int i = length - 1; i >= 0; i --) {
        if (digits[i] == 6) {
            digits[i] = 9;
            break;
        }
    }
    
    temp = 0;
    for(int i = length - 1; i > 0; i --) {
        temp += digits[i];
        temp *= 10;
    }
    
    temp += digits[0];
    
    return temp;
}

