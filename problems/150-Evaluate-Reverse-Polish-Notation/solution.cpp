class Solution {
private:
    bool isOperator(string &s) {
        if(s.length() == 1) {
            if(s[0] == '+' || s[0] == '-' || s[0] == '*' || s[0] == '/') {
                return true;
            } else {
                return false;
            }
        } else {
            return false;
        }
    };
    long long int parseOperand(string &s) {
        int flag = 0, index = 0;
        long long int result = 0;
        if(s[0] == '-') {
            flag = 1;
            index = 1;
        }

        for(int i = index; i < s.length(); i++) {
            result += s[i] - '0';
            result *= 10;
        }

        if(flag) {
            return 0 - result / 10;
        } else {
            return result / 10;
        }
    };
public:
    int evalRPN(vector<string>& tokens) {
        stack<long long int> numStack;
        for(int i = 0; i < tokens.size(); i++) {
            if(isOperator(tokens[i])) {
                long long int num2 = numStack.top();
                numStack.pop();
                long long int num1 = numStack.top();
                numStack.pop();
                switch(tokens[i][0]) {
                    case '+':
                        numStack.push(num1 + num2);
                        break;
                    case '-':
                        numStack.push(num1 - num2);
                        break;
                    case '*':
                        numStack.push(num1 * num2);
                        break;
                    case '/':
                        numStack.push(num1 / num2);
                        break;
                }
            } else {
                numStack.push(parseOperand(tokens[i]));
            }
        }

        return numStack.top();
    }
};

