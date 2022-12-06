class Solution {
public:
    int longestPalindrome(vector<string>& words) {
        map<string, int> wordsMap;
        int result = 0;
        
        int size = words.size();
        char pal[3];
        for(int i = 0; i < size; i++) {
            pal[0] = words[i][1];
            pal[1] = words[i][0];
            pal[2] = '\0';
            string palStr = string(pal);
            if(wordsMap.find(palStr) != wordsMap.end()) {
                if(wordsMap[palStr] == 1) {
                    wordsMap.erase(palStr);
                } else {
                    wordsMap[palStr]--;
                }
                result += 4;
            } else if (wordsMap.find(words[i]) != wordsMap.end()) {
                wordsMap[words[i]]++;
            } else {
                wordsMap.insert({words[i], 1});
            }
        }
        
        for (auto iter = wordsMap.begin(); iter != wordsMap.end(); iter++) {
            if(iter->first[0] == iter->first[1]) {
                result += 2;
                break;
            }
        }
        
        return result;
    }
};

