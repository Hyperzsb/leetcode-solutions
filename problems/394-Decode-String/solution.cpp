class Solution {
public:
    string decodeString(string s) {
        stack<int> numStack;
        stack<char> strStack;
        
        int num = 0;
        for(int i = 0; i < s.size(); i++) {
            if(s[i] >= '0' && s[i] <= '9') {
                num += s[i] - '0';
                num *= 10;
            } else if(s[i] == '[') {
                numStack.push(num / 10);
                num = 0;
                strStack.push(s[i]);
            } else if(s[i] == ']') {
                string temp = "";
                while(strStack.top() != '[') {
                    temp += strStack.top();
                    strStack.pop();
                }
                strStack.pop();
                for(int i = 0; i < numStack.top(); i++) {
                    for(int j = temp.size() - 1; j >= 0; j--) {
                        strStack.push(temp[j]);
                    }
                }
                numStack.pop();
            } else {
                strStack.push(s[i]);
            }
        }
        
        string result = "";
        while(!strStack.empty()) {
            result += strStack.top();
            strStack.pop();
        }
        
        reverse(result.begin(), result.end());
        
        return result;
    }
};

