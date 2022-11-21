/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * int guess(int num);
 */

int guessNumber(int n){
	int start = 1, end = n, num = (end - start) / 2 + start, result = 0;
    
    while(true) {
        result = guess(num);
        
        if(result == -1) {
            end = num - 1;
            num = (end - start) / 2 + start;
        } else if(result == 1) {
            start = num + 1;
            num = (end - start) / 2 + start;
        } else {
            break;
        }
    }
    
    return num;
}

