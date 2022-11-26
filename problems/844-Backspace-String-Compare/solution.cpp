class Solution {
public:
    bool backspaceCompare(string s, string t) {
        stack<char> editorS, editorT;
        
        int lenS = s.size();
        for(int i = 0; i < lenS; i++) {
            if(s[i] == '#') {
                if(!editorS.empty()) {
                    editorS.pop();
                }
            } else {
                editorS.push(s[i]);
            }
        }
        
        int lenT = t.size();
        for(int i = 0; i < lenT; i++) {
            if(t[i] == '#') {
                if(!editorT.empty()) {
                    editorT.pop();
                }
            } else {
                editorT.push(t[i]);
            }
        }
        
        if(editorS.size() != editorT.size()) {
            return false;
        }
        
        while(!editorS.empty() && !editorT.empty()) {
            char sChar = editorS.top();
            editorS.pop();
            char tChar = editorT.top();
            editorT.pop();
            if(sChar != tChar) {
                return false;
            }
        }
        
        return true;
    }
};

