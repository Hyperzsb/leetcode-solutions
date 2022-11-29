class Solution {
private:
    int calc(int num) {
        int result = 0;
        while(num != 0) {
            result += (num % 10) * (num % 10);
            num /= 10;
        }
        
        return result;
    }
    
public:
    bool isHappy(int n) {
        map<int, int> trace;
        trace.insert({n, 1});
        
        while(n != 1) {
            n = calc(n);
            
            if(trace.find(n) != trace.end()) {
                return false;
            } else {
                trace.insert({n, 1});
            }
        }
        
        return true;
    }
};

